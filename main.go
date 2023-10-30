package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	for j, arg := range os.Args[0:] {
		fmt.Printf("%d %s\n", j, arg)
	}

	fmt.Println("Hello, МИР")
	fmt.Println(s)
	fmt.Println(os.Args[0:])
	fmt.Println(strings.Join(os.Args[0:], " "))

}
