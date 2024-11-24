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
	Rotten         bool
}

// Map to hold crops by their name
var crops = map[string]*Crop{
	"carrot":  {"Vegetable", "carrot", 20, 4, 0, "🥕", 10, false, 0, false},
	"potato":  {"Vegetable", "potato", 40, 6, 0, "🥔", 20, false, 0, false},
	"garlic":  {"Vegetable", "garlic", 60, 8, 0, "🧄", 40, false, 0, false},
	"corn":    {"Vegetable", "corn", 100, 8, 0, "🌽", 30, false, 0, false},
	"pumpkin": {"Vegetable", "pumpkin", 100, 14, 0, "🎃", 160, false, 0, false},
	"apple":   {"Fruit", "apple", 70, 6, 0, "🍎", 10, false, 200, false},
	"orange":  {"Fruit", "orange", 70, 6, 0, "🍊", 10, false, 200, false},
	"mango":   {"Fruit", "mango", 110, 8, 0, "🥭", 20, false, 300, false},
	"peach":   {"Fruit", "peach", 140, 8, 0, "🍑", 30, false, 400, false},
	"banana":  {"Fruit", "banana", 180, 12, 0, "🍌", 45, false, 500, false},
}

// A list of all crop keys for iteration
var CropKeys = []string{"carrot", "potato", "garlic", "corn", "pumpkin", "apple", "orange", "mango", "peach", "banana"}

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

// Function to get a crop by its name

// func getCropObject(cropName string) (*Crop, error) {
// 	crop, exists := crops[cropName]
// 	if !exists {
// 		return nil, fmt.Errorf("crop '%s' not found", cropName)
// 	}

// 	return crop, nil
// }

// potential difference: fruits continuing after 1 harvest
