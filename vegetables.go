package main

// This file is for managing fruit and vegetable objects.
// Creates Crop struct, maps unique structs to cropDict by string Crop.Name.

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

var cropDict = map[string]Crop{

	//Vegetables: object deleted after harvest
	"carrot": Crop{"Vegetable", "Carrot", 20, 4, 0, "🥕", 10, false},
	"potato": Crop{"Vegetable", "Potato", 40, 6, 0, "🥔", 20, false},
	"garlic": Crop{"Vegetable", "Garlic", 60, 8, 0, "🧄", 40, false},
	"corn": Crop{"Vegetable", "Corn", 100, 8, 0, "🌽", 30, false},
	"pumpkin": Crop{"Vegetable", "Pumpkin", 100, 14, 0, "🎃", 320, false},

	//Fruits: TimePlanted reset to zero after harvest
	"apple": Crop{"Fruit", "Apple", 70, 6, 0, "🍎", 10, false},
	"orange": Crop{"Fruit", "Orange", 70, 6, 0, "🍊", 10, false},
	"mango": Crop{"Fruit", "Mango", 110, 8, 0, "🥭", 20, false},
	"peach": Crop{"Fruit", "Peach", 140, 8, 0, "🍑", 30, false},
	"banana": Crop{"Fruit", "Banana", 180, 12, 0, "🍌", 45, false},

},

//takes lowercase crop name, returns data structure of new crop
func createCrop(cropName string) *Crop {
	return cropDict[cropName]
}

/* DEFINING CROPS AS VARIABLES, OBSOLETE DUE TO CROPDICT

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

*/
