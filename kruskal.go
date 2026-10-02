package main

import "fmt"

type Edge struct {
	u, v, w int
}

func find(p []int, x int) int {
	if p[x] != x {
		p[x] = find(p, p[x])
	}
	return p[x]
}

func main() {
	e := []Edge{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}

	p := []int{0, 1, 2, 3}

	for i := 0; i < len(e)-1; i++ {
		for j := i + 1; j < len(e); j++ {
			if e[i].w > e[j].w {
				e[i], e[j] = e[j], e[i]
			}
		}
	}

	for _, x := range e {
		a, b := find(p, x.u), find(p, x.v)
		if a != b {
			p[a] = b
			fmt.Println(x.u, "-", x.v, "=", x.w)
		}
	}
}
