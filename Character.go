package main

import "fmt"

type Item struct {
	ID          int
	Name        string
	description string
	Quantity    int
}

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory []Item
}

func initCharacter(name string, class string, maxHP int, inventory []Item) *Character {
	c := &Character{
		Name:      name,
		Class:     class,
		Level:     1,
		MaxHP:     maxHP,
		CurrentHP: maxHP / 2,
		Inventory: inventory,
	}
	return c
}

func askclass() (string, int) {
	var choix int
	fmt.Println("Choose your class :")
	fmt.Println("1. alien (100 PV)")
	fmt.Println("2. humain (80 PV)")
	fmt.Println("3. reptile (120 PV)")
	fmt.Println("Votre choix :")
	fmt.Scanln(&choix)
	if choix == 1 {
		return "alien", 100
	} else if choix == 2 {
		return "humain", 80
	} else if choix == 3 {
		return "reptile", 120
	} else {
		fmt.Println("Choix invalide, reessayez")
	}
	return askclass()
}

func createCharacter() *Character {
	class, maxHP := askclass()
	name := askname()

	c := initCharacter(name, class, maxHP, []Item{itemDft1, itemDft2, itemDft3})
	return c

}

func askname() string {
	var name string
	fmt.Println("Enter your character's name:")
	fmt.Scanln(&name)
	return name
}

func displayInfo(c *Character) {
	fmt.Println("Character Details:")
	fmt.Println("Name:", c.Name)
	fmt.Println("Class:", c.Class)
	fmt.Println("Level:", c.Level)
	fmt.Println("Max HP:", c.MaxHP)
	fmt.Println("Current HP:", c.CurrentHP)
	fmt.Println("Inventory:")
	for _, item := range c.Inventory {
		fmt.Printf(" : %s  Quantity: %d\n", item.Name, item.Quantity)
	}
}
