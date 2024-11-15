package main

type Crop struct {
	Name       string
	Symbol     string
	fullyGrown bool
}

// go run .
func main() {
	garden := CreatePlot(5, 5)
	potato := Crop{Name: "Flower", Symbol: "🥔", fullyGrown: true}
	mango := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: false}
	mango2 := Crop{Name: "Mango", Symbol: "🥭", fullyGrown: true}
	garden.Plant(0, 0, &potato)
	garden.Plant(2, 2, &mango)
	garden.Plant(2, 3, &mango2)
	garden.printGarden()
}
