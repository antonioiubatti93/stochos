package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Surface_Value_Flat(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0.2, NewFlatSurface(0.2).Value(1.0, 100.0))
}
