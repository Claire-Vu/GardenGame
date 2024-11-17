package main

// If implemented, this file is for assigning characteristics to fruit/vegetables.

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
}

// data structures created by yen!
type Vegetable struct {
	Name           string
	Cost           int
	ProductionDays int
	TimePlanted    int
	Symbol         string
	SellPrice      int
	fullyGrown     bool // (0 = not planted, 1 = growing, 2 = fully grown)
}

var carrot = Vegetable{"Carrot", 20, 4, 0, "🥕", 10, false}
var potato = Vegetable{"Potato", 40, 6, 0, "🥔", 20, false}
var garlic = Vegetable{"Garlic", 60, 8, 0, "🧄", 40, false}
var corn = Vegetable{"Corn", 100, 8, 0, "🌽", 30, false}
var pumpkin = Vegetable{"Pumpkin", 100, 14, 0, "🎃", 320, false}

// potential difference: fruits continuing after 1 harvest
// data structure created by yen!
type Fruit struct {
	Name           string
	Cost           int
	ProductionDays int
	TimePlanted    int
	Symbol         string
	SellPrice      int
	GrowthStage    int // (0 = not planted, 1 = growing, 2 = fully grown)
}

var apple = Fruit{"Apple", 70, 6, 0, "🍎", 10, 0}
var orange = Fruit{"Orange", 70, 6, 0, "🍊", 10, 0}
var mango = Fruit{"Mango", 110, 8, 0, "🥭", 20, 0}
var peach = Fruit{"Peach", 140, 8, 0, "🍑", 30, 0}
var banana = Fruit{"Banana", 180, 12, 0, "🍌", 45, 0}
