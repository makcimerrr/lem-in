package mypackage

import (
	"sort"
	"strings"
)

func Filter(chemins [][]string) [][]string {
	uniquePaths := removeDuplicates(chemins)

	// Trier les chemins par longueur croissante
	sort.Slice(uniquePaths, func(i, j int) bool {
		return len(uniquePaths[i]) < len(uniquePaths[j])
	})

	// Filtrer les chemins pour s'assurer qu'ils ont au moins 2 pièces
	validPaths := make([][]string, 0)
	for _, path := range uniquePaths {
		if len(path) >= 2 {
			validPaths = append(validPaths, path)
		}
	}

	// Déterminer le nombre de chemins à sélectionner en fonction du nombre total
	numPathsToSelect := 2
	if len(validPaths) >= 9 {
		numPathsToSelect = 3
	} else if len(validPaths) == 3 {
		numPathsToSelect = 1
	} else if len(validPaths) == 1 {
		numPathsToSelect = 1
	}

	// Si nous n'avons pas suffisamment de chemins, essaie avec d'autres chemins
	for i := 0; i < len(validPaths); i++ {
		// Sélectionner le premier chemin
		selectedPaths := []string{strings.Join(validPaths[i], " ")}

		// Collecter les pièces utilisées dans le premier chemin
		usedPieces := make(map[string]bool)
		for j := 1; j < len(validPaths[i])-1; j++ {
			usedPieces[validPaths[i][j]] = true
		}

		// Essayer de trouver les autres chemins avec des pièces non utilisées
		for j := i + 1; j < len(validPaths); j++ {
			if len(selectedPaths) >= numPathsToSelect {
				break
			}

			hasUniquePieces := true
			for k := 1; k < len(validPaths[j])-1; k++ {
				if usedPieces[validPaths[j][k]] {
					hasUniquePieces = false
					break
				}
			}

			if hasUniquePieces {
				selectedPaths = append(selectedPaths, strings.Join(validPaths[j], " "))
				for k := 1; k < len(validPaths[j])-1; k++ {
					usedPieces[validPaths[j][k]] = true
				}
			}
		}

		if len(selectedPaths) == numPathsToSelect {
			result := make([][]string, numPathsToSelect)
			for j, path := range selectedPaths {
				result[j] = strings.Fields(path)
			}
			return result
		}
	}

	// Si on ne trouve pas le nombre requis de chemins avec des pièces différentes, retourner une liste vide
	return [][]string{}
}

// Fonction pour supprimer les doublons dans les chemins
func removeDuplicates(chemins [][]string) [][]string {
	uniquePaths := make([][]string, 0)
	seen := make(map[string]bool)

	for _, path := range chemins {
		key := strings.Join(path, " ")
		if !seen[key] {
			seen[key] = true
			uniquePaths = append(uniquePaths, path)
		}
	}

	return uniquePaths
}
