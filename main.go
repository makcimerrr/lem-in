package main

import (
	"fmt"
	"log"
	"package/mypackage"
	"strconv"
	"strings"
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
	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	}

	fmt.Println("====== AVEC TRIS ======")

	sortedpaths := mypackage.FilterAndReturnShortestPaths(paths, 3)

	// Afficher les chemins triés
	for _, path := range sortedpaths {
		fmt.Println(strings.Join(path, " "))
	}

	fmt.Println("====== SIMULATION ======")
	// Vous pouvez appeler sendAnts avec les données appropriées
	mypackage.SendAnts(roomsMap, start, end, ants)

}
