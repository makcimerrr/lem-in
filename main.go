package main

import (
	"fmt"
	"package/mypackage"
)

func main() {
	numberofant, start, end := mypackage.Checktxt()
	fmt.Println(numberofant)
	fmt.Println(start)
	fmt.Println(end)

}
