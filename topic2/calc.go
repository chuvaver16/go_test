package main

import "fmt"

type Operation string

const (
	Plus     Operation = "+"
	Minus              = "-"
	Multiply           = "*"
	Division           = "/"
)

type Expression struct {
	X    float64
	Y    float64
	Oper Operation
}

func (exp *Expression) Calc() float64 {
	switch exp.Oper {
	case Plus:
		return exp.X + exp.Y
	case Minus:
		return exp.X - exp.Y
	case Multiply:
		return exp.X * exp.Y
	case Division:
		return exp.X / exp.Y
	}
	return 666
}

func DoCalc() {
	var x, y float64
	var oper Operation

	fmt.Println("Введите первое чмсло: ")
	fmt.Scan(&x)
	fmt.Println("Введите второе чмсло: ")
	fmt.Scan(&y)
	fmt.Println("Введите операцию: ")
	fmt.Scan(&oper)

	var exp = Expression{X: x, Y: y, Oper: oper}

	res := exp.Calc()

	fmt.Printf("Результат: %f", res)

}
