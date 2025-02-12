package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	loadtex  *sdl.Texture
	loadsurf *sdl.Surface

	floortiles, walltiles TILESHEET
)

type TILESHEET struct {
	tex *sdl.Texture
	r   []sdl.Rect
}

type IM struct {
}

func makeimages() {

	loadsurf, _ = img.Load("im/floor2.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	floortiles = maketiles(loadtex, 64, 64, 3)
	loadsurf, _ = img.Load("im/wall.png")
	loadtex, _ = RNDR.CreateTextureFromSurface(loadsurf)
	walltiles = maketiles(loadtex, 64, 64, 4)
}

func maketiles(t *sdl.Texture, w, h, num int32) TILESHEET {

	ts := TILESHEET{}
	ts.tex = t
	var x, y int32 = 0, 0

	for num > 0 {
		r := sdl.Rect{x, y, w, h}
		ts.r = append(ts.r, r)
		x += w
		num--
	}

	return ts
}
