package main

// If implemented, this file is for assigning characteristics to fruit/vegetables.

// data structures created by yen!
type Vegetable struct {
	Name           string
	Cost           int
	ProductionDays int
	TimePlanted    int
	Symbol         string
	SellPrice      int
	fullyGrown     int // (0 = not planted, 1 = growing, 2 = fully grown)
}

const parsnip = Vegetable("Carrot", 20, 4, 0, "🥕", 10, 0)
const potato = Vegetable("Potato", 40, 6, 0, "🥔", 20, 0)
const garlic = Vegetable("Garlic", 60, 8, 0, "🧄", 40, 0)
const corn = Vegetable("Corn", 100, 8, 0, "🌽", 30, 0)
const pumpkin = Vegetable("Pumpkin", 100, 14, 0, "🎃", 320, 0)

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

const apple = Fruit("Apple", 70, 6, 0, "🍎", 10, 0)
const orange = Fruit("Orange", 70, 6, 0, "🍊", 10, 0)
const mango = Fruit("Mango", 110, 8, 0, "🥭", 20, 0)
const peach = Fruit("Peach", 140, 8, 0, "🍑", 30, 0)
const banana = Fruit("Banana", 180, 12, 0, "🍌", 45, 0)
