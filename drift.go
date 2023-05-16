package stochos

type Drift interface {
	Value(s, t float64) float64
}

type DriftFunc func(s, t float64) float64

var _ Drift = DriftFunc(nil)

func (f DriftFunc) Value(s, t float64) float64 {
	return f(s, t)
}

func NewConstantDrift(c float64) DriftFunc {
	return func(_, _ float64) float64 {
		return c
	}
}

func NewZeroDrift() DriftFunc {
	return NewConstantDrift(0.0)
}
