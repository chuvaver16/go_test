package main

import (
	"fmt"
	"math"
)

type Figure2D struct {
	Square    float64
	Perimeter float64
}

type Rectangle struct {
	X float64
	Y float64
	Figure2D
}

type Cyrcle struct {
	R float64
	Figure2D
}

func (f *Figure2D) Print() {
	fmt.Printf("[Square = %f] [Perimeter = %f]\n", f.Square, f.Perimeter)
}

func (r *Rectangle) Print() {
	fmt.Printf("[X = %f] [Y = %f]\n", r.X, r.Y)
}

func (c *Cyrcle) Print() {
	fmt.Printf("[R = %f] [D = %f]\n", c.R, 2*c.R)
}

func (r *Rectangle) Calc() {
	r.Square = r.X * r.Y
	r.Perimeter = 2 * (r.X + r.Y)
}

func (c *Cyrcle) Calc() {
	c.Square = math.Pi * math.Pow(c.R, 2)
	c.Perimeter = 2 * math.Pi * c.R
}

func (c *Cyrcle) Calc2() {
	c.R = math.Pow(c.Square/math.Pi, 0.5)
	c.Perimeter = 2 * math.Pi * c.R
}
