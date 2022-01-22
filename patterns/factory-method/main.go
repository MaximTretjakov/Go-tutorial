package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	mushket, _ := getGun("mushket")

	printDetails(ak47)
	printDetails(mushket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
