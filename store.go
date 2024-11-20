//Elaine-started store-managing file, started on 11/10/24

//store.go manages the shop, primarily used for buying and selling seeds and crops

package main

import (
	"fmt"
	"strconv"
	"strings"
	//possibly "player.go" or a player file, if needed to access the player's points/gold
)

func main() {
	fmt.Println("Welcome to the shop!")
	//TODO: assign player to current player
	const purchasables = getUnlocked(Player.points)
	fmt.Println("To select item, type its name and press Enter.")
	fmt.Println("To sell items, type \"sell\".")
	fmt.println("To leave the shop, type 'E'.")
	var shopChoice string // user input
	fmt.Scanln(&shopChoice)
	if strings.ToLower(shopChoice) == "e" {
		fmt.Println("Goodbye! We hope you'll shop with us again soon :)")
		//TODO: LEAVE SHOP, RETURN TO PLOT
	}
	if strings.ToLower(shopChoice) == "sell" {
		sell()
	} else {
		var chosen = getCropObject(strings.ToLower(shopChoice))
		buy(chosen)
	}
}

//Returns formatted string list of unlocked crops and buy price
func getUnlockedBuy(player.Points) []string {
	var buyList []string
	for i := 0; i < len(Crops); i++ {
		buyList[i] = Crops[i][Name] + " - " + Crops[i][Cost]
		//"FRUIT/V'S NAME - FRUIT/V'S COST" "Pumpkin - 100 g"
		//if player.Points >= Crop.Unlock... TODO: if we want to do by point
	}
	// buyList[len(Crops)] = "SELL CROPS" //opens sell menu
	return buyList
}

//Returns formatted string list of crops in player inventory and sell price
func getUnlockedSell() []string {
	var sellList []string
	if len(player.CropInventory) == 0 { //no items in list
		return sellList
	} else {
		for i := 0; i < len(player.CropInventory); i++ {
			sellList[i] = player.CropInventory[i][Name] + " - " + player.CropInventory[i][SellPrice]
			//FRUIT/V'S NAME - FRUIT/V'S SELL PRICE" "Pumpkin - 320 g"
		}
	}
	return sellList
}

// prints the getUnlockedBuy() or getUnlockedSell() list line by line
func printUnlocked(unlockedList []string) {
	for i := 0; i < len(unlockedList); i++ {
		fmt.Println(unlockedList[i])
	}
}

// TODO: storefront(list): interacting with all unlocked items (could be just the getUnlocked() list)
// allows for selecting an item by entering the number preceding it to buy()
//func storefront(unlockedList []string) {
	
	//shopChoice = strconv.Atoi(shopChoice)                  //converts input to integer
	///if shopChoice >= len(unlockedList) || shopChoice < 1 { //out of range
	///	fmt.Println("ERROR: given number not in range")
	///} else {
	///	if shopChoice == len(unlockedList) { // if chose final sell option
	///		sell()
	///	} else {
	///		buy()
	///	}

	//after buying, players back at shop menu list until deciding to leave()
//}

// TODO: buy(Fruit/Vegetable.cost, Player.gp): calculates min/max item quantity
// that can be purchased with current gold,
// transfers item(s) to player and gold from player if confirmed
// returns to storefront() previous menu if cancelled
func buy(cropName) Crop {
	//grab cropName[Cost]
}

// function to add crop seeds and remove gold from the player file
func buyPlayer() *Player {
	//error checking: if quantity is not in range
	//add seeds of specified quantity
	//remove gold corresponding to seeds
}

// sell(Player.inventory(?)): allows for player to select how many items to sell
func sell() *Player {
	if length player.CropInventory == 0 { //empty inventory
		fmt.Println("You don't currently have any crops to sell.")
		return
	} else {
		var sellList = getUnlockedSell()

	}
}

// function to remove fully grown crops and add gold to the player file
func sellPlayer(quantity int) *Player {
	//error checking: if quantity is not in range
	//increase player gold
	//remove corresponding fruits/vegetables from player inventory
	//TODO: if points implemented, increase player points
}

//TODO leave(): exits shop menu ('E'?)
//func leave() {
//fmt.Println("Thank you for shopping with us!")
//return to field code
//}
