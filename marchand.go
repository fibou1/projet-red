package main

import "fmt"

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
	fmt.Println("list des items disponibles à l'achat :")
	fmt.Printf("1. %v\n2. %v\n3. %v\n4. %v\n5. %v\n", itemDft1.Name, itemDft2.Name, itemDft3.Name, itemDft4.Name, itemDft5.Name)
	var choix int
	fmt.Scanf("%d", &choix)

}
