package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const windowWidth, windowHeight = 800, 600

type color struct {
	r, g, b byte
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*windowWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+1] = c.b
	}
}

func main() {

	window, err := sdl.CreateWindow("testing SLD2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer texture.Destroy()

	pixels := make([]byte, windowWidth*windowHeight*4) // capacity of width * height * 4 bytes for RGBA

	for y := 0; y < windowHeight; y++ {
		for x := 0; x < windowWidth; x++ {
			setPixel(x, y, color{255, 0, 0}, pixels)
		}
	}

	// pitch is width of screen * how many bytes per pixel
	// internally it will probably divide by 4 bytes to get the width
	// ie each square (pixel) in the grid is 4 bytes
	texture.Update(nil, pixels, windowWidth*4)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	sdl.Delay(3000)
}
