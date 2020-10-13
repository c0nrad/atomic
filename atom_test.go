package main

import "testing"

func TestAtomConfiguration(t *testing.T) {
	a := NewAtom(65)
	if a.Configuration() != "1s^2 2s^2 2p^6 3s^2 3p^6 4s^2 3d^10 4p^6 5s^2 4d^10 5p^6 6s^2 4f^9" {
		t.Error("failed to generate config", a.Configuration())
	}
}

func TestAtomIonization(t *testing.T) {
	a := NewAtom(1)
	if a.IonizationEnergy[0]-13.6 > .01 {
		t.Error("invalid ionization", a.IonizationEnergy[0])
	}
}
