package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Hex code for %v is %v\n", key, value)
	}
}
