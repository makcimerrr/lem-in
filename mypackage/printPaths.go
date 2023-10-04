package mypackage

func FindAndPrintPaths(roomsMap map[string]Room, start, end string) [][]string {
	// Trouver tous les chemins possibles de la pièce de départ à la pièce de fin
	paths := FindPaths(roomsMap, start, end, []string{start}, make(map[string]bool))

	return paths //Sort tous les chemins possible de la fourmillière.
}
