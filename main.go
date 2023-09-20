package main

import (
	"fmt"
	"log"
	"package/mypackage"
	"strconv"
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

	// Trouver et afficher tous les chemins possibles
	mypackage.FindAndPrintPaths(roomsMap, start, end)

	fmt.Println("====== AVEC TRIS ======")

	// Trouver et afficher les chemins uniques les plus courts
	mypackage.FindAndPrintUniquePaths(roomsMap, start, end)
}
