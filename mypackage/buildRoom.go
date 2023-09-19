package mypackage

import "strings"

func BuildRoomMap(rooms []string, links []string) map[string]Room {
	roomsMap := make(map[string]Room)

	// Construire la carte des pièces adjacentes
	for _, link := range links {
		roomIn, roomOut := parseTunnel(rooms, link)
		roomsMap[roomIn] = Room{Name: roomIn, Adjacent: append(roomsMap[roomIn].Adjacent, roomOut)}
		roomsMap[roomOut] = Room{Name: roomOut, Adjacent: append(roomsMap[roomOut].Adjacent, roomIn)}
	}

	return roomsMap
}

func parseTunnel(roomsParam []string, line string) (string, string) {
	parts := strings.Fields(line)
	roomsNames := strings.Split(parts[0], "-")
	return roomsNames[0], roomsNames[1]
}
