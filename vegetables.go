package main

import (
	"fmt"
)

// This file is for assigning characteristics to fruit/vegetables.

// Crop is an interface that all crops should implement
type Crop struct {
	Type           string // Vegetable or Fruit
	Name           string // Full name of crop
	Cost           int    // Price to buy
	ProductionDays int    // Days required to grow
	TimePlanted    int    // Time individual seed has been planted
	Symbol         string // emoji symbol for crop
	SellPrice      int    // What shop will pay for fully grown crop
	FullyGrown     bool   // (0 = not planted, 1 = growing, 2 = fully grown)
	UnlockPoints   int    // Player points required to unlock crops
	Rotten         bool   // Can become "true" if crop left in field
}

// Map to hold crops by their name
var crops = map[string]*Crop{
	"carrot":  {"Vegetable", "carrot", 10, 4, 0, "🥕", 20, false, 0, false},
	"potato":  {"Vegetable", "potato", 20, 6, 0, "🥔", 40, false, 0, false},
	"garlic":  {"Vegetable", "garlic", 40, 8, 0, "🧄", 60, false, 0, false},
	"corn":    {"Vegetable", "corn", 75, 8, 0, "🌽", 110, false, 0, false},
	"pumpkin": {"Vegetable", "pumpkin", 100, 14, 0, "🎃", 160, false, 0, false},
	"apple":   {"Fruit", "apple", 70, 6, 0, "🍎", 10, false, 10, false},
	"orange":  {"Fruit", "orange", 70, 6, 0, "🍊", 10, false, 10, false},
	"mango":   {"Fruit", "mango", 110, 8, 0, "🥭", 20, false, 20, false},
	"peach":   {"Fruit", "peach", 140, 8, 0, "🍑", 30, false, 30, false},
	"banana":  {"Fruit", "banana", 180, 12, 0, "🍌", 45, false, 40, false},
}

// A list of all crop keys for iteration in other files
var CropKeys = []string{"carrot", "potato", "garlic", "corn", "pumpkin", "apple", "orange", "mango", "peach", "banana"}

// A function for creating an instance of specific crop object
// If successful, creates and returns new Crop object and nil for the error
// If not successful, returns nil for Crop object and an error
func getCropObject(name string) (*Crop, error) {
	//
	crop, exists := crops[name]
	if !exists {
		return nil, fmt.Errorf("Crop %s not found", name)
	}

	// Create a new instance of the crop
	newCrop := Crop{
		Type:           crop.Type,
		Name:           crop.Name,
		Cost:           crop.Cost,
		ProductionDays: crop.ProductionDays,
		TimePlanted:    0,
		Symbol:         crop.Symbol,
		SellPrice:      crop.SellPrice,
		FullyGrown:     false,
		UnlockPoints:   crop.UnlockPoints,
		Rotten:         false,
	}

	return &newCrop, nil
}
