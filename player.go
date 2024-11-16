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
	Username  string
	Points    int
	Resources map[string]int
	Plot      *Plot
}

// FOR NEW PLAYER ONLY!
func CreateNewPlayer(name string, rows int, cols int) Player {
	username := name + "_" + strconv.Itoa(rand.Intn(1000))
	player := Player{
		Username:  username,
		Points:    0,
		Resources: map[string]int{"Carrot": 1}, // start with one parsnip
		Plot:      CreatePlot(rows, cols),
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
func (p *Player) PlantCrop(row, col int, crop Crop) {
	// Call the Plant method from the Plot struct
	p.Plot.Plant(row, col, &crop)
}

// GROWING THE PLAYER'S PLOT
func (p *Player) GrowPlot(numRows, numCols int) {
	// Call the GrowPlot method from the Plot struct
	p.Plot = p.Plot.GrowPlot(numRows, numCols)
}
