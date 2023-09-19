package mypackage

import (
	"fmt"
)

// Ant représente une fourmi.
type Ant struct {
	ID     int
	Room   string
	Path   []string
	Active bool
}

// SimulateAnts simule le mouvement des fourmis sur les chemins.
func SimulateAnts(roomsMap map[string]Room, ants int, start, end string, paths [][]string) {
	antPositions := make([]int, ants)    // Stocke la position actuelle de chaque fourmi
	antLocations := make([]string, ants) // Stocke la pièce actuelle de chaque fourmi

	adjacentRooms := roomsMap[start].Adjacent // Pièces adjacentes à la pièce de départ

	if len(adjacentRooms) == 0 {
		fmt.Println("Il y a aucunes pièces adjacentes.")
	}

	// Initialisation des positions des fourmis dans les pièces adjacentes
	for antID := 0; antID < ants && antID < len(adjacentRooms); antID++ {
		antLocations[antID] = adjacentRooms[antID]
		fmt.Printf("L%d-%s ", antID+1, antLocations[antID])
		antPositions[antID]++
	}
	fmt.Println()

	// Simuler le mouvement des fourmis
	for !allAntsReachedEnd(antPositions, end) {
		for antID := 0; antID < ants; antID++ {
			if antPositions[antID] < len(paths[antID]) {
				nextRoom := paths[antID][antPositions[antID]]
				antPositions[antID]++
				antLocations[antID] = nextRoom
				fmt.Printf("L%d-%s ", antID+1, antLocations[antID])
			}
		}
		fmt.Println()
	}
}

func allAntsReachedEnd(antPositions []int, end string) bool {
	for _, position := range antPositions {
		if position < len(end) && end[position] != end[0] {
			return false
		}
	}
	return true
}
