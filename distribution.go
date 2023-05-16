package main

import "math/rand"

type Distribution interface {
	Sample() float64
}

type DistributionFunc func() float64

var _ Distribution = DistributionFunc(nil)

func (f DistributionFunc) Sample() float64 {
	return f()
}

func NewStandardNormal(seed int64) DistributionFunc {
	return rand.New(rand.NewSource(seed)).NormFloat64
}
