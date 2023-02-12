package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type Camera struct {
	Transform

	//add camera specific fields
	OutputWidth  int
	OutputHeight int

	FOV float32

	World *World

	Window *sdl.Window
}

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func (c *Camera) Render() {
	//render camera using fmt.printf

	//for every pixel in output
	for y := 0; y < c.OutputHeight; y++ {
		for x := 0; x < c.OutputWidth; x++ {
			surface, err := c.Window.GetSurface()
			if err != nil {
				panic(err)
			}

			//shoot ray from camera into world and if ray intersects tri than print a character
			surface.Set(x, y, c.ShootRay(c.PixelToRay(x, y)))
		}
		fmt.Printf("\n")
	}

}

// PixelToRay pixel coordinates to Vector3 in world
func (c *Camera) PixelToRay(x, y int) Vector3 {
	// find Camera aspect ratio
	aspectRatio := float32(c.OutputWidth) / float32(c.OutputHeight)

	//Convert screen space to Normalized Device Coordinates
	NDCX := (float32(x) + 0.5) / float32(c.OutputWidth)
	NDCY := (float32(y) + 0.5) / float32(c.OutputHeight)

	//Convert to Screen Space
	ScreenX := ((2.0 * NDCX) - 1.0) * aspectRatio
	ScreenY := (2.0 * NDCY) - 1.0

	// get point on plane 1 unit away from camera
	planePoint := Vector3{ScreenX, ScreenY, 1.0}

	// get angle between camera forward and plane point
	return c.Transform.Rotation.RayBetween(planePoint)
}

func (c *Camera) ShootRay(dir Vector3) color.Color {
	//shoot ray from camera at angle into world and if ray intersects tri than return true

	r := Ray{}
	r.Origin = c.Transform.Position
	r.Direction = dir

	for _, tri := range c.World.Tris {
		if r.IntersectsTri(tri) {

			// make triangle color random seeded on its index

			return tri.Color
		}
	}
	return color.Black
}

func (r *Ray) IntersectsTri(tri Tri) bool {
	//check if ray intersects tri and if it does than return true

	// if Direction is parallel to plane, then no intersection
	if r.Direction.Dot(tri.GetNormal()) == 0 {
		return false
	}

	// get intersection point between ray and tri plane

	// get distance from ray origin to plane
	sub := tri.p1.Sub(r.Origin)
	d := sub.Dot(tri.GetNormal()) / r.Direction.Dot(tri.GetNormal())

	// get intersection point
	intersection := r.Origin.Add(r.Direction.Mul(d))

	// check if intersection point is inside tri
	// get cross product of p1-p2 and p1-intersection
	A := tri.p2.Sub(tri.p1)
	B := intersection.Sub(tri.p1)
	C := A.Cross(B)

	// get cross product of p2-p3 and p2-intersection
	A = tri.p3.Sub(tri.p2)
	B = intersection.Sub(tri.p2)
	D := A.Cross(B)

	// get cross product of p3-p1 and p3-intersection
	A = tri.p1.Sub(tri.p3)
	B = intersection.Sub(tri.p3)
	E := A.Cross(B)

	// if all cross products have same sign than intersection point is inside tri
	if C.Dot(D) >= 0 && D.Dot(E) >= 0 {
		return true
	}

	return false
}
