package main

import "strings"

type Orbital struct {
	N int
	L string
}

var OrbitalCountMap = map[string]int{
	"s": 2, "p": 6, "d": 10, "f": 14, "g": 18, "h": 22,
}

var Orbitals = []Orbital{
	{1, "s"}, {2, "s"}, {2, "p"}, {3, "s"}, {3, "p"}, {4, "s"}, {3, "d"}, {4, "p"}, {5, "s"}, {4, "d"}, {5, "p"}, {6, "s"}, {4, "f"}, {5, "d"}, {6, "p"}, {7, "s"}, {5, "f"},
}

func (o Orbital) Size() int {
	l := strings.ToLower(o.L)
	c, ok := OrbitalCountMap[l]
	if !ok {
		panic("invalid orbital:" + l)
	}
	return c
}

func N(n int) int {
	count := []int{2, 8, 18, 18 + 14, 18 + 14 + 22}
	for i, c := range count {
		if n <= c {
			return i + 1
		}
		n -= c
	}
	panic("what")
}
