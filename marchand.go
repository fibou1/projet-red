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
		Acheterobject(c)
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

func Acheterobject(c *Character) {
	fmt.Println("list des items disponibles à l'achat :")
	fmt.Printf("1. %v\n2. %v\n3. %v\n4. %v\n5. %v\n6. %v\n7. %v\n", itemDft1.Name, itemDft2.Name, itemDft3.Name, itemDft4.Name, itemDft5.Name, itemDft6.Name, itemDft7.Name)
	fmt.Println("chosir un item à acheter (1-7) :")
	fmt.Println("8 pour revenir au menu principal")

	var choix int
	fmt.Scanf("%d", &choix)
	switch {
	case choix == 1:
		if c.coin >= itemDft1.Coinprix {
			c.coin -= itemDft1.Coinprix
			Additem(c, itemDft1)

		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 2:
		if c.coin >= itemDft2.Coinprix {
			c.coin -= itemDft2.Coinprix
			Additem(c, itemDft2)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 3:
		if c.coin >= itemDft3.Coinprix {
			c.coin -= itemDft3.Coinprix
			Additem(c, itemDft3)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 4:
		if c.coin >= itemDft4.Coinprix {
			c.coin -= itemDft4.Coinprix
			Additem(c, itemDft4)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 5:
		if c.coin >= itemDft5.Coinprix {
			c.coin -= itemDft5.Coinprix
			Additem(c, itemDft5)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 6:
		if c.coin >= itemDft6.Coinprix {
			c.coin -= itemDft6.Coinprix
			Additem(c, itemDft6)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 7:
		if c.coin >= itemDft7.Coinprix {
			c.coin -= itemDft7.Coinprix
			Additem(c, itemDft7)
		} else {
			fmt.Println("Vous n'avez pas assez de coins pour acheter cet objet.")
		}
	case choix == 8:
		Mainmenu(c)

	default:
		fmt.Println("Choix invalide, reessayez")
		Acheterobject(c)
	}

}
