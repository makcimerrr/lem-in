package mypackage

import (
	"sort"
	"strings"
)

func FilterAndReturnShortestPaths(paths [][]string, numPaths int) [][]string {
	// Triez les chemins par longueur
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Filtrer les chemins non-conflictuels
	var shortestPaths [][]string
	seenRooms := make(map[string]bool)

	for _, path := range paths {
		conflict := false

		for _, room := range path[1 : len(path)-1] { // Ignorer la première et la dernière pièce
			if seenRooms[room] {
				conflict = true
				break
			}
		}

		if !conflict {
			shortestPaths = append(shortestPaths, path)
			for _, room := range path[1 : len(path)-1] { // Marquer les pièces comme vues
				seenRooms[room] = true
			}
		}

		if len(shortestPaths) == numPaths {
			break
		}
	}

	return shortestPaths
}

func arePiecesUnique(path1, path2 []string) bool {
	// Crée une carte pour stocker les pièces utilisées dans le chemin
	pieceMap := make(map[string]bool)

	// Remplit la carte avec les pièces du chemin 1
	for _, piece := range path1 {
		pieceMap[getPieceName(piece)] = true
	}

	// Vérifie si les pièces du chemin 2 sont déjà utilisées dans le chemin 1
	for _, piece := range path2 {
		if pieceMap[getPieceName(piece)] {
			return false
		}
	}

	return true
}

func getPieceName(piece string) string {
	parts := strings.Split(piece, "-")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}
