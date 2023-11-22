package main

import (
	"fmt"
	"math"
	"sync"
)

func Eratosthene(n int) []int {
	a := make([]int, n+1)
	var res []int

	for p := 2; p < n+1; p++ {
		if a[p] == 0 {
			res = append(res, p)
			for j := p * p; j < n+1; j += p {
				a[j] = 1
			}
		}
	}

	return res
}

func EratostheneParallel(n int) []int {
	n_sqrt := int(math.Sqrt((float64)(n + 1)))
	a := make([]int, n+1)
	var wg sync.WaitGroup

	res := Eratosthene(n_sqrt)

	for _, p := range res {
		wg.Add(1)
		go func(p int, n int, a *[]int, wg *sync.WaitGroup) {
			defer wg.Done()

			for i := p * p; i <= n; i += p {
				(*a)[i] = 1
			}
		}(p, n, &a, &wg)
	}
	wg.Wait()

	for i := n_sqrt + 1; i <= n; i++ {
		if a[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

func DoPrime() {
	var n int

	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	//res1 := Eratosthene(n)
	res2 := EratostheneParallel(n)
	fmt.Println()
	//fmt.Printf("%v\n", res1)
	fmt.Printf("%v\n", res2)
}
