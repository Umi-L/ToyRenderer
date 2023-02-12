package main

import "math"

type Transform struct {
	Position Vector3
	Rotation Vector3
	Scale    Vector3
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// math operations on Vector3
func (v *Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v *Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v *Vector3) Mul(value float32) Vector3 {
	return Vector3{v.X * value, v.Y * value, v.Z * value}
}

func (v *Vector3) Div(v2 Vector3) Vector3 {
	return Vector3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}

func (v *Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{v.Y*v2.Z - v.Z*v2.Y, v.Z*v2.X - v.X*v2.Z, v.X*v2.Y - v.Y*v2.X}
}

func (v *Vector3) Dot(v2 Vector3) float32 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v *Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v *Vector3) AngleBetween(v2 Vector3) Vector3 {
	// get angle between two vectors

	// get dot product
	dot := v.Dot(v2)

	// get magnitude
	mag1 := v.Magnitude()

	mag2 := v2.Magnitude()

	// get angle
	angle := float32(math.Acos(float64(dot / (mag1 * mag2))))

	return Vector3{angle, angle, angle}
}

func (v *Vector3) Normalize() Vector3 {
	// get magnitude
	mag := v.Magnitude()

	// divide by magnitude
	return Vector3{v.X / mag, v.Y / mag, v.Z / mag}
}

func (v *Vector3) RayBetween(v2 Vector3) Vector3 {
	// get vector between two points
	return v2.Sub(*v)
}

func NewVector3(x, y, z float32) Vector3 {
	return Vector3{x, y, z}
}
