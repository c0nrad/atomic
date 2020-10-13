package main

import (
	"math"
	"testing"
)

func TestEffectiveNuclearCharge1(t *testing.T) {
	na := NewAtom(11)
	fl := NewAtom(9)

	if EffectiveNuclearCharge1(na) != 1 {
		t.Error("invalid enc", EffectiveNuclearCharge1(na), na.Configuration())
	}

	if EffectiveNuclearCharge1(fl) != 7 {
		t.Error("invalid enc", EffectiveNuclearCharge1(fl), fl.Configuration())
	}

	k := NewAtom(19)
	if EffectiveNuclearCharge1(k) != 1 {
		t.Error("invalid enc", EffectiveNuclearCharge1(k), k.Configuration())
	}
}

func TestIonizationEnergy(t *testing.T) {
	h := NewAtom(1)

	if math.Abs(IonizationEnergy(h, EffectiveNuclearCharge1(h))-13.6) > .01 {
		t.Error("Failed to calculate ionization energy for hydrogen", IonizationEnergy(h, EffectiveNuclearCharge1(h)))
	}

	// na := NewAtom(11)
	// if math.Abs(IonizationEnergy(na) - -13.6) > .01 {
	// 	t.Error("Failed to calculate ionization energy for hydrogen", IonizationEnergy())
	// }
}
