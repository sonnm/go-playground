package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["red"] = "#ff0000"

	// delete(colors, "red")

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("%s : %s\n", color, hex)
	}
}
