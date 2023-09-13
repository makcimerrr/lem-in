package mypackage

import (
	"bufio"
	"fmt"
	"os"
)

func Checktxt() {

	if len(os.Args) == 2 {
		fichier := os.Args[2]
		file, err := os.Open(fichier)
		if err != nil {
			fmt.Println("cannot open the file !!")
			return
		}
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		file.Close()

		for _, line := range fileLines {
			fmt.Println(line)
		}

		fmt.Println(fileLines)
	} else {
		fmt.Println("Besoin d'un fichier txt pour executer le projet !!!")
		return
	}

}
