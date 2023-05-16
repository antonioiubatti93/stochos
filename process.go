package main

type State = float64

type Process interface {
	Current() State
	Next(s, t float64) State
}

type Path []State

func Simulate(process Process, grid TimeGrid) Path {
	n := len(grid)
	if n == 0 {
		return Path{}
	}

	path := make(Path, 0, n)
	path = append(path, process.Current())

	if n == 1 {
		return path
	}

	for i, t := range grid[:n-1] {
		path = append(path, process.Next(t, grid[i+1]))
	}

	return path
}
