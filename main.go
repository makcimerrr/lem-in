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

	fmt.Println("====== SIMULATION ======")
	// Vous pouvez appeler sendAnts avec les données appropriées
	sendAnts(roomsMap, start, end, ants)

}

func sendAnts(roomsMap map[string]mypackage.Room, startRoom, endRoom string, numAnts int) {
	paths := mypackage.FindUniquePaths(roomsMap, startRoom, endRoom)
	numPaths := len(paths)

	if numPaths == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}

	antPositions := make([]int, numAnts)

	for i := range antPositions {
		antPositions[i] = 1 // Ignorer la première étape pour chaque fourmi
	}

	for {
		for i := 0; i < numAnts; i++ {
			if antPositions[i] < len(paths[i%numPaths])-1 {
				fmt.Printf("L%d-%s ", i+1, paths[i%numPaths][antPositions[i]])
				antPositions[i]++
			} else {
				fmt.Printf("L%d-0 ", i+1)
			}
		}
		fmt.Println()

		allReachedEnd := true
		for i := 0; i < numAnts; i++ {
			if antPositions[i] < len(paths[i%numPaths])-1 {
				allReachedEnd = false
				break
			}
		}

		if allReachedEnd {
			break
		}
	}
}
