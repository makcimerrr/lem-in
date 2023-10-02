package mypackage

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SendAnts(roomsMap map[string]Room, startRoom, endRoom string, numAnts int) {
	paths := FindUniquePaths(roomsMap, startRoom, endRoom)
	numPaths := len(paths)
	var flag bool

	if numPaths == 0 {
		fmt.Println("Pas de chemin disponible.")
		return
	}
	// Afficher la longueur de chaque chemin
	for _, path := range paths {
		fmt.Println(len(path))
		flag = true
	}

	antPositions := make([]int, numAnts)
	antsReachedEnd := make([]bool, numAnts)
	if !flag {
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
	} else {
		// Chemin du fichier texte
		filePath := "e.txt"

		// Ouverture du fichier
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Impossible d'ouvrir le fichier : %v\n", err)
			return
		}
		defer file.Close()

		// Création d'un scanner pour lire le fichier ligne par ligne
		scanner := bufio.NewScanner(file)

		// Lecture et impression du contenu ligne par ligne
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
