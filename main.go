package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
	"math/rand"
)

func main() {
	//make new Camera object

	world := World{}

	world.Camera.OutputWidth = 420
	world.Camera.OutputHeight = 420
	world.Camera.FOV = 90
	world.Camera.World = &world

	// make new cube
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: 1, Y: -1, Z: -1}, p3: Vector3{X: 1, Y: 1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: 1, Y: 1, Z: -1}, p3: Vector3{X: -1, Y: 1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: -1, Y: 1, Z: -1}, p3: Vector3{X: -1, Y: 1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: -1, Y: 1, Z: 1}, p3: Vector3{X: -1, Y: -1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: -1, Y: -1, Z: 1}, p3: Vector3{X: 1, Y: -1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: -1, Y: -1, Z: -1}, p2: Vector3{X: 1, Y: -1, Z: 1}, p3: Vector3{X: 1, Y: -1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: -1, Y: 1, Z: 1}, p3: Vector3{X: -1, Y: -1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: -1, Y: -1, Z: 1}, p3: Vector3{X: 1, Y: -1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: 1, Y: -1, Z: 1}, p3: Vector3{X: 1, Y: -1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: 1, Y: -1, Z: -1}, p3: Vector3{X: 1, Y: 1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: 1, Y: 1, Z: -1}, p3: Vector3{X: -1, Y: 1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: 1}, p2: Vector3{X: -1, Y: 1, Z: -1}, p3: Vector3{X: -1, Y: 1, Z: 1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: -1}, p2: Vector3{X: -1, Y: 1, Z: -1}, p3: Vector3{X: -1, Y: -1, Z: -1}})
	world.Tris = append(world.Tris, Tri{p1: Vector3{X: 1, Y: 1, Z: -1}, p2: Vector3{X: -1, Y: -1, Z: -1}, p3: Vector3{X: 1, Y: -1, Z: -1}})

	// for every tri
	for i := 0; i < len(world.Tris); i++ {
		// set color to random color
		world.Tris[i].Color = color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
	}

	// translate tris 5 units away from camera
	for i := 0; i < len(world.Tris); i++ {
		world.Tris[i].p1 = world.Tris[i].p1.Add(Vector3{X: 0, Y: 0, Z: 5})
		world.Tris[i].p2 = world.Tris[i].p2.Add(Vector3{X: 0, Y: 0, Z: 5})
		world.Tris[i].p3 = world.Tris[i].p3.Add(Vector3{X: 0, Y: 0, Z: 5})
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(world.Camera.OutputWidth), int32(world.Camera.OutputHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	world.Camera.Window = window

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}

		// move camera when w,a,s,d is pressed
		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_W] != 0 {
			world.Camera.Transform.Position.Z += 0.1
		}
		if keys[sdl.SCANCODE_S] != 0 {
			world.Camera.Transform.Position.Z -= 0.1
		}
		if keys[sdl.SCANCODE_A] != 0 {
			world.Camera.Transform.Position.X -= 0.1
		}
		if keys[sdl.SCANCODE_D] != 0 {
			world.Camera.Transform.Position.X += 0.1
		}

		//space and shift to move up and down
		if keys[sdl.SCANCODE_SPACE] != 0 {
			world.Camera.Transform.Position.Y -= 0.1
		}
		if keys[sdl.SCANCODE_LSHIFT] != 0 {
			world.Camera.Transform.Position.Y += 0.1
		}

		//rotate camera when arrow keys are pressed
		if keys[sdl.SCANCODE_UP] != 0 {
			world.Camera.Transform.Rotation.Y += 0.1
		}
		if keys[sdl.SCANCODE_DOWN] != 0 {
			world.Camera.Transform.Rotation.Y -= 0.1
		}
		if keys[sdl.SCANCODE_LEFT] != 0 {
			world.Camera.Transform.Rotation.X += 0.1
		}
		if keys[sdl.SCANCODE_RIGHT] != 0 {
			world.Camera.Transform.Rotation.X -= 0.1
		}

		world.Camera.Render()

		window.UpdateSurface()
	}
}
