package main

import (
	"fmt"
)

// If implemented, this file is for assigning characteristics to fruit/vegetables.

// Crop is struct for both fruits and vegetables
type Crop struct {
	Type           string // Vegetable or Fruit
	Name           string // name of crop
	Cost           int    // Price in shop
	ProductionDays int    // TimePlanted needed to be FullyGrown
	TimePlanted    int    // how long crop has been growing
	Symbol         string // emoji for printing current field
	SellPrice      int    // What shop will pay for FullyGrown object
	FullyGrown     bool   // TimePlanted == ProductionDays, ready to sell
}

// Map to hold crops by their name
var crops = map[string]*Crop{
	"carrot":  {"Vegetable", "carrot", 20, 4, 0, "🥕", 10, false},
	"potato":  {"Vegetable", "potato", 40, 6, 0, "🥔", 20, false},
	"garlic":  {"Vegetable", "garlic", 60, 8, 0, "🧄", 40, false},
	"corn":    {"Vegetable", "corn", 100, 8, 0, "🌽", 30, false},
	"pumpkin": {"Vegetable", "pumpkin", 100, 14, 0, "🎃", 320, false},
	"apple":   {"Fruit", "apple", 70, 6, 0, "🍎", 10, false},
	"orange":  {"Fruit", "orange", 70, 6, 0, "🍊", 10, false},
	"mango":   {"Fruit", "mango", 110, 8, 0, "🥭", 20, false},
	"peach":   {"Fruit", "peach", 140, 8, 0, "🍑", 30, false},
	"banana":  {"Fruit", "banana", 180, 12, 0, "🍌", 45, false},
}

// Function to get a crop by its name
func getCropObject(cropName string) (*Crop, error) {
	crop, exists := crops[cropName]
	if !exists {
		return nil, fmt.Errorf("crop '%s' not found", cropName)
	}

	return crop, nil
}

// potential difference: fruits continuing after 1 harvest
