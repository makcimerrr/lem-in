package mypackage

import (
	"fmt"
	"strings"
)

func FindAndPrintPaths(roomsMap map[string]Room, start, end string) [][]string {
	// Trouver tous les chemins possibles de la pièce de départ à la pièce de fin
	paths := findPaths(roomsMap, start, end, []string{start}, make(map[string]bool))

	// Afficher les chemins trouvés
	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	}
	return paths
}
