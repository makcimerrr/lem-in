package mypackage

func Paths(roomsMap map[string]Room, currentRoom, endRoom string, currentPath []string, uniquePaths *[][]string, visited map[string]bool) {
	if currentRoom == endRoom {
		// La pièce de fin a été atteinte, ajoute le chemin à la liste des chemins uniques
		*uniquePaths = append(*uniquePaths, append([]string{}, currentPath...))
		return
	}

	visited[currentRoom] = true

	for _, nextRoom := range roomsMap[currentRoom].Adjacent {
		if !visited[nextRoom] && !contains(currentPath, nextRoom) {
			// Évite les boucles infinies en vérifiant si la pièce n'a pas déjà été visitée
			newPath := append(currentPath, nextRoom)
			Paths(roomsMap, nextRoom, endRoom, newPath, uniquePaths, visited)
		}
	}

	delete(visited, currentRoom)
}
