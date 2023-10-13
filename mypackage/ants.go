package mypackage

import (
	"fmt"
)

func Ants(ant int) {
	if ant != 1 {
		fmt.Println("\tAnts :", ant)
	} else {
		fmt.Println("\tAnt :", ant)
	}
}
