package mypackage

func FindUniquePaths(roomsMap map[string]Room, startRoom, endRoom string) [][]string {
	uniquePaths := [][]string{}
	currentPath := []string{start}

	Paths(roomsMap, start, end, currentPath, &uniquePaths, map[string]bool{})

	uniquePaths = filterUniquePaths(uniquePaths)

	return uniquePaths
}
