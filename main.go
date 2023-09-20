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
	antsReachedEnd := make([]bool, numAnts)

	for {
		allReachedEnd := true
		usedRooms := make(map[string]bool)

		var antMovements [][]string

		for i := 0; i < numAnts; i++ {
			if !antsReachedEnd[i] {
				if antPositions[i] < len(paths[i%numPaths])-1 {
					room := paths[i%numPaths][antPositions[i]+1] // Commence à la deuxième pièce
					if !usedRooms[room] || (room == endRoom && !antsReachedEnd[i]) {
						antMovements = append(antMovements, []string{fmt.Sprintf("L%d-%s", i+1, room)})
						antPositions[i]++
						usedRooms[room] = true

						if room == endRoom {
							antsReachedEnd[i] = true
						}
					}
				} else {
					antsReachedEnd[i] = true
				}
				allReachedEnd = allReachedEnd && antsReachedEnd[i]
			}
		}

		if len(antMovements) > 0 {
			for i := 0; i < len(antMovements[0]); i++ {
				for j := 0; j < len(antMovements); j++ {
					if i < len(antMovements[j]) {
						fmt.Printf("%s ", antMovements[j][i])
					}
				}
				fmt.Println()
			}
		}

		if allReachedEnd {
			break
		}
	}
}
