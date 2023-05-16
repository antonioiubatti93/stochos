package main

type VolatilitySurface interface {
	Value(t, m float64) float64
}

type VolatilitySurfaceFunc func(t, m float64) float64

var _ VolatilitySurface = VolatilitySurfaceFunc(nil)

func (f VolatilitySurfaceFunc) Value(t, m float64) float64 {
	return f(t, m)
}

func NewFlatSurface(value float64) VolatilitySurfaceFunc {
	return func(_, _ float64) float64 {
		return value
	}
}
