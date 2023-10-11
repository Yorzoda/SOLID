package main

import "math"

type Circle struct {
	Radius float64
}

func (r Circle) area() float64 {
	return r.Radius * math.Pi
}
