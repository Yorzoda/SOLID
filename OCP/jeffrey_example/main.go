package main

func main() {
	calculate := CalculateGeometry{}
	square := Square{With: 2.1, Height: 1.4}
	circle := Circle{Radius: 5.2}

	calculate.CalculateArea(square)
	calculate.CalculateArea(circle)

}
