package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Simulate_Numéraire(t *testing.T) {
	t.Parallel()

	const tol = 1.0e-15

	for _, tc := range []struct {
		name     string
		process  Process
		grid     TimeGrid
		expected Path
	}{
		{
			"empty grid",
			NewNuméraire(NewZeroDrift(), 1.0),
			NewTimeGrid(),
			Path{},
		},
		{
			"single point",
			NewNuméraire(NewZeroDrift(), 1.0),
			NewTimeGrid(0.0),
			Path{1.0},
		},
		{
			"flat numéraire",
			NewNuméraire(NewZeroDrift(), 1.0),
			NewTimeGrid(0.0, 1.0, 2.0, 3.0),
			Path{1.0, 1.0, 1.0, 1.0},
		},
		{
			"numéraire",
			NewNuméraire(NewConstantDrift(math.Log(1.01)), 1.0),
			NewUniformTimeGrid(0.0, 3.0, 3),
			Path{1.0, 1.01, 1.0201, 1.030301},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDeltaSlice(t, tc.expected, Simulate(tc.process, tc.grid), tol)
		})
	}
}
