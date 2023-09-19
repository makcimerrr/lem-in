package main

import (
	"fmt"
	"package/mypackage"
)

var (
	ant   string
	start string
	end   string
	links []string
	rooms []string
)

func main() {

	ant, start, end, rooms, links = mypackage.Checktxt() // Appel de la fonction pour obtenir les données

	fmt.Println("\tAnt(s) :", ant)

	// Lire les tunnels et construire la carte des pièces adjacentes
	roomsMap := mypackage.BuildRoomMap(rooms, links)

	// Trouver et afficher tous les chemins possibles
	mypackage.FindAndPrintPaths(roomsMap, start, end)

}
