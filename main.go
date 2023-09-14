package main

import (
	"fmt"
	"package/mypackage"
	"strings"
)

type Room struct {
	Name     string
	Adjacent []string
}

var roomsMap map[string]Room

var (
	ant   string
	start string
	end   string
	links []string
	rooms []string
)

func main() {
	ant, start, end, rooms, links = mypackage.Checktxt() // Appel de la fonction pour obtenir les données

	// Lire les tunnels et construire la carte des pièces adjacentes
	buildRoomMap(rooms, links)

	// Trouver et afficher tous les chemins possibles
	FindAndPrintPaths(start, end)
}

func buildRoomMap(rooms []string, links []string) {

	roomsMap = make(map[string]Room)

	// Construire la carte des pièces adjacentes
	for _, link := range links {
		roomIn, roomOut := parseTunnel(rooms, link)
		roomsMap[roomIn] = Room{Name: roomIn, Adjacent: append(roomsMap[roomIn].Adjacent, roomOut)}
		roomsMap[roomOut] = Room{Name: roomOut, Adjacent: append(roomsMap[roomOut].Adjacent, roomIn)}
	}
}

func parseTunnel(roomsParam []string, line string) (string, string) {
	parts := strings.Fields(line)
	roomsNames := strings.Split(parts[0], "-")
	return roomsNames[0], roomsNames[1]
}

func findPaths(currentRoom, endRoom string, path []string, visited map[string]bool) [][]string {
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
			subPaths := findPaths(nextRoom, endRoom, newPath, visited)
			paths = append(paths, subPaths...)
		}
	}

	delete(visited, currentRoom)
	return paths
}

func FindAndPrintPaths(start, end string) {
	// Trouver tous les chemins possibles de la pièce de départ à la pièce de fin
	paths := findPaths(start, end, []string{start}, make(map[string]bool))

	// Afficher les chemins trouvés
	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	}
}
