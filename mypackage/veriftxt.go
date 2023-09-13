package mypackage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ant int

func Checktxt() int {
	//ant := false
	//start := false
	//end := false

	if len(os.Args) == 2 { //verif si il y bien 2 arg donc main + l'example sinon message err
		fichier := os.Args[1]
		file, err := os.Open(fichier)
		if err != nil {
			fmt.Println("cannot open the file !!")
			os.Exit(3)
		}
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)

		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		file.Close()

		for _, line := range fileLines { // ici verif du txt
			fmt.Println(line)
			intfirstline, err := strconv.Atoi(fileLines[0])

			if err != nil {
				//fmt.Print("impossible de convertir la premiere ligne en int ")
				os.Exit(3)

			}

			if intfirstline > 0 {
				//ant = true
				ant = intfirstline
			} else {
				fmt.Println("le nombre de ants doit etre suprieur a 0 ")
				os.Exit(3)

			}
		}

	} else {
		fmt.Println("Besoin d'un fichier txt pour executer le projet !!!")
		os.Exit(3)
	}
	return ant
}
