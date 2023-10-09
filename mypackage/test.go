package mypackage

import (
	"fmt"
	"sort"
)

func SendAntsOpti(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {
	paths := SelectPaths(FindAndPrintPaths(roomsMap, startRoom, endRoom))

	if len(paths) == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}

	antPositions := make([]int, numAnts)
	occupiedRooms := make([]map[string]bool, numAnts)
	occupiedTunnels := make(map[string]bool)

	for i := range occupiedRooms {
		occupiedRooms[i] = make(map[string]bool)
	}

	// Indice de la dernière fourmi
	lastAntIndex := numAnts - 1

	for {
		allReachedEnd := true

		// Trier les chemins en fonction de leur longueur, en privilégiant le chemin le plus court pour la dernière fourmi
		sort.Slice(paths, func(i, j int) bool {
			if j == lastAntIndex {
				return len(paths[i]) <= len(paths[j])
			}
			return len(paths[i]) < len(paths[j])
		})

		for i := 0; i < numAnts; i++ {
			if antPositions[i] < len(paths[i%len(paths)])-1 {
				currentRoom := paths[i%len(paths)][antPositions[i]]
				nextRoom := paths[i%len(paths)][antPositions[i]+1]
				tunnel := currentRoom + "-" + nextRoom

				// Vérifier si le tunnel et la pièce sont disponibles
				if !occupiedTunnels[tunnel] && !occupiedRooms[i][nextRoom] {
					fmt.Printf("L%d-%s ", i+1, nextRoom)
					antPositions[i]++
					occupiedTunnels[tunnel] = true
					occupiedRooms[i][nextRoom] = true
				}

				allReachedEnd = allReachedEnd && antPositions[i] == len(paths[i%len(paths)])-1
			}
		}

		fmt.Println()

		// Vérification pour terminer la boucle
		if allReachedEnd {
			break
		}

		// Réinitialiser les tunnels occupés à la fin de chaque tour
		for tunnel := range occupiedTunnels {
			occupiedTunnels[tunnel] = false
		}
	}
}
