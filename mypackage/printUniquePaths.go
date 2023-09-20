package mypackage

import (
	"fmt"
	"strings"
)

func FindAndPrintUniquePaths(roomsMap map[string]Room, start, end string) {
	uniquePaths := FindUniquePaths(roomsMap, start, end)
	PrintPaths(uniquePaths)
}

func PrintPaths(paths [][]string) {
	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	} //Sort tous les chemins possible non-conflictuels, trié après avoir sorti tous les chemins.
}
