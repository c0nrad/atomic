package main

import (
	"fmt"
	"math"
	"math/big"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Effective Nuclear Charge

func EffectiveNuclearCharge1(a Atom) float64 {
	count := []int{2, 8, 8, 8, 18}
	n := a.Number
	inside := float64(0)
	for _, c := range count {
		if n <= c {
			return float64(a.Number) - inside
		}
		n -= c
		inside += float64(c)
	}
	panic("what")
}

func EffectiveNuclearCharge2(a Atom) float64 {
	count := []int{2, 8, 8, 8, 18}
	n := a.Number
	inside := float64(0)
	for _, c := range count {
		if n <= c {
			return float64(a.Number) - inside - float64(n-1)/2
		}
		n -= c
		inside += float64(c)
	}
	panic("what")
}

func IonizationEnergy(a Atom, enc float64) float64 {

	// enc := EffectiveNuclearCharge1(a)

	// SI units
	electronMass := big.NewFloat(9.1093837015e-31)
	hBar := big.NewFloat(1.054571817e-34)
	permitivity := big.NewFloat(8.8541878128e-12)
	pi := big.NewFloat(math.Pi)
	e := big.NewFloat(1.602176634e-19)
	jToEv := big.NewFloat(1.602176634e-19)

	// (e^2 / 4*pi*e_0)^2
	tmp := new(big.Float).Mul(e, e)
	tmp.Mul(tmp, big.NewFloat(float64(enc))) // Effective Nuclear Charge
	tmp.Quo(tmp, big.NewFloat(4))
	tmp.Quo(tmp, pi)
	tmp.Quo(tmp, permitivity)
	tmp.Mul(tmp, tmp)

	// m_e / 2 hbar^2 (tmp)
	tmp.Mul(tmp, electronMass)
	tmp.Quo(tmp, big.NewFloat(2))
	tmp.Quo(tmp, hBar)
	tmp.Quo(tmp, hBar)

	// -tmp (1/n^2)
	// tmp.Quo(tmp, big.NewFloat(1))
	// tmp.Quo(tmp, big.NewFloat(1))
	tmp.Quo(tmp, big.NewFloat(float64(N(a.Number))))
	tmp.Quo(tmp, big.NewFloat(float64(N(a.Number))))

	// tmp.Mul(tmp, big.NewFloat(-1))

	tmp.Quo(tmp, jToEv)

	out, _ := tmp.Float64()
	return out

}

func PlotIonizationEnergy() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Experimental Vs Calculated Effective Nuclear Charge Ionization Energy"
	p.X.Label.Text = "Atomic Number"
	p.Y.Label.Text = "eV"

	experimental := make(plotter.XYs, 25)
	enc1 := make(plotter.XYs, 25)
	enc2 := make(plotter.XYs, 25)
	for i := 1; i <= 25; i++ {
		a := NewAtom(i)

		experimental[i-1].Y = a.IonizationEnergy[0]
		experimental[i-1].X = float64(i)

		enc1[i-1].X = float64(i)
		enc1[i-1].Y = IonizationEnergy(a, EffectiveNuclearCharge1(a))

		enc2[i-1].X = float64(i)
		enc2[i-1].Y = IonizationEnergy(a, EffectiveNuclearCharge2(a))
		fmt.Println(i, EffectiveNuclearCharge1(a), EffectiveNuclearCharge2(a))
		// fmt.Println(experimental[i-1])
	}

	err = plotutil.AddLinePoints(p,
		"Experimental", experimental,
		"Effective Nuclear Charge 1", enc1,
		"Effective Nuclear Charge 2", enc2)
	if err != nil {
		panic(err)
	}
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "ionization.png"); err != nil {
		panic(err)
	}
}
