package stochos

import "golang.org/x/exp/slices"

type TimeGrid []float64

func NewTimeGrid(ts ...float64) TimeGrid {
	slices.Sort(ts)

	return TimeGrid(ts)
}

func NewUniformTimeGrid(start, end float64, m int) TimeGrid {
	h := (end - start) / float64(m)

	g := make(TimeGrid, m+1)
	for i := 0; i <= m; i++ {
		g[i] = start + h*float64(i)
	}

	return g
}
