package mypackage

import "sort"

func FilterAndReturnShortestPaths(paths [][]string, maxPaths, count int) [][]string {
	var validPaths [][]string

	// Ajouter le premier chemin
	validPaths = append(validPaths, paths[0])

	// Parcourir les chemins et sélectionner ceux qui n'ont pas de pièces en commun
	for i := 1; i < len(paths) && len(validPaths) < maxPaths; i++ {
		currentPath := paths[i]
		isUnique := true

		// Vérifier si le chemin a la même longueur que le dernier chemin ajouté
		if len(currentPath) == len(validPaths[0]) {
			// Vérifier si les chemins partagent une pièce en commun
			for _, room := range currentPath {
				if contain(validPaths[0], room) {
					isUnique = false
					break
				}
			}
		}

		if isUnique {
			validPaths = append(validPaths, currentPath)
		}
	}

	// Trier les chemins par longueur
	sort.Slice(validPaths, func(i, j int) bool {
		return len(validPaths[i]) < len(validPaths[j])
	})

	// Afficher soit 2 ou 3 chemins en fonction de count et maxPaths
	if count >= maxPaths {
		return validPaths[:3]
	}

	return validPaths[:2]
}

// Fonction utilitaire pour vérifier si une slice contient une chaîne donnée
func contain(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
