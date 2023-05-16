package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Simulate_Numéraire(t *testing.T) {
	t.Parallel()

	const (
		tol  = 1.0e-15
		seed = 1234
	)

	for _, tc := range []struct {
		name     string
		process  Process
		grid     TimeGrid
		expected Path
	}{
		{
			"empty grid",
			NewNuméraire(1.0, NewZeroDrift()),
			NewTimeGrid(),
			Path{},
		},
		{
			"single point",
			NewNuméraire(1.0, NewZeroDrift()),
			NewTimeGrid(0.0),
			Path{1.0},
		},
		{
			"flat numéraire",
			NewNuméraire(1.0, NewZeroDrift()),
			NewTimeGrid(0.0, 1.0, 2.0, 3.0),
			Path{1.0, 1.0, 1.0, 1.0},
		},
		{
			"numéraire",
			NewNuméraire(1.0, NewConstantDrift(math.Log(1.01))),
			NewUniformTimeGrid(0.0, 3.0, 3),
			Path{1.0, 1.01, 1.0201, 1.030301},
		},
		{
			"geometric brownian/no drift/no volatility",
			NewGeometricBrownian(1.0, NewZeroDrift(), 0.0, seed),
			NewTimeGrid(0.0, 1.0, 2.0),
			Path{1.0, 1.0, 1.0},
		},
		{
			"geometric brownian/with drift/no volatility",
			NewGeometricBrownian(1.0, NewConstantDrift(math.Log(1.01)), 0.0, seed),
			NewTimeGrid(0.0, 1.0, 2.0),
			Path{1.0, 1.01, 1.0201},
		},
		{
			"geometric/no drift/with diffusion",
			NewGeometric(1.0, NewZeroDrift(), 0.2, DistributionFunc(func() float64 {
				return -1.0
			})),
			NewTimeGrid(0.0, 1.0),
			Path{1.0, math.Exp(-0.5*0.04 - 1.0*0.2)},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDeltaSlice(t, tc.expected, Simulate(tc.process, tc.grid), tol)
		})
	}
}

func Test_Simulate_LocalVolatility_FlatSurface(t *testing.T) {
	t.Parallel()

	const (
		value = 1.0
		mu    = 0.02
		vol   = 0.2
		seed  = 123243
		tol   = 1.0e-15
	)

	grid := NewUniformTimeGrid(0.0, 10.0, 1.0)
	drift := NewConstantDrift(mu)

	lv := Simulate(NewLocalVolatility(value, drift, NewFlatSurface(vol), seed), grid)
	bs := Simulate(NewGeometricBrownian(value, drift, vol, seed), grid)

	assert.InDeltaSlice(t, bs, lv, tol)
}
