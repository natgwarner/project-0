package managers

import "fmt"

//Inventory is a slice that stores the items within the player's inventory.
var Inventory []string

//InventoryCheck allows the player to check their inventory.
func InventoryCheck() {
	if len(Inventory) >= 1 {
		fmt.Println("You search your backpack and find:")
		for i := 0; i < len(Inventory); i++ {
			fmt.Println(Inventory[i] + "\n")
		}
	} else {
		fmt.Println("You search your backpack, but find nothing.")
	}
}

//InventoryAdd allows to the program to add items to the player's inventory.
func InventoryAdd(s string) {
	Inventory = append(Inventory, s)
}
