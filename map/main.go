package main

import "fmt"

func main() {
	// declare a map

	// option1:
	// varName := map[keyType]valueType{values}
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	// option 2:
	// var name map[keyType]valueType
	// var colors map[string]string

	// option 3:
	// varName := make(map[keyType]valueType)
	// colors := make(map[string]string)

	// adding key-value pairs to a map
	colors["white"] = "#ffffff" //note - go maps don't have dot notation!

	// remove key-value pairs from map
	// delete(mapName, keyToDelete)
	delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code of %s is %s\n", color, hex)
	}
}
