package mypackage

import (
	"fmt"
	"sort"
)

func SendAntsOpti(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {
	allpaths := FindAndPrintPaths(roomsMap, start, end)
	uniquePaths := SelectPaths(allpaths)
	numPaths := len(uniquePaths)

	sort.Slice(uniquePaths, func(i, j int) bool {
		return len(uniquePaths[i]) > len(uniquePaths[j]) // Triez à l'envers pour obtenir le plus court en dernier
	})

	if numPaths == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}

	// Afficher la longueur de chaque chemin
	for _, path := range uniquePaths {
		fmt.Println(len(path))
	}

	antPositions := make([]int, numAnts)
	antsReachedEnd := make([]bool, numAnts)

	// Utiliser une carte pour suivre les pièces occupées par les fourmis
	occupiedRooms := make([]map[string]bool, numAnts)
	for i := 0; i < numAnts; i++ {
		occupiedRooms[i] = make(map[string]bool)
	}

	tunnelOccupied := make(map[string]bool)

	for {
		allReachedEnd := true

		for i := 0; i < numAnts; i++ {
			if !antsReachedEnd[i] {
				currentRoom := uniquePaths[i%numPaths][antPositions[i]]
				nextRoom := uniquePaths[i%numPaths][antPositions[i]+1]
				tunnel := currentRoom + "-" + nextRoom

				// Vérifier si le tunnel est disponible
				if !tunnelOccupied[tunnel] && !occupiedRooms[i][nextRoom] {
					fmt.Printf("L%d-%s ", i+1, nextRoom)
					antPositions[i]++
					tunnelOccupied[tunnel] = true
					occupiedRooms[i][nextRoom] = true

					if nextRoom == endRoom {
						antsReachedEnd[i] = true
					}
				}
				allReachedEnd = allReachedEnd && antsReachedEnd[i]
			}
		}

		fmt.Println()

		// Vérification pour terminer la boucle
		if allReachedEnd {
			break
		}

		// Si tous les tunnels sont occupés, libérer les tunnels
		for tunnel := range tunnelOccupied {
			tunnelOccupied[tunnel] = false
		}
	}
}
