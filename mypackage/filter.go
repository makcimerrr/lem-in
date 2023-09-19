package mypackage

func filterUniquePaths(paths [][]string) [][]string {
	uniquePaths := [][]string{}
	visitedRooms := map[string]bool{}

	for _, path := range paths {
		isUnique := true
		for _, room := range path[1 : len(path)-1] { // Évite les pièces de départ et d'arrivée
			if visitedRooms[room] {
				isUnique = false
				break
			}
		}

		if isUnique {
			uniquePaths = append(uniquePaths, path)
			for _, room := range path[1 : len(path)-1] { // Évite les pièces de départ et d'arrivée
				visitedRooms[room] = true
			}
		}
	}

	return uniquePaths
}
