package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	WIN  *sdl.Window
	RNDR *sdl.Renderer

	winW, winH int32 = 1920, 1080

	displayMode sdl.DisplayMode

	surface *sdl.Surface
	texture *sdl.Texture

	OFF bool

	setFPS = 60

	CNTR sdl.Point
)

func PLAY() {

	//GET SCREEN SIZE
	WIN, _ = sdl.CreateWindow("", 0, 0, 0, 0, sdl.WINDOW_HIDDEN)
	defer WIN.Destroy()
	displayMode, _ = sdl.GetCurrentDisplayMode(0)
	if displayMode.W > 0 {
		winW = displayMode.W
	}
	if displayMode.H > 0 {
		winH = displayMode.H
	}
	WIN.Destroy()

	//CREATE WINDOW
	WIN, _ = sdl.CreateWindow("Not a 3D dungeon", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winW, winH, sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_SHOWN)
	RNDR, _ = sdl.CreateRenderer(WIN, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)

	RNDR.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	RNDR.SetRenderTarget(texture)

	defer WIN.Destroy()
	defer RNDR.Destroy()
	defer texture.Destroy()

	CNTR = sdl.Point{winW / 2, winH / 2}

	//CREATE TRANSPARENT SURFACE WINDOW SIZE
	/*
		if sdl.BYTEORDER == sdl.BIG_ENDIAN {
			surface, _ = sdl.CreateRGBSurface(sdl.SWSURFACE, int32(winW), int32(winH), 32, 0xFF000000, 0x00FF0000, 0x0000FF00, 0x000000FF)
		} else {
			surface, _ = sdl.CreateRGBSurface(sdl.SWSURFACE, int32(winW), int32(winH), 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
		}
	*/

	//FPS
	targetFPS = float64(setFPS)
	frameDelay = 1000 / targetFPS // Delay in milliseconds

	INITIAL()

	for !OFF {

		B4()

		DRAW()

		AFTER()

	}

}

func main() {
	PLAY()
}
