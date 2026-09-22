package main

import "fmt"

var itemDft1 = Item{ID: 1, Name: "Potion", description: "Restores 50 HP", Quantity: 1}
var itemDft2 = Item{ID: 2, Name: "Sword", description: "A sharp blade", Quantity: 1}
var itemDft3 = Item{ID: 3, Name: "Shield", description: "Protects from attacks", Quantity: 1}
var itemDft4 = Item{ID: 4, Name: "Helmet", description: "Protects the head", Quantity: 1}
var itemDft5 = Item{ID: 5, Name: "superBoots", description: "Increases speed", Quantity: 1}

func DisplayInventory(c *Character) {
	fmt.Println("Inventory:")
	for _, item := range c.Inventory {
		fmt.Printf("ID: %d, Name: %s, Description: %s, Quantity: %d\n", item.ID, item.Name, item.description, item.Quantity)
	}
}

func Additem(c *Character, newItem Item) {
	length := len(c.Inventory)
	if length >= 10 {
		fmt.Println("Inventory is full. Cannot add more items.")
	} else {
		fmt.Println("Adding item to inventory:", newItem.Name)
		c.Inventory = append(c.Inventory, newItem)
	}

}
