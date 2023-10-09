package mypackage

import (
	"sort"
)

func SelectPaths(chemins [][]string) [][]string {
	// Trier les chemins par longueur
	sort.Slice(chemins, func(i, j int) bool {
		return len(chemins[i]) < len(chemins[j])
	})

	// Vérifier si tous les chemins ont la même taille
	allSameSize := true
	for i := 1; i < len(chemins); i++ {
		if len(chemins[i]) != len(chemins[0]) {
			allSameSize = false
			break
		}
	}

	// Sélectionner les chemins en fonction du nombre de chemins total et de la taille des chemins
	var numPathsToSelect int
	if len(chemins) >= 9 {
		numPathsToSelect = 3
	} else {
		numPathsToSelect = 2
	}
	if len(chemins) >= 30 {
		numPathsToSelect = 4
	}

	// Si tous les chemins ont la même taille, n'afficher qu'un chemin
	if allSameSize {
		numPathsToSelect = 1
	}

	// Sélectionner les chemins tout en vérifiant que les pièces enfants des pièces "start" ne sont pas identiques
	selectedPaths := make([][]string, 0)
	childPieces := make(map[string]bool)
	for _, chemin := range chemins {
		if len(selectedPaths) >= numPathsToSelect {
			break
		}

		startChildren := make([]string, 0)
		for i := 1; i < len(chemin)-1; i++ {
			if chemin[i-1] == "start" {
				startChildren = append(startChildren, chemin[i])
			}
		}

		isUnique := true
		for _, piece := range startChildren {
			if childPieces[piece] {
				isUnique = false
				break
			}
		}

		if isUnique {
			for _, piece := range startChildren {
				childPieces[piece] = true
			}

			selectedPaths = append(selectedPaths, chemin)
		}
	}

	return selectedPaths
}
