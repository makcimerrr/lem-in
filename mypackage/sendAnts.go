package mypackage

import (
	"fmt"
	"sort"
)

func SendAnts(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {

	allpaths := FindAndPrintPaths(roomsMap, start, end)
	paths := SelectPaths(allpaths)
	numPaths := len(paths)

	var instantEnd = true

	if numPaths == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}

	// Afficher la longueur de chaque chemin
	for nb, path := range paths {
		fmt.Println("longeur du", nb+1, "chemin ", len(path))
	}

	antPositions := make([]int, numAnts)
	antsReachedEnd := make([]bool, numAnts)

	// Utiliser un tableau pour suivre les pièces occupées par les fourmis
	occupiedRooms := make([]map[string]bool, numAnts)
	for i := 0; i < numAnts; i++ {
		occupiedRooms[i] = make(map[string]bool)
	}

	tunnelOccupied := make(map[string]bool)

	for {
		allReachedEnd := true

		// Trier les chemins en fonction de leur longueur
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) < len(paths[j])
		})

		for i := 0; i < numAnts; i++ {
			if !antsReachedEnd[i] {
				currentRoom := paths[i%numPaths][antPositions[i]]
				nextRoom := paths[i%numPaths][antPositions[i]+1]
				tunnel := currentRoom + "-" + nextRoom

				// Vérifier si le tunnel est disponible
				if !tunnelOccupied[tunnel] && !occupiedRooms[i][nextRoom] {
					fmt.Printf("L%d-%s ", i+1, nextRoom)
					antPositions[i]++
					tunnelOccupied[tunnel] = true
					occupiedRooms[i][nextRoom] = true

<<<<<<< HEAD
					room := paths[i%numPaths][antPositions[i]+1]

					if !usedRooms[room] || (room == endRoom) {
						/* fmt.Println(antPositions)
						fmt.Println(room)
						fmt.Println(instantEnd) */

						if antPositions[i] == 0 && room == "3" {
							if instantEnd == true {
								antMovements = append(antMovements, []string{fmt.Sprintf("L%d-%s", i+1, room)})
								usedRooms[room] = true
								antPositions[i]++

								if room == endRoom {
									antsReachedEnd[i] = true
								}
								instantEnd = false
							}
							

						} else {
							antMovements = append(antMovements, []string{fmt.Sprintf("L%d-%s", i+1, room)})
							usedRooms[room] = true
							antPositions[i]++

							if room == endRoom {
								antsReachedEnd[i] = true
							}
						}

=======
					if nextRoom == endRoom {
						antsReachedEnd[i] = true
>>>>>>> 288511c9aceace1661a7635e350ad6043bb7d921
					}
				}
				allReachedEnd = allReachedEnd && antsReachedEnd[i]
			}

		}

<<<<<<< HEAD
		instantEnd = true

		//faire l'affichage
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
=======
		fmt.Println()
>>>>>>> 288511c9aceace1661a7635e350ad6043bb7d921

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
