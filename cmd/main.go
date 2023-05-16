package main

import (
	"flag"
	"image/color"
	"log"
	"math"

	"github.com/antonioiubatti93/stochos"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	out := flag.String("o", "paths.png", "output plot image .png")
	flag.Parse()

	p := plot.New()
	p.X.Label.Text = "t"
	p.Y.Label.Text = "St"
	p.Title.Text = "stochastic processes"
	p.Add(plotter.NewGrid())

	const (
		value = 1.0
		mu    = 0.02
		sigma = 0.4
		seed  = 1256
	)

	grid := stochos.NewUniformTimeGrid(0.0, 10.0, 1000)
	drift := stochos.NewConstantDrift(mu)

	for _, tc := range []struct {
		name    string
		process stochos.Process
		color   color.RGBA
	}{
		{
			"numéraire",
			stochos.NewNuméraire(value, drift),
			color.RGBA{A: 255, B: 255},
		},
		{
			"black scholes",
			stochos.NewGeometricBrownian(value, drift, sigma, seed),
			color.RGBA{A: 255, R: 255},
		},
		{
			"local volatility",
			stochos.NewLocalVolatility(value, drift, stochos.VolatilitySurfaceFunc(func(_, m float64) float64 {
				x := m/value - 1.0
				return sigma * math.Sqrt(1.0+x*x)
			}), seed),
			color.RGBA{A: 255, R: 128, B: 128},
		},
	} {
		xys := make(plotter.XYs, 0, len(grid))
		for k, v := range stochos.Simulate(tc.process, grid) {
			xys = append(xys, plotter.XY{
				X: grid[k],
				Y: v,
			})
		}

		line, err := plotter.NewLine(xys)
		if err != nil {
			log.Fatal(err)
		}
		line.Color = tc.color

		p.Add(line)
		p.Legend.Add(tc.name, line)
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, *out); err != nil {
		log.Fatal(err)
	}
}
