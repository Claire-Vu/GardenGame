package main

import (
	"fmt"
)

type Crop struct {
	Name       string
	Symbol     string
	fullyGrown bool
}

// Main Game Logic
func main() {
	fmt.Println("Welcome to the Farming Simulation Game!")
	fmt.Println("Are you a new player or continuing? (Type 'new' or 'continue')")
	var playerType string
	fmt.Scanln(&playerType) // taking input from user

	var player Player // variable of type Player to store player information

	if playerType == "new" {
		var name string
		fmt.Println("Enter your name:")
		fmt.Scanln(&name)
		player = CreateNewPlayer(name)
		SavePlayer(player)

	} else if playerType == "continue" {
		fmt.Println("Enter your username:")
		var username string
		fmt.Scanln(&username)
		var err error
		player, err = LoadPlayer(username)
		if err != nil {
			fmt.Println("Could not find player.")
			return
		}
		fmt.Printf("\nWelcome back, %s!\n", player.Username)

	} else {
		fmt.Println("Invalid input.")
		return
	}

	// Example actions
	player.Resources["Potato"] = 0 // add some resources as an example
	SavePlayer(player)

	// display current player info
	fmt.Println("---Current Resources---")
	// range gives you the key and value pairs
	for resource, amount := range player.Resources {
		fmt.Printf("%s: %d\n", resource, amount)
	}
	fmt.Printf("Current Points: %d\n", player.Points)

	fmt.Println("Game saved!")

	// Garden Simulation
	garden := CreatePlot(5, 5)
	potato := Crop{Name: "Flower", Symbol: "🥔", fullyGrown: true}
	mango := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: false}
	mango2 := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: true}
	garden.Plant(0, 0, &potato)
	garden.Plant(2, 2, &mango)
	garden.Plant(2, 3, &mango2)
	garden.printGarden()
}
