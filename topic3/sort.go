package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InsertionSort(arr *[]int) {

	n := len(*arr)

	for i := 1; i < n; i++ {
		x := (*arr)[i]
		j := i
		for ; j > 0 && (*arr)[j-1] > x; j-- {
			(*arr)[j] = (*arr)[j-1]
		}
		(*arr)[j] = x
	}
}

func DoInsertionSort() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	strs := strings.Split(s, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}

	InsertionSort(&ary)

	fmt.Println(ary)

}
