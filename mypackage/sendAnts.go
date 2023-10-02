package mypackage

import (
	"fmt"
	"sort"
)

func SendAnts(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {
	paths := FindUniquePaths(roomsMap, startRoom, endRoom)
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

	for {
		allReachedEnd := true
		usedRooms := make(map[string]bool)

		var antMovements [][]string

		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) < len(paths[j])
		})

		for i := 0; i < numAnts; i++ {
			if !antsReachedEnd[i] {

				if antPositions[i] < len(paths[i%numPaths])-1 {

					room := paths[i%numPaths][antPositions[i]+1]

					if !usedRooms[room] || (room == endRoom) {

						antMovements = append(antMovements, []string{fmt.Sprintf("L%d-%s", i+1, room)})
						usedRooms[room] = true
						antPositions[i]++

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
