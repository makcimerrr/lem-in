package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"package/mypackage"
)

type Room struct {
	Name     string
	Adjacent []string
}

var (
	ant   string
	start string
	end   string
	links []string
	rooms []string
)

func main() {
	ant, start, end, rooms, links = mypackage.Checktxt() // Appel de la fonction pour obtenir les donnée

	ants, err := strconv.Atoi(ant)

	if ants != 1 {
		fmt.Println("\tAnts :", ant)
	} else {
		fmt.Println("\tAnt :", ant)
	}

	if err != nil {
		log.Println("Error")
	}

	// Construire la carte des pièces adjacentes
	roomsMap := mypackage.BuildRoomMap(rooms, links)

	fmt.Println("====== ALL PATHS ======")
	// Trouver et afficher tous les chemins possibles
	paths := mypackage.FindAndPrintPaths(roomsMap, start, end)

	// Afficher les chemins trouvés
	count := 0
	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
		count++
	}

	fmt.Println(count)

	fmt.Println("====== AVEC TRIS ======")

	// Appeler la fonction pour sélectionner les chemins
	selectedPaths := mypackage.SelectPaths(paths)

	// Afficher les chemins sélectionnés
	fmt.Println("Chemins sélectionnés :")
	for _, chemin := range selectedPaths {
		fmt.Println(chemin)
	}

	fmt.Println("====== AVEC TRIS OPTI ======")

	// Appeler la fonction pour sélectionner les chemins
	selectedPaths2 := mypackage.Filter(paths)

	// Afficher les chemins sélectionnés
	fmt.Println("Chemins sélectionnés :")
	for _, chemin := range selectedPaths2 {
		fmt.Println(chemin)
	}

	fmt.Println("====== SIMULATION ======")
	// Vous pouvez appeler sendAnts avec les données appropriées
	mypackage.SendAnts(roomsMap, start, end, ants)

	fmt.Println("====== SIMULATION OPTI ======")
	// Vous pouvez appeler sendAnts avec les données appropriées
	mypackage.SendAntsOpti(roomsMap, start, end, ants)
}
