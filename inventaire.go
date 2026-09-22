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
func Additem(c *Character, newItem Item) bool {
	if c.MaxSlots > len(c.Inventory) {
		fmt.Println("Inventory is full. Cannot add more items.")
		return false
	} else {
		fmt.Println("Adding item to inventory:", newItem.Name)
		c.Inventory = append(c.Inventory, newItem)
		return true
	}

}
func RemoveItem(c *Character, item Item) bool {
	for i, invItem := range c.Inventory {
		if invItem.ID == item.ID {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Println("Removed item from inventory:", item.Name)
			return true
		}
	}
	fmt.Println("Item not found in inventory.")
	return false
}
func takePot(c *Character) {
	for i, item := range c.Inventory {
		if item.ID == 1 && item.Quantity > 0 {
			if c.MaxHP > c.CurrentHP {
				if c.MaxHP-c.CurrentHP <= 50 {
					c.CurrentHP = c.MaxHP
					c.Inventory[i].Quantity -= 1
					fmt.Println("Your HP is now full ! you have used x1 potion.")
					fmt.Printf(" Your current HP is now: %d/%d\n", c.CurrentHP, c.MaxHP)

					return
				} else {
					c.CurrentHP += 50
					c.Inventory[i].Quantity -= 1
					fmt.Println("Your HP is now +50  ! you have used a potion.")
					fmt.Printf(" Your current HP is now: %d/%d\n", c.CurrentHP, c.MaxHP)
					return
				}
			} else {
				fmt.Println("Your HP is already full. You cannot use a potion.")
				return
			}
		}
	}

}
