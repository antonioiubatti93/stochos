package stochos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTimeGrid(t *testing.T) {
	t.Parallel()

	const tol = 1.0e-15

	for _, tc := range []struct {
		name     string
		grid     TimeGrid
		expected TimeGrid
	}{
		{
			"from input/ordered",
			NewTimeGrid(0.0, 1.0, 2.0),
			TimeGrid{0.0, 1.0, 2.0},
		},
		{
			"from input/unordered",
			NewTimeGrid(2.0, 0.0, 1.0),
			TimeGrid{0.0, 1.0, 2.0},
		},
		{
			"uniform",
			NewUniformTimeGrid(0.0, 1.0, 5),
			NewTimeGrid(0.0, 0.2, 0.4, 0.6, 0.8, 1.0),
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDeltaSlice(t, tc.expected, tc.grid, tol)
		})
	}
}
