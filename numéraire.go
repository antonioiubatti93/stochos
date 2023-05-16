package main

import "math"

type Numéraire struct {
	drift Drift
	value State
}

var _ Process = &Numéraire{}

func (n Numéraire) Current() State {
	return n.value
}

func (n *Numéraire) Next(s, t float64) State {
	n.value *= math.Exp(n.drift.Value(s, t) * (t - s))

	return n.value
}

func NewNuméraire(drift Drift, value float64) *Numéraire {
	return &Numéraire{
		drift: drift,
		value: value,
	}
}
