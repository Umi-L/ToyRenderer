package main

import (
	"image/color"
	"math/rand"
)

type Tri struct {
	p1 Vector3
	p2 Vector3
	p3 Vector3

	Color color.Color
}

func (t *Tri) GetNormal() Vector3 {
	// get cross product of p1-p2 and p1-p3 https://stackoverflow.com/questions/19350792/calculate-normal-of-a-single-triangle-in-3d-space

	A := t.p2.Sub(t.p1)
	B := t.p3.Sub(t.p1)

	return A.Cross(B)
}

func (t *Tri) SetRandomColor() {
	// make triangle color random

	t.Color = color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}