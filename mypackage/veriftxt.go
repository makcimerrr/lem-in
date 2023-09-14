package mypackage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ant string
var start string
var end string
var links []string
var rooms []string
var temp []string

func Checktxt() (string, string, string, []string, []string) {

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

		intfirstline, err := strconv.Atoi(fileLines[0])
		if err != nil {
			fmt.Print("impossible de convertir la premiere ligne en int ")
			os.Exit(3)
		}
		if intfirstline > 0 {
			stringfirstline := strconv.Itoa(intfirstline)

			ant = stringfirstline
		} else {
			fmt.Println("le nombre de ants doit etre suprieur a 0 ")
			os.Exit(3)

		}

		for i := 1; i < len(fileLines); i++ {

			if fileLines[i-1] == "##start" {

				tabl := strings.Split(fileLines[i], " ")

				start = tabl[0]

			}

			if fileLines[i-1] == "##end" {

				tab := strings.Split(fileLines[i], " ")

				end = tab[0]

			} else {

				if strings.Contains(fileLines[i], "#") {

				} else {

					temp = strings.Split(fileLines[i], " ")
					if len(temp) == 3 {
						rooms = append(rooms, temp[0])
					}

					if strings.Contains(fileLines[i], "-") {
						links = append(links, temp[0])

					}

				}

			}
		}

	} else {
		fmt.Println("Besoin d'un fichier txt pour executer le projet !!!")
		os.Exit(3)
	}
	return ant, start, end, rooms, links

}
