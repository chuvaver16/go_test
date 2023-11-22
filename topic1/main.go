package main

import (
	"fmt"
	"strconv"
)

func CalcRectangle() {
	var x, y float64
	fmt.Println("Введите длину: ")
	fmt.Scan(&x)
	fmt.Println("Введите высоту: ")
	fmt.Scan(&y)

	var rect = Rectangle{X: x, Y: y}
	rect.Calc()
	rect.Figure2D.Print()
}

func CalcCyrcle() {
	var s float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scan(&s)

	var cyrcle = Cyrcle{}
	cyrcle.Square = s
	cyrcle.Calc2()
	cyrcle.Print()
	cyrcle.Figure2D.Print()
}

func GetDigits() {
	var n int
	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	digits := []rune(strconv.Itoa(n))
	for _, digit := range digits[0:] {
		fmt.Print(string(digit) + " ")
	}
	fmt.Println()

	for n > 0 {
		fmt.Printf("%d ", n%10)
		n = n / 10
	}
	fmt.Println()

}

func main() {

	// Calc rectangle's square
	//CalcRectangle()

	// Calc diamter and perimeter for cyrcle
	//CalcCyrcle()

	// Get numbers' digits
	GetDigits()
}
