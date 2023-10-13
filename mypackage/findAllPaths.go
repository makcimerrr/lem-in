package mypackage

func FindPaths(roomsMap map[string]Room, currentRoom, endRoom string, path []string, visited map[string]bool) [][]string {
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
			subPaths := FindPaths(roomsMap, nextRoom, endRoom, newPath, visited) // Correction ici
			paths = append(paths, subPaths...)
		}
	}

	delete(visited, currentRoom)
	return paths
}
