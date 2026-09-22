package main

import "fmt"

var shopItems = []Item{itemDft4, itemDft5}

func Menumarchand(c *Character) {
	fmt.Println("************Welcome to the Merchant!************")
	fmt.Println("1. Acheter un objet")
	fmt.Println("2. Vendre un objet")
	fmt.Println("3. Retour au menu principal")

	var choix int
	fmt.Scanf("%d", &choix)
	switch {
	case choix == 1:
		fmt.Println("1. Acheter un objet")
		acheterobject()
		Menumarchand(c)
	case choix == 2:
		fmt.Println("2. Vendre un objet")
		Menumarchand(c)
	case choix == 3:
		fmt.Println("3. Retour au menu principal")
		Mainmenu(c)

	default:
		fmt.Println("Choix invalide, reessayez")
		Mainmenu(c)
	}
}

func acheterobject() {
	for i, item := range shopItems {
		fmt.Println(i+1, "-", item.Name, ":", item.description)
	}

}
