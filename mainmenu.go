package main

import (
	"fmt"
	"os"
)

func clearScreen() {
	fmt.Print("\033[3J\033[2J\033[H")
}

func Mainmenu(c *Character) {
	clearScreen()

	fmt.Println("************Welcome to the game!************")
	fmt.Println("----- Menu principal -----")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder à l'inventaire")
	fmt.Println("3. Marchand")
	fmt.Println("4. Forgeron")
	fmt.Println("5. Entrainement")
	fmt.Println("6. Quitter")
	fmt.Println("Votre choix :")

	var choix int
	fmt.Scan(&choix)
	switch {
	case choix == 1:
		clearScreen()
		fmt.Println("Informations du personnage")

		displayInfo(c)

	case choix == 2:
		clearScreen()
		fmt.Println("2. Accéder à l'inventaire")
	case choix == 3:
		clearScreen()
		fmt.Println("3. Marchand")
		Menumarchand(c)
	case choix == 4:
		clearScreen()
		fmt.Println("4. Forgeron")
	case choix == 5:
		clearScreen()
		fmt.Println("5. Entrainement")
	case choix == 6:
		clearScreen()
		fmt.Println("bye bye !! ╭∩╮( º.º )╭∩╮")
		os.Exit(0)
	default:
		clearScreen()
		fmt.Println("Choix invalide, reessayez")
		Mainmenu(c)

	}
}
