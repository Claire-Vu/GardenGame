//store.go manages the shop, used for buying seeds and selling crops

//functions:
//StoreFront(): main store loop
//getUnlocked(): returns list of strings with all unlocked crops, along with their buy prices
//getInventory(): returns a list of strings with all crops present in player inventory, along with their sell prices
//printLists(): prints getUnlocked() or getInventory() lists one by one
//buyItems(): manages player inventory adding seeds, player gold decrease
//sellItems(): manages player inventory removing fully grown crops, player gold and points++

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Main store loop:
func (p *Player) StoreFront() string {
	ClearConsole()
	fmt.Printf("----------------------SHOP--------------------------------\n")
	fmt.Printf("Welcome to the shop! You currently have %d gold.\n", p.Gold)
	fmt.Println("To buy items, type \"buy\".")
	fmt.Println("To sell items, type \"sell\".")
	fmt.Println("To leave the shop, type \"E\".")

	// Scanner to prompt for input
	scanner := bufio.NewScanner(os.Stdin)

	var shopChoice string // user input
	if scanner.Scan() {
		shopChoice = scanner.Text()
	} else {
		shopChoice = "invalid"
	}

	// LEAVEING SHOP
	if strings.ToLower(shopChoice) == "e" {
		fmt.Println("Goodbye! We hope you'll shop with us again soon :)")
		return "Exit"

		// BUYING
	} else if strings.ToLower(shopChoice) == "buy" {
		var cropList = p.getUnlocked() // unlocked crop list
		printLists(cropList)           // prints shop stock

		fmt.Println("To buy a listed item, type its name and press Enter.")
		var buyChoice string // user input for what to buy
		if scanner.Scan() {
			buyChoice = strings.TrimSpace(scanner.Text())
			_, notFound := getCropObject(buyChoice)
			if notFound != nil {
				buyChoice = "Invalid"
			}
		}

		if buyChoice != "Invalid" {
			// quantityToBuyStr holds the input value
			// and quantityToBuy checks only valid if input value
			// was a string
			var quantityToBuyStr string
			var quantityToBuy int

			// Validates input is an integer
			fmt.Printf("How many %s(s) would you like to buy? (%d gold held) ", buyChoice, p.Gold)
			if scanner.Scan() {
				quantityToBuyStr = strings.TrimSpace(scanner.Text())
				// If the input cannot be converted to a string then set quantityToBuy
				// to -1 (invalid input)
				inputVal, errStr := strconv.Atoi(quantityToBuyStr)
				if errStr != nil {
					quantityToBuy = -1
				} else {
					quantityToBuy = inputVal
				}
			}
			// If quantityToBuy and buyChoice are valid
			if quantityToBuy > 0 {
				// We proceed to buy the item and check for any errors
				// associated with executing that function
				errBuy := p.buyItems(buyChoice, quantityToBuy)
				if errBuy != nil {
					fmt.Println(strings.ToUpper(errBuy.Error()))
				} else {
					fmt.Printf("Successfully purchased %d %s(s)!", quantityToBuy, buyChoice)
				}
				// If the quantityToBuy was not valid
			} else {
				fmt.Printf("\n Invalid quantity input")
			}

			// If the buyChoice was not valid
		} else {
			fmt.Printf("\n Invalid crop input")
		}

		// SELLING
	} else if strings.ToLower(shopChoice) == "sell" {
		fmt.Println("To sell a listed item, type its name and press Enter.")
		var inventory = p.getInventory()
		printLists(inventory) // prints player-held crops & quantities

		var sellChoice string // user input
		if scanner.Scan() {
			sellChoice = strings.TrimSpace(scanner.Text())
		}

		if (sellChoice != "") && (p.CropInventory[sellChoice] > 0) { // added check to make sure item in inventory
			fmt.Printf("How many %s would you like to sell? (You have %d)", sellChoice, p.CropInventory[sellChoice])

			var quantityToSellStr string
			var quantityToSell int

			// Validates input is an integer
			if scanner.Scan() {
				quantityToSellStr = strings.TrimSpace(scanner.Text())
				// If the input cannot be converted to a string then set quantityToSell
				// to -1 (invalid input)
				inputVal, errStr := strconv.Atoi(quantityToSellStr)
				if errStr != nil {
					quantityToSell = -1
				} else {
					quantityToSell = inputVal
				}
			}

			// If valid int input then sell the items and check all
			// associated errors with that action
			if quantityToSell > 0 {
				p.sellItems(sellChoice, quantityToSell)
			} else {
				fmt.Printf("\nInvalid quantity input")
			}
		} else if sellChoice != "" {
			fmt.Printf("\n Input \"%s\" not understood. Returning to shop menu", sellChoice)
		}
	} else { // INVALID ACTION INPUT:
		fmt.Println("\nInput not understood. Please type 'buy', 'sell', or 'E'.")
	}
	// Leaves 2 second delay for user to look at err/success messages
	fmt.Println("\n Loading game ...")
	time.Sleep(2 * time.Second)

	return "notExit"
}

// For buying. Generates a list of the products the shop is carrying.
// This list only includes crops the player has enough points to unlock.
// Returns all unlocked crops and their seed prices as a list of strings.
func (p *Player) getUnlocked() []string {
	var cropList []string
	for i := 0; i < len(CropKeys); i++ {
		curCrop, err := getCropObject(CropKeys[i])
		if err != nil {
			// print error message
		}
		if curCrop.UnlockPoints <= p.Points { // If player has enough points for crop
			// Format and append the crop details to the list
			cropDetails := curCrop.Symbol + " " + curCrop.Name + " - Buy for " + strconv.Itoa(curCrop.Cost) + " gold"
			cropList = append(cropList, cropDetails)
		}
	}
	return cropList
}

// For selling. Display's player's current sellable crops.
// This list only includes crops the player currently has in CropInventory.
// Returns all crops player is holding and sell prices as a list of strings.
func (p *Player) getInventory() []string {
	index := 1 // index of strLi after first entry
	var strLi []string
	strLi = append(strLi, "You have... ")
	for i := 0; i < len(CropKeys); i++ {
		if p.CropInventory[CropKeys[i]] > 0 { //if anything in crop inventory slot
			curCrop, err := getCropObject(CropKeys[i])
			if err != nil {
				// print error message
			}
			sellDetails := CropKeys[i] + " - " + strconv.Itoa(p.CropInventory[CropKeys[i]]) + " held - sell for " + strconv.Itoa(curCrop.SellPrice) + " gold each"
			strLi = append(strLi, sellDetails)
			index++
		}
	}
	if len(strLi) == 1 { //if nothing in player inventory
		strLi = append(strLi, "No fully grown crops.")
		strLi = append(strLi, "Press \"Enter\" to return to the shop.") // button to return player to main shop
	}
	return strLi
}

// prints the getUnlocked() or getInventory() lists line by line
func printLists(curList []string) {
	for i := 0; i < len(curList); i++ {
		fmt.Println(curList[i])
	}
}

// Handles cropObjects being bought by the player
// Error checks, adds seeds to storage, removes gold from player
func (p *Player) buyItems(cropToBuy string, quantityToBuy int) error {
	cropObject, notValidCrop := getCropObject(cropToBuy)

	if notValidCrop == nil { // if no error (cropObject found)
		if p.Gold >= (cropObject.Cost * quantityToBuy) {
			// removes gold from player, adds purchased seeds
			p.Gold -= (cropObject.Cost * quantityToBuy)
			p.SeedStorage[cropToBuy] += quantityToBuy
			return nil // return no error
		} else {
			return fmt.Errorf("invalid: not enough gold to purchase %d %s crops", quantityToBuy, cropToBuy)
		}
	} else { // if getCropObject() returned an error
		return fmt.Errorf("invalid: not a valid crop")
	}
}

// Handles fully grown crops being sold by player
// Error checks, removes crops from inventory, adds gold
func (p *Player) sellItems(cropToSell string, quantityToSell int) error {
	// Ok: Identifies if player cropInventory has key cropToSell
	if quantityInInventory, ok := p.CropInventory[cropToSell]; ok {
		// Checks if has enough quantity to sell
		if quantityInInventory >= quantityToSell {

			// getCropObject returns the cropObject or error
			cropObject, notValidCrop := getCropObject(cropToSell)

			if notValidCrop == nil { // If no error, enough to sell
				//adds gold and points, removes crops
				p.Gold += (cropObject.SellPrice * quantityToSell)
				p.Points += quantityToSell
				p.CropInventory[cropToSell] -= quantityToSell
				// if crop quantity becomes 0 then remove it from the inventory
				if p.CropInventory[cropToSell] == 0 {
					delete(p.CropInventory, cropToSell)
				}
				fmt.Printf("Successfully sold %d %s(s) for %d gold!", quantityToSell, cropObject.Name, (cropObject.SellPrice * quantityToSell))
				// return no error
				return nil
			}

		} else {
			fmt.Println("invalid: not enough crops to sell")
			return fmt.Errorf("invalid: not enough crops to sell")
		}
	}
	return fmt.Errorf("invalid: you do not have crop: %s", cropToSell)
}
