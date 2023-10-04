package mypackage

import (
	"fmt"
	"sort"
)

func SendAnts(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {
	allpaths := FindAndPrintPaths(roomsMap, start, end)
	paths := FilterAndReturnShortestPaths(allpaths, 3)
	numPaths := len(paths)

	if numPaths == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}
	// Afficher la longueur de chaque chemin
	for _, path := range paths {
		fmt.Println(len(path))
	}

	antPositions := make([]int, numAnts)
	antsReachedEnd := make([]bool, numAnts)

	tunnelOccupied := make(map[string]bool)

	for {
		allReachedEnd := true

		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) < len(paths[j])
		})

		for i := 0; i < numAnts; i++ {
			if !antsReachedEnd[i] {
				if antPositions[i] < len(paths[i%numPaths])-1 {
					room := paths[i%numPaths][antPositions[i]+1]
					tunnel := paths[i%numPaths][antPositions[i]] + "-" + room

					if !tunnelOccupied[tunnel] {
						fmt.Printf("L%d-%s ", i+1, room)
						antPositions[i]++
						tunnelOccupied[tunnel] = true

						if room == endRoom {
							antsReachedEnd[i] = true
						}
					}
				} else {
					antsReachedEnd[i] = true
				}
			}
			allReachedEnd = allReachedEnd && antsReachedEnd[i]
		}

		fmt.Println()

		// Vérification pour terminer la boucle
		if allReachedEnd {
			break
		}

		// Si tous les tunnels sont occupés, libérer un tunnel
		for tunnel := range tunnelOccupied {
			tunnelOccupied[tunnel] = false
		}
	}
}
