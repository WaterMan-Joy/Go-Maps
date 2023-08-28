package main

import (
	"fmt"
)

func main() {
	var colors = map[string]string{
		// "red":   "#ff0000",
		// "green": "#ff00ff",
	}

	colors["blue"] = "blue"
	colors["y"] = "yellow"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
