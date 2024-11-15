package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// PLAYER DATA
type Player struct {
	Username  string
	Points    int
	Resources map[string]int
}

// FOR NEW PLAYER ONLY!
func CreateNewPlayer(name string) Player {
	username := name + "_" + strconv.Itoa(rand.Intn(1000))
	player := Player{
		Username:  username,
		Points:    0,
		Resources: map[string]int{"Parsnip": 1}, // start with one parsnip
	}
	fmt.Printf("Welcome, %s! Your username is %s. Remember this for future logins.\n", name, username)
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

	// create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// initialize Player with username and empty resources map
	player := Player{Username: username,
		Resources: make(map[string]int)}

	// read the first line for player points
	if scanner.Scan() {
		// convert points from string to int
		player.Points, _ = strconv.Atoi(scanner.Text())
	}

	// read the remaining lines for resources
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			item := parts[0]
			// convert string to int, ignoring errors
			amount, _ := strconv.Atoi(parts[1])
			player.Resources[item] = amount
		}
	}

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
