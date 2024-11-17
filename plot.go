package main

import "fmt"

type Plot struct {
	Rows, Cols int
	//Initializes a 2D array of pointers to Crop objects
	Plot [][]*Crop
}

// Initializes a garden Plot of size Rows x Cols
func CreatePlot(Rows, Cols int) *Plot {
	g := &Plot{
		Rows: Rows,
		Cols: Cols,
		// Makes an array with of Rows length
		Plot: make([][]*Crop, Rows),
	}

	// Initialize the 2D slice with empty soil (nil)
	for i := 0; i < Rows; i++ {
		// Each row has Cols elements
		g.Plot[i] = make([]*Crop, Cols)

		// nil for each cell
		for j := 0; j < Cols; j++ {
			g.Plot[i][j] = nil
		}
	}
	return g
}

// GrowPlot expands the current Plot by numRows and numCols
// (g *Plot) represents a receiver (which is just a pointer to the current Plot struct)
// structure of go functions: func functionName(params) returnType{}
func (g *Plot) GrowPlot(numRows, numCols int) *Plot {
	// Calculate new dimensions
	newHeight := g.Rows + numRows
	newWidth := g.Cols + numCols

	// Create a new expanded Plot
	newPlot := CreatePlot(newHeight, newWidth)

	// Copy over old Plot to new Plot
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			newPlot.Plot[i][j] = g.Plot[i][j]
		}
	}

	return newPlot
}

func (g *Plot) printGarden() {
	// Print column number
	fmt.Print(" ") // Padding for row numbers
	for col := 0; col < g.Cols; col++ {
		fmt.Printf("%2d ", col) // Print column numbers with 2-digit width
	}
	fmt.Println()

	for i := 0; i < g.Rows; i++ {
		// Print row number and row
		fmt.Print(i, " ")
		for j := 0; j < g.Cols; j++ {
			if g.Plot[i][j] == nil {
				fmt.Print("🟫 ")
			} else {
				// Plant object should have a symbol value
				// If plant is not yet fully grown then
				if g.Plot[i][j].FullyGrown {
					fmt.Print(g.Plot[i][j].Symbol, " ")
				} else {
					fmt.Print("🌱 ")
				}
			}
		}
		fmt.Println()
	}
}

func (g *Plot) Plant(row, col int, crop *Crop) {
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		if g.Plot[row][col] == nil {
			g.Plot[row][col] = crop
			fmt.Println("Crop successfully planted!")
		} else {
			fmt.Println("Cannot plant at that location: There is already a crop at that location!")
		}
	} else {
		fmt.Println("Cannot plant: the row/col is outside the Plot")
	}
}

// Function that harvests all fully grown crops and returns map of all harvested items
// with key being the crop object that was harvested and value being the quantity harvested

func (g *Plot) HarvestAll() map[*Crop]int {
	harvestedCrop := make(map[*Crop]int)
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			// If there is a crop and it is fullyGrown
			if g.Plot[i][j] != nil && g.Plot[i][j].FullyGrown {

				// Updates harvestedCrop map, if the crop object
				// already exist then add to quantity, else add it to map
				if quantity, ok := harvestedCrop[g.Plot[i][j]]; ok {
					harvestedCrop[g.Plot[i][j]] = quantity + 1
				} else {
					harvestedCrop[g.Plot[i][j]] = 1
				}
				// If it is a fruit then the growth days gets reset
				if g.Plot[i][j].Type == "Fruit" {
					g.Plot[i][j].TimePlanted = 0
					g.Plot[i][j].FullyGrown = false
				} else {
					g.Plot[i][j] = nil
				}
			}
		}
	}
	return harvestedCrop
}

func (g *Plot) removeItem(row, col int) {
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		if g.Plot[row][col] != nil {
			g.Plot[row][col] = nil
			fmt.Println("Crop successfully removed!")
		} else {
			fmt.Println("There is no crop to remove.")
		}
	} else {
		fmt.Println("Cannot remove: the row/col is outside the Plot")
	}
}

func (g *Plot) updateCrops() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			// For every crop increase day by 1 and change Fully grown
			if g.Plot[i][j] != nil {
				g.Plot[i][j].TimePlanted += 1

				// Crop is Fully Grown if days exceed production days
				if g.Plot[i][j].TimePlanted >= g.Plot[i][j].ProductionDays {
					g.Plot[i][j].FullyGrown = true
				}
			}
		}
	}
}

// player.Points = player.Points + 1
// fmt.Println(player.Points = player.Points + 1)
