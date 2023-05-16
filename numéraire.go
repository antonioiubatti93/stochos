package stochos

import "math"

type Numéraire struct {
	value State
	drift Drift
}

var _ Process = &Numéraire{}

func (n Numéraire) Current() State {
	return n.value
}

func (n *Numéraire) Next(s, t float64) State {
	n.value *= math.Exp(n.drift.Value(s, t) * (t - s))

	return n.value
}

func NewNuméraire(value State, drift Drift) *Numéraire {
	return &Numéraire{
		value: value,
		drift: drift,
	}
}
