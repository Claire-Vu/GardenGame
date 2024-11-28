package main

import "fmt"

type Plot struct {
	Rows, Cols int
	//Initializes a 2D array of pointers to Crop objects
	Plot      [][]*Crop
	PlotLevel int
}

// INITIALIZES AND RETURNS A GARDEN PLOT STRUCT
func CreatePlot(Rows, Cols int) *Plot {
	// Initalizes new instance of Plot struct
	g := &Plot{
		Rows: Rows,
		Cols: Cols,
		// 2D array with Rows length
		Plot:      make([][]*Crop, Rows),
		PlotLevel: 0,
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

// (g *Plot) represents a receiver (which is just a pointer to object
// that calls the function).
// EXPANDS CURRENT PLOT BY NUMROWS AND NUMCOLS
func (g *Plot) GrowPlot(numRows, numCols int) *Plot {
	// Calculates new dimensions
	newHeight := g.Rows + numRows
	newWidth := g.Cols + numCols

	// Creates a new Plot with new dimensions
	newPlot := CreatePlot(newHeight, newWidth)

	// Copying over the old plot
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			newPlot.Plot[i][j] = g.Plot[i][j]
		}
	}
	return newPlot
}

// PRINTS THE GARDEN USING EMOJI'S
func (g *Plot) printGarden() {
	// Print column numbers
	fmt.Print(" ") // Padding for row numbers
	for col := 0; col < g.Cols; col++ {
		fmt.Printf("%2d ", col) // Each number takes the space of 2 digits
	}
	fmt.Println()

	for i := 0; i < g.Rows; i++ {
		// Prints row number and then prints row contents
		fmt.Print(i, " ")
		for j := 0; j < g.Cols; j++ {
			// If nothing on plot
			if g.Plot[i][j] == nil {
				fmt.Print("🟫 ")
			} else {
				// If plant is not yet fully grown then it is a leaf or tree
				// when it is fully grown it's symbol is shown

				if !g.Plot[i][j].Rotten {
					if g.Plot[i][j].FullyGrown {
						fmt.Print(g.Plot[i][j].Symbol, " ")
					} else {
						if g.Plot[i][j].Type == "Fruit" {
							fmt.Print("🌳 ")
						} else {
							fmt.Print("🌱 ")
						}
					}
				} else {
					fmt.Print("🦠 ")
				}
			}
		}
		fmt.Println()
	}
}

// PLANTS CROP IN THE PLOT AT COORDINATES (ROW,COL)
func (g *Plot) Plant(row, col int, crop *Crop) error {
	// Checks that the cell is within bounds of the garden
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		// Only plants if there is no crop present at location
		if g.Plot[row][col] == nil {
			g.Plot[row][col] = crop
			return nil
		} else {
			return fmt.Errorf("there is already a crop at that location")
		}
	} else {
		return fmt.Errorf("cannot plant: the row/col is outside the plot")
	}
}

// HARVEST ALL FULLYGROWN CROPS ON PLOT
func (g *Plot) HarvestAll() map[string]int {
	// map of all harvested items, key is the crop name and value is the quantity
	harvestedCrop := make(map[string]int)
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			// If there is a crop at the current location, it is fullyGrown,
			// and it's not rotten, then it can be harvested
			if g.Plot[i][j] != nil && g.Plot[i][j].FullyGrown && !g.Plot[i][j].Rotten {

				// Updates harvestedCrop map: ok checks if the item
				// is in the list, it is is then it only updates quantity
				// else it adds the new crop as a new key
				if quantity, ok := harvestedCrop[g.Plot[i][j].Name]; ok {
					harvestedCrop[g.Plot[i][j].Name] = quantity + 1
				} else {
					harvestedCrop[g.Plot[i][j].Name] = 1
				}

				// Vegetables can only be harvested once, Fruits are harvested
				// multiple times
				if g.Plot[i][j].Type == "Vegetable" {
					g.Plot[i][j] = nil
				} else {
					// Fruits gets reset
					g.Plot[i][j].TimePlanted = 0
					g.Plot[i][j].FullyGrown = false
				}

			}
		}
	}
	return harvestedCrop
}

// REMOVES/DIGS OUT ITEM AT LOCATION (ROW,COL)
func (g *Plot) removeItem(row, col int) error {
	// Checks if input within bounds of plot
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		// If there is an item there, then remove it
		if g.Plot[row][col] != nil {
			g.Plot[row][col] = nil
			return nil
		} else {
			return fmt.Errorf("there is no crop to remove")
		}
	} else {
		return fmt.Errorf("cannot remove: the row/col is outside the plot")
	}
}

// UPDATES CROP AT END OF DAY
func (g *Plot) updateCrops() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			// For every crop increase day by 1
			if g.Plot[i][j] != nil {
				g.Plot[i][j].TimePlanted += 1

				// Crop is Fully Grown if TimePlanted exceeds production days
				if g.Plot[i][j].TimePlanted >= g.Plot[i][j].ProductionDays {
					g.Plot[i][j].FullyGrown = true

					// If crop has not been harvested within 5 days since it became
					// fully grown it becomes rotten
					if g.Plot[i][j].TimePlanted-g.Plot[i][j].ProductionDays > 5 {
						g.Plot[i][j].Rotten = true
					}
				}
			}
		}
	}
}

// AUTOMATICALLY GROWS PLAYER PLOT IF THEY HAVE ENOUGH POINTS
func (p *Player) updatePlot() {
	// Plot grows when players obtain 200 points, maximun plot size is 10x10
	if p.Points == 200 && p.Plot.PlotLevel == 0 {
		p.GrowPlotPlayer(2, 2)
		p.Plot.PlotLevel++
		fmt.Println("Your plot was automatically upgraded!")
	}
	// only upgrades the plot when player reaches the specified points
	// and when plot hasn't been updated yet
	if p.Points == 400 && p.Plot.PlotLevel == 1 {
		p.GrowPlotPlayer(3, 3)
		p.Plot.PlotLevel++
		fmt.Println("Your plot was automatically upgraded!")
	}
	if p.Points == 600 && p.Plot.PlotLevel == 2 {
		p.GrowPlotPlayer(4, 4)
		p.Plot.PlotLevel++
		fmt.Println("Your plot was automatically upgraded!")
	}
}
