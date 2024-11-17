package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// PLAYER DATA
type Player struct {
	Username      string
	Points        int
	SeedStorage   map[string]int // Tracks the player's available seeds (e.g., carrot seeds)
	CropInventory map[string]int // Tracks harvested crops (e.g., carrots, potatoes)
	Plot          *Plot
}

// FOR NEW PLAYER ONLY!
func CreateNewPlayer(name string, rows int, cols int) Player {
	username := name + "_" + strconv.Itoa(rand.Intn(1000))
	player := Player{
		Username: username,
		Points:   0,
		SeedStorage: map[string]int{
			"carrot":  1,
			"potato":  1,
			"garlic":  1,
			"corn":    1,
			"pumpkin": 1,
		}, // Start with one of each vegetable seed
		CropInventory: make(map[string]int),
		Plot:          CreatePlot(rows, cols),
	}
	fmt.Printf("\nWelcome, %s! Your username is %s. Remember this for future logins.\n", name, username)
	return player
}

// LOADS AN EXISTING PLAYER
func LoadPlayer(username string) (Player, error) {
	fileName := username + ".json"
	file, err := os.Open(fileName)
	if err != nil {
		return Player{}, fmt.Errorf("player not found")
	}
	defer file.Close()

	// Create a variable to hold the player data
	var player Player

	// Decode the entire file into the player struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&player)
	if err != nil {
		return Player{}, fmt.Errorf("error decoding player data: %v", err)
	}

	// Return the loaded player struct
	return player, nil
}

// SAVES PLAYER DATA TO A FILE
func SavePlayer(player Player) {
	fileName := player.Username + ".json"
	file, err := os.Create(fileName)
	// if there is an error
	if err != nil {
		fmt.Println("Error saving player data.")
		return
	}
	defer file.Close()

	// convert player struct to JSON format
	data, err := json.MarshalIndent(player, "", "  ")
	if err != nil {
		fmt.Println("Error converting player data to JSON.")
		return
	}

	// write the JSON data to the file
	file.Write(data)
}

// PLANTING CROP IN THE PLAYER'S PLOT
func (p *Player) PlantCrop(row, col int, crop Crop) error {
	//Check if the player has enough seeds to plant
	if p.SeedStorage[crop.Name] <= 0 {
		return fmt.Errorf("not enough %s seeds to plant", crop.Name)
	}

	p.Plot.Plant(row, col, &crop)
	fmt.Printf("Planted %s at row %d, column %d.\n", crop.Name, row, col)

	p.SeedStorage[crop.Name]--
	return nil
}

// GROWING THE PLAYER'S PLOT
func (p *Player) GrowPlot(numRows, numCols int) {
	p.Plot = p.Plot.GrowPlot(numRows, numCols)
}

// DISPLAY PLAYER'S INVENTORY
func (p *Player) DisplayInfo() {
	fmt.Printf("Username: %s\n", p.Username)
	fmt.Printf("Points: %d\n", p.Points)
	fmt.Println("Seed Storage:")
	for crop, count := range p.SeedStorage {
		fmt.Printf("  %s: %d\n", crop, count)
	}

	if len(p.CropInventory) == 0 {
		fmt.Println("Crop Inventory: No harvest yet.")
	} else {
		fmt.Println("Crop Inventory:")
		for crop, count := range p.CropInventory {
			fmt.Printf("  %s: %d\n", crop, count)
		}
	}
}
