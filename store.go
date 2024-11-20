//Elaine-started store-managing file, started on 11/10/24
//LIKELY GOING TO TEST-PUSH THIS CODE TO A BRANCH TO MAKE SURE I CAN DO IT

//store.go should work with, in a text-based format:
// -how many, what kinds of F&V objects currently held by player object
// -player's gold
// -unlocked items in store (if unlocks based on playthrough points)
// -how much each store item costs (Fruit and Vegetable objects)
// -option to buy a quantity of the store item up until all player gold used
// -[MAYBE] a way to sell grown crops (depending on if they're sold from the field or here)

package main

import (
	"fmt"
)

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

//TODO: bring in the player object/connect the player object to the code
//TODO: get prices of different Fruit and Vegetable Objects

//func main() {
//fmt.Println("Welcome to the shop!")
//const items = getUnlocked(Player.gp)
//storefront(items)
//}

//TODO: getUnlocked(Player.gp): determines what is unlocked by the amount of
//points the player has earned. Returns an array formatted like: "8 - Pumpkin Seed (2) - 320 g"
//Each list item = "NUM TO SELECT - FRUIT/V'S NAME - FRUIT/V'S COST"
//func getUnlocked(Player.gp) []str {
//TODO: either construct the list recursively or with a for loop (if go has those)
//}

// TODO: storefront(getUnlocked's list): interacting with all unlocked items (could be just the getUnlocked() list)
// allows for selecting an item by entering the number preceding it to buy()
//func storefront(items) {
//fmt.Println(items)
//possible overarcing loop: while 'E' has not been typed
//fmt.Println("To select item, type its number and press Enter.")
//fmt.println("To leave the shop, type 'E'.")
//player input
//if player input is valid number:
//buy(Fruit/V's cost, Player.gp)
//if player input is 'E' (and not determined in overarcing loop):
//leave()
//after buying, players back at shop menu list until deciding to leave()
//}

//TODO: buy(Fruit/Vegetable.cost, Player.gp): calculates min/max item quantity
//that can be purchased with current gold,
//transfers item(s) to player and gold from player if confirmed
//returns to storefront() previous menu if cancelled

//TODO leave(): exits shop menu ('E'?)
//func leave() {
//fmt.Println("Thank you for shopping with us!")
//return to field code
//}

//MAYBE TODO: sell(Player.inventory(?)): allows for player to select how many items to sell. Could
//be done in a handful of ways (sell all? sell per item?)
