package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"package/mypackage"
)

type Room struct {
	Name     string
	Adjacent []string
}

var rooms map[string]Room

var (
	ant   int
	start int
	end   int
)

func main() {
	ant, start, end = mypackage.Checktxt() // Appel de la fonction pour obtenir ant, start et end

	rooms = make(map[string]Room)

	file := os.Args[1]
	ParseFile(file)

	paths := findPaths(fmt.Sprintf("%d", start), fmt.Sprintf("%d", end), []string{fmt.Sprintf("%d", start)}, make(map[string]bool))

	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	}
}

func findPaths(currentRoom, endRoom string, path []string, visited map[string]bool) [][]string {
	if currentRoom == endRoom {
		// La pièce de fin a été atteinte, ajoute le chemin à la liste des chemins
		return [][]string{append([]string{}, path...)}
	}

	visited[currentRoom] = true
	var paths [][]string

	for _, nextRoom := range rooms[currentRoom].Adjacent {
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

func parseTunnel(line string) (string, string) {
	parts := strings.Fields(line)
	rooms := strings.Split(parts[0], "-")
	return rooms[0], rooms[1]
}

func ParseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open the file !!")
		os.Exit(3)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // Ignorer les lignes de commentaire
		}
		if !strings.Contains(line, "-") {
			continue // Ignorer les lignes qui ne sont pas des tunnels
		}

		roomIn, roomOut := parseTunnel(line)
		rooms[roomIn] = Room{Name: roomIn, Adjacent: append(rooms[roomIn].Adjacent, roomOut)}
		rooms[roomOut] = Room{Name: roomOut, Adjacent: append(rooms[roomOut].Adjacent, roomIn)}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("cannot read the file !!")
		os.Exit(3)
	}
}
