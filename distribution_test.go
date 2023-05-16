package stochos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Distribution_Sample(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0.0, DistributionFunc(func() float64 {
		return 0.0
	}).Sample())
}

func Test_NewStandardNormal(t *testing.T) {
	t.Parallel()

	assert.NotEqual(t, NewStandardNormal(0).Sample(), NewStandardNormal(1).Sample())
}
