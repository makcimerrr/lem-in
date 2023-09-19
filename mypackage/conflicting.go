package mypackage

// FilterConflictingPaths filtre les chemins conflictuels pour obtenir uniquement des chemins uniques les plus courts.
func FilterConflictingPaths(paths [][]string) [][]string {
	uniquePaths := make([][]string, 0)
	visitedRooms := make(map[string]bool)

	for _, path := range paths {
		isUnique := true
		for _, room := range path {
			if visitedRooms[room] {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniquePaths = append(uniquePaths, path)
			// Marquer les pièces visitées comme occupées
			for _, room := range path {
				visitedRooms[room] = true
			}
		}
	}

	return uniquePaths
}

func FindPossiblePaths(roomsMap map[string]Room, start, end string) [][]string {
	visited := make(map[string]bool)
	paths := findPossiblePathsRecursive(roomsMap, start, end, []string{start}, visited)
	return paths
}

func findPossiblePathsRecursive(roomsMap map[string]Room, currentRoom, endRoom string, path []string, visited map[string]bool) [][]string {
	if currentRoom == endRoom {
		// La pièce de fin a été atteinte, ajoute le chemin à la liste des chemins
		return [][]string{append([]string{}, path...)}
	}

	visited[currentRoom] = true
	var paths [][]string

	for _, nextRoom := range roomsMap[currentRoom].Adjacent {
		if !visited[nextRoom] {
			// Évite les boucles infinies en vérifiant si la pièce n'a pas déjà été visitée
			newPath := append(path, nextRoom)
			subPaths := findPossiblePathsRecursive(roomsMap, nextRoom, endRoom, newPath, visited)
			paths = append(paths, subPaths...)
		}
	}

	delete(visited, currentRoom)
	return paths
}
