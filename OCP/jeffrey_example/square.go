package main

type Square struct {
	With   float64
	Height float64
}

func (s Square) area() float64 {
	return s.With * s.Height
}
