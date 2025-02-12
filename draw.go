package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var ()

type REC struct {
	r   sdl.Rect
	col []uint8
}

// DRAW TO SCREEN
func DRAW() {

	gridDRAW(grid)

	if DEBUG {
		gridlinesDRAW(grid)
		//tilesheetDRAW(walltiles, 350, 10, 4, 5)
		debug()
	}

}

// GRID
func gridDRAW(o []OBJ) {
	for i := 0; i < len(o); i++ {
		RNDR.Copy(o[i].ts.tex, &o[i].i, &o[i].ir)
		if o[i].wall && !o[i].opac {
			r := o[i].ir
			r.Y -= o[i].ir.H / 2
			/*
				alphaChange := false
				if o[i].opac {
					o[i].ts2.tex.SetAlphaMod(70)
					alphaChange = true
				}
			*/
			RNDR.Copy(o[i].ts2.tex, &o[i].i2, &r)
			r.Y -= o[i].ir.H / 2
			RNDR.Copy(o[i].ts2.tex, &o[i].i2, &r)
			/*
				if alphaChange {
					o[i].ts2.tex.SetAlphaMod(255)
				}
			*/
		}
		if o[i].wall2 {
			r := o[i].ir
			r.Y -= o[i].ir.H / 2
			RNDR.Copy(o[i].ts2.tex, &o[i].i3, &r)
		}
	}
}

// ISO
func gridlinesDRAW(o []OBJ) {
	RNDR.SetDrawColor(GREENǁ2())
	for i := 0; i < len(o); i++ {
		isorecDRAW(o[i].p)
		txtXY(fmt.Sprint(o[i].num), o[i].p[0].X-int32(font.Height()/2), o[i].p[1].Y-int32(font.Height()/2))
	}
}

// ISO RECS
func isorecDRAW(p []sdl.Point) {
	lineDRAW(p[0], p[1])
	lineDRAW(p[1], p[2])
	lineDRAW(p[2], p[3])
	lineDRAW(p[0], p[3])
}

// IMAGES
func tilesheetDRAW(t TILESHEET, x, y, space, zoom int32) {

	r := sdl.Rect{x, y, t.r[0].W * zoom, t.r[0].H * zoom}

	for i := 0; i < len(t.r); i++ {
		RNDR.Copy(t.tex, &t.r[i], &r)
		txtXY(fmt.Sprint(i), r.X+2, r.Y+2)
		r.X += r.W + space
	}

}

// LINES
func lineDRAW(p1, p2 sdl.Point) {
	RNDR.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}

// RECS
func recDRAW(r REC) {
	RNDR.SetDrawColor(r.col[0], r.col[1], r.col[2], r.col[3])
	RNDR.FillRect(&r.r)
}
