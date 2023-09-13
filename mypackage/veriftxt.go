package mypackage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ant   int
	start int
	end   int
)

func Checktxt() (int, int, int) {
	if len(os.Args) == 2 {
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
		i := 0
		for range fileLines {
			intfirstline, err := strconv.Atoi(fileLines[0])
			if err != nil {
				fmt.Print("impossible de convertir la premiere ligne en int ")
				os.Exit(3)
			}
			if intfirstline > 0 {
				ant = intfirstline
			} else {
				fmt.Println("le nombre de ants doit être supérieur à 0 ")
				os.Exit(3)
			}
			if fileLines[i] != fileLines[0] {
				if fileLines[i-1] == "##start" {
					tabl := strings.Split(fileLines[i], " ")
					intstart, err := strconv.Atoi(tabl[0])
					if err != nil {
						fmt.Print("impossible de convertir la premiere ligne en int ")
						os.Exit(3)
					}
					start = intstart
				}
				if fileLines[i-1] == "##end" {
					tab := strings.Split(fileLines[i], " ")
					intend, err := strconv.Atoi(tab[0])
					if err != nil {
						fmt.Print("impossible de convertir la premiere ligne en int ")
						os.Exit(3)
					}
					end = intend
				}
			}
			i++
		}
	} else {
		fmt.Println("Besoin d'un fichier txt pour exécuter le projet !!!")
		os.Exit(3)
	}
	return ant, start, end
}
