package main

import "fmt"

var fibs = make(map[uint]uint64)

func Fibonacci(n uint) uint64 {
	if fibs[n] != 0 {
		return fibs[n]
	}

	var fib uint64

	if n == 1 {
		fib = 0
	} else if n == 2 {
		fib = 1
	} else {
		fib = Fibonacci(n-1) + Fibonacci(n-2)
	}

	fibs[n] = fib
	return fib
}

func DoFibonacci() {
	var n uint
	var f uint64

	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	f = Fibonacci(n)
	fmt.Println(f)
	f = Fibonacci(n)
	fmt.Println(f)
	f = Fibonacci(n)
	fmt.Println(f)

}
