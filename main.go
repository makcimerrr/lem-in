package main

import (
	"fmt"
	"package/mypackage"
)

func main() {
	numberofant, start, end, rooms, links := mypackage.Checktxt()
	fmt.Printf("nombre de ants : %s\n", numberofant)
	fmt.Printf("début du circuit : %s\n", start)
	fmt.Printf("fin du circuit : %s\n", end)
	fmt.Println(rooms)
	fmt.Println(links)

}
