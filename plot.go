package main

import "fmt"

type Plot struct {
	Rows, Cols int
	//Initializes a 2D array of pointers to Crop objects
	Plot [][]*Crop
}

type Crop struct {
	Name       string
	Symbol     string
	FullyGrown bool
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
	for i := 0; i < g.Rows; i++ {
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

// player.Points = player.Points + 1
// fmt.Println(player.Points = player.Points + 1)
