package main

import "fmt"

type Plot struct {
	rows, cols int
	//Initializes a 2D array of pointers to Crop objects
	plot [][]*Crop
}

// Initializes a garden plot of size rows x cols
func CreatePlot(rows, cols int) *Plot {
	g := &Plot{
		rows: rows,
		cols: cols,
		// Makes an array with of rows length
		plot: make([][]*Crop, rows),
	}

	// Initialize the 2D slice with empty soil (nil)
	for i := 0; i < rows; i++ {
		// Each row has cols elements
		g.plot[i] = make([]*Crop, cols)

		// nil for each cell
		for j := 0; j < cols; j++ {
			g.plot[i][j] = nil
		}
	}
	return g
}

// GrowPlot expands the current plot by numRows and numCols
// (g *Plot) represents a receiver (which is just a pointer to the current Plot struct)
// structure of go functions: func functionName(params) returnType{}
func (g *Plot) GrowPlot(numRows, numCols int) *Plot {
	// Calculate new dimensions
	newHeight := g.rows + numRows
	newWidth := g.cols + numCols

	// Create a new expanded plot
	newPlot := CreatePlot(newHeight, newWidth)

	// Copy over old plot to new plot
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			newPlot.plot[i][j] = g.plot[i][j]
		}
	}

	return newPlot
}

func (g *Plot) printGarden() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			if g.plot[i][j] == nil {
				fmt.Print("🟫 ")
			} else {
				// Plant object should have a symbol value
				// If plant is not yet fully grown then
				if g.plot[i][j].fullyGrown {
					fmt.Print(g.plot[i][j].Symbol, " ")
				} else {
					fmt.Print("🌱 ")
				}
			}
		}
		fmt.Println()
	}
}

func (g *Plot) Plant(row, col int, crop *Crop) {
	if row >= 0 && row < g.rows && col >= 0 && col < g.cols {
		if g.plot[row][col] == nil {
			g.plot[row][col] = crop
			fmt.Println("Crop successfully planted!")
		} else {
			fmt.Println("Cannot plant at that location: There is already a crop at that location!")
		}
	} else {
		fmt.Println("Cannot plant: the row/col is outside the plot")
	}
}

// player.Points = player.Points + 1
// fmt.Println(player.Points = player.Points + 1)
