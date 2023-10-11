package main

type CalculateGeometry struct {
}

type CalculateGeometryMethod interface {
	area() float64
}

func (c CalculateGeometry) CalculateArea(cm CalculateGeometryMethod) float64 {
	return cm.area()
}
