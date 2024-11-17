package main

// If implemented, this file is for assigning characteristics to fruit/vegetables.

// data structures created by yen!
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

//Vegetables: object deleted after harvest
var carrot = Crop{"Vegetable", "Carrot", 20, 4, 0, "🥕", 10, false}
var potato = Crop{"Vegetable", "Potato", 40, 6, 0, "🥔", 20, false}
var garlic = Crop{"Vegetable", "Garlic", 60, 8, 0, "🧄", 40, false}
var corn = Crop{"Vegetable", "Corn", 100, 8, 0, "🌽", 30, false}
var pumpkin = Crop{"Vegetable", "Pumpkin", 100, 14, 0, "🎃", 320, false}

//Fruits: TimePlanted reset to zero after harvest
var apple = Crop{"Fruit", "Apple", 70, 6, 0, "🍎", 10, false}
var orange = Crop{"Fruit", "Orange", 70, 6, 0, "🍊", 10, false}
var mango = Crop{"Fruit", "Mango", 110, 8, 0, "🥭", 20, false}
var peach = Crop{"Fruit", "Peach", 140, 8, 0, "🍑", 30, false}
var banana = Crop{"Fruit", "Banana", 180, 12, 0, "🍌", 45, false}
