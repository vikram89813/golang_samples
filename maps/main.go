package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// fmt.Println(colors)

	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	colors["white"] = "#ffffff"

	// delete(colors, "green")

	//fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
