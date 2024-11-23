package main

import (
	"fmt"
)

// This file is for assigning characteristics to fruit/vegetables.

// Crop is an interface that all crops should implement
type Crop struct {
	Type           string
	Name           string
	Cost           int
	ProductionDays int
	TimePlanted    int
	Symbol         string
	SellPrice      int
	FullyGrown     bool // (0 = not planted, 1 = growing, 2 = fully grown)
	UnlockPoints   int  // Player points required to unlock crop

}

// Map to hold crops by their name
var crops = map[string]*Crop{
	"carrot":  {"Vegetable", "carrot", 20, 4, 0, "🥕", 10, false, 0},
	"potato":  {"Vegetable", "potato", 40, 6, 0, "🥔", 20, false, 0},
	"garlic":  {"Vegetable", "garlic", 60, 8, 0, "🧄", 40, false, 0},
	"corn":    {"Vegetable", "corn", 100, 8, 0, "🌽", 30, false, 0},
	"pumpkin": {"Vegetable", "pumpkin", 100, 14, 0, "🎃", 160, false, 0},
	"apple":   {"Fruit", "apple", 70, 6, 0, "🍎", 10, false, 200},
	"orange":  {"Fruit", "orange", 70, 6, 0, "🍊", 10, false, 200},
	"mango":   {"Fruit", "mango", 110, 8, 0, "🥭", 20, false, 300},
	"peach":   {"Fruit", "peach", 140, 8, 0, "🍑", 30, false, 400},
	"banana":  {"Fruit", "banana", 180, 12, 0, "🍌", 45, false, 500},
}

//A list of all crop keys for iteration
var CropKeys = []string {["carrot","potato","garlic","corn","pumpkin","apple","orange","mango","peach","banana"]}

// Function to get a crop by its name
func getCropObject(cropName string) (*Crop, error) {
	crop, exists := crops[cropName]
	if !exists {
		return nil, fmt.Errorf("crop '%s' not found", cropName)
	}

	return crop, nil
}

// potential difference: fruits continuing after 1 harvest
