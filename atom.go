package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

var RawAtoms = []Atom{}

func LoadAtoms() {

	data, err := ioutil.ReadFile("./elements.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &RawAtoms)
	if err != nil {
		panic(err)
	}
}

func init() {
	LoadAtoms()
}

type Atom struct {
	Name             string
	Symbol           string
	Number           int
	IonizationEnergy []float64 `json:"ionization_energies"`
}

func KJMoleToEV(a float64) float64 {
	return a * .010364
}

func NewAtom(n int) Atom {
	a := RawAtoms[n-1]

	a.IonizationEnergy[0] = KJMoleToEV(a.IonizationEnergy[0])

	return RawAtoms[n-1]
}

func (a *Atom) Configuration() string {
	remaining := a.Number
	out := ""
	for _, o := range Orbitals {
		if o.Size() < remaining {
			out += fmt.Sprintf("%d%s^%d ", o.N, o.L, o.Size())
			remaining -= o.Size()
			continue
		}

		out += fmt.Sprintf("%d%s^%d", o.N, o.L, remaining)
		return out
	}
	panic("not enough orbitals for atom number:" + strconv.Itoa(a.Number))
}
