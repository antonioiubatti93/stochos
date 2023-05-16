package main

import "math"

type BlackScholes struct {
	value      State
	drift      Drift
	volatility float64
	normal     Distribution
}

var _ Process = &BlackScholes{}

func NewBlackScholes(value State, drift Drift, volatility float64, seed int64) *BlackScholes {
	return &BlackScholes{
		value:      value,
		drift:      drift,
		volatility: volatility,
		normal:     NewStandardNormal(seed),
	}
}

func (bs BlackScholes) Current() State {
	return bs.value
}

func (bs *BlackScholes) Next(s, t float64) State {
	dt := t - s
	drift := (bs.drift.Value(s, t) - 0.5*bs.volatility*bs.volatility) * dt
	diffusion := bs.volatility * bs.normal.Sample() * math.Sqrt(dt)
	bs.value *= math.Exp(drift + diffusion)

	return bs.Current()
}
