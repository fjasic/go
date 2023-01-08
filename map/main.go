package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#003399",
	}
	colors["green"] = "#231231"
	fmt.Println(colors)
	delete(colors, "green")
	fmt.Println(colors)
	iterate(colors)
}

func iterate(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
