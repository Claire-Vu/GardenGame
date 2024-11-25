//Elaine-started store-managing file, started on 11/10/24

//store.go manages the shop, primarily used for buying and selling seeds and crops

//functions:
//main(): runs code
//stringGold(): returns current player object's gold as a string
//getUnlocked(): returns list of strings with all unlocked crops, along with their buy prices
//getInventory(): returns a list of strings with all crops present in player inventory, along with their sell prices
//printUnlocked(): prints getUnlocked() or getInventory() lists one by one
//buyItems(): manages player inventory adding seeds, player gold decrease
//sellItems(): manages player inventory removing fully grown crops, player gold and points++

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: IS THERE A WAY TO PASS PLAYER OBJECT INTO MAIN
func (p *Player) StoreFront() string {
	ClearConsole()
	fmt.Printf("----------------------SHOP--------------------------------\n")
	fmt.Printf("Welcome to the shop! You currently have %d gold.\n", p.Gold)
	fmt.Println("To buy items, type \"buy\".")
	fmt.Println("To sell items, type \"sell\".")
	fmt.Println("To leave the shop, type \"E\".")
	var shopChoice string // user input
	fmt.Scanln(&shopChoice)
	if strings.ToLower(shopChoice) == "e" {
		fmt.Println("Goodbye! We hope you'll shop with us again soon :)")
		return "Exit"
	}
	// BUYING:
	if strings.ToLower(shopChoice) == "buy" {
		var cropList = p.getUnlocked()
		printUnlocked(cropList)
		fmt.Println("To buy a listed item, type its name and press Enter.")
		var buyChoice string // user input
		fmt.Scanln(&buyChoice)
		if buyChoice != "" {
			fmt.Printf("How many %s would you like to buy? You have %d gold.", buyChoice, p.Gold)
			var quantityToBuy int // user input
			fmt.Scanln(&quantityToBuy)
			p.buyItems(buyChoice, quantityToBuy)
		}
	}
	// SELLING:
	if strings.ToLower(shopChoice) == "sell" {
		//var cropList = p.getUnlocked()
		//printUnlocked(cropList)
		var inventory = p.getInventory()
		printUnlocked(inventory)
		fmt.Println("To sell a listed item, type its name and press Enter.")
		var sellChoice string // user input
		fmt.Scanln(&sellChoice)
		if sellChoice != "" {
			fmt.Printf("How many %s would you like to sell? (%d total)", sellChoice, p.CropInventory[sellChoice])
			var quantityToSell int // user input
			fmt.Scanln(&quantityToSell)
			p.sellItems(sellChoice, quantityToSell)
		}
	} else { // invalid input:
		fmt.Println("Input not understood. Please type 'buy', 'sell', or 'E'.")
	}
	return "notExit"
}

// creates a list of strings -- all unlocked crops and their prices
func (p *Player) getUnlocked() []string {
	var cropList []string
	for i := 0; i < len(CropKeys); i++ {
		curCrop, err := getCropObject(CropKeys[i])
		if err != nil {
			// print error message
		}
		if curCrop.UnlockPoints <= p.Points { // If player has enough points for crop
			// Format and append the crop details to the list
			cropDetails := curCrop.Name + " - Buy for " + strconv.Itoa(curCrop.Cost) + " gold"
			cropList = append(cropList, cropDetails)
		}
	}
	return cropList
}

// Creates a list of strings: all crops player is holding and sell prices
func (p *Player) getInventory() []string {
	var index int
	index = 1
	var strLi []string
	strLi = append(strLi, "You have... ")
	for i := 0; i < len(CropKeys); i++ {
		if p.CropInventory[CropKeys[i]] > 0 { //if anything in crop inventory slot
			curCrop, err := getCropObject(CropKeys[i])
			if err != nil {
				// print error message
			}
			sellDetails := strconv.Itoa(p.CropInventory[CropKeys[i]]) + " " + CropKeys[i] + " - sell for " + strconv.Itoa(curCrop.SellPrice) + "gold each"
			strLi = append(strLi, sellDetails)
			index++
		}
	}
	if len(strLi) == 1 { //if nothing in player inventory
		strLi = append(strLi, "No crops in inventory. Plant and harvest them from the field!")
	}
	return strLi
}

// prints the getUnlocked() or getInventory() lists line by line
func printUnlocked(unlockedList []string) {
	for i := 0; i < len(unlockedList); i++ {
		fmt.Println(unlockedList[i])
	}
}

func (p *Player) buyItems(cropToBuy string, quantityToBuy int) error {
	cropObject, notValidCrop := getCropObject(cropToBuy)
	// Checks player gold >= cost
	if p.Gold >= (cropObject.Cost * quantityToBuy) {
		// Buys the crops

		// If no error
		if notValidCrop == nil {
			// removes gold from player
			p.Gold -= (cropObject.Cost * quantityToBuy)
			// adds crop seed to inventory
			// TODO: DO I NEED TO CHECK TO MAKE SURE SEED INVENTORY ENTRY EXISTS FIRST
			p.SeedStorage[cropToBuy] += quantityToBuy
			// return no error
			return nil
		}

	} else {
		return fmt.Errorf("INVALID: Not enough gold to purchase %s %s crops.", strconv.Itoa(quantityToBuy), cropToBuy)
	}
	return nil
}

func (p *Player) sellItems(cropToSell string, quantityToSell int) error {
	// Ok: Identifies if player cropInventory has key cropToSell
	if quantityInInventory, ok := p.CropInventory[cropToSell]; ok {
		// Checks if has enough quantity to sell
		if quantityInInventory >= quantityToSell {
			// Sells the crops

			// Crop object we are selling
			// getCropObject returns the cropObject or error
			cropObject, notValidCrop := getCropObject(cropToSell)

			// If no error
			if notValidCrop == nil {
				// gives gold to player
				p.Gold += (cropObject.SellPrice * quantityToSell)
				//updates points
				p.Points += quantityToSell
				// takes away crop
				p.CropInventory[cropToSell] -= quantityToSell
				// if crop quantity becomes 0 then remove it from the inventory
				if p.CropInventory[cropToSell] == 0 {
					delete(p.CropInventory, cropToSell)
				}
				// return no error
				return nil
			}

		} else {
			return fmt.Errorf("INVALID: Not enough crops to sell")
		}
	}
	return fmt.Errorf("INVALID: You do not have crop: %s", cropToSell)
}
