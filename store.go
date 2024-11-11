//Elaine-started store-managing file, started on 11/10/24
//LIKELY GOING TO TEST-PUSH THIS CODE TO A BRANCH TO MAKE SURE I CAN DO IT

//store.go should work with, in a text-based format:
// -how many, what kinds of F&V objects currently held by player object
// -player object's gold
// -unlocked items in store (if unlocks based on playthrough points)
// -how much each store item costs (Fruit and Vegetable objects)
// -option to buy a quantity of the store item up until all player gold used
// -[MAYBE] a way to sell grown crops (depending on if they're sold from the field or here)

package main

//TODO: bring in the player object/connect the player object to the code
//TODO: look at player points to determine what items unlocked
//TODO: get prices of different Fruit and Vegetable Objects

//TODO: start(): puts other functions together

//TODO: storefront(): a display of all unlocked items, could be
//just a list mapped 1-9 and formatted like "Pumpkin Seed (2) - 320 g"

//TODO: purchase(): calculates min/max items that can be purchased with current gold,
//transfers item to player and gold from player if confirmed

//TODO leave(): exits shop menu

//MAYBE TODO: sell(): allows for player to select how many items to sell. Could
//be done in a handful of ways (sell all? sell per item?)
