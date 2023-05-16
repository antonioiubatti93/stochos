package main

import "math"

type LocalVolatility struct {
	value   State
	drift   Drift
	surface VolatilitySurface
	normal  Distribution
}

var _ Process = &LocalVolatility{}

func NewLocalVolatility(value State, drift Drift, surface VolatilitySurface, seed int64) *LocalVolatility {
	return &LocalVolatility{
		value:   value,
		drift:   drift,
		surface: surface,
		normal:  NewStandardNormal(seed),
	}
}

func (lv LocalVolatility) Current() State {
	return lv.value
}

func (lv *LocalVolatility) Next(s, t float64) State {
	dt := t - s
	vol := lv.surface.Value(s, lv.value)
	drift := (lv.drift.Value(s, t) - 0.5*vol*vol) * dt
	diffusion := vol * lv.normal.Sample() * math.Sqrt(dt)
	lv.value *= math.Exp(drift + diffusion)

	return lv.Current()
}
