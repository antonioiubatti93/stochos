package main

import "math"

type Geometric struct {
	value      State
	drift      Drift
	volatility float64
	normal     Distribution
}

var _ Process = &Geometric{}

func NewGeometric(value State, drift Drift, volatility float64, distribution Distribution) *Geometric {
	return &Geometric{
		value:      value,
		drift:      drift,
		volatility: volatility,
		normal:     distribution,
	}
}

func NewGeometricBrownian(value State, drift Drift, volatility float64, seed int64) *Geometric {
	return NewGeometric(value, drift, volatility, NewStandardNormal(seed))
}

func (bs Geometric) Current() State {
	return bs.value
}

func (bs *Geometric) Next(s, t float64) State {
	dt := t - s
	drift := (bs.drift.Value(s, t) - 0.5*bs.volatility*bs.volatility) * dt
	diffusion := bs.volatility * bs.normal.Sample() * math.Sqrt(dt)
	bs.value *= math.Exp(drift + diffusion)

	return bs.Current()
}
