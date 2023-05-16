package stochos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Drift_Value(t *testing.T) {
	t.Parallel()

	const (
		s   = 1.0
		dt  = 0.1
		tol = 1.0e-15
	)

	for _, tc := range []struct {
		name     string
		drift    Drift
		expected float64
	}{
		{
			"zero",
			NewZeroDrift(),
			0.0,
		},
		{
			"constant",
			NewConstantDrift(0.02),
			0.02,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.drift.Value(s, s+dt), tol)
		})
	}
}
