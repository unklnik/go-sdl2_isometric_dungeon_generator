package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (

	//TEXT
	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "

	font                    *ttf.Font
	tx1                     = 18
	fontchars               []FONTCH
	txSurf                  *sdl.Surface
	LineHeight, LetterSpace int32 = int32(tx1 + 4), 1

	//KEYS
	kEscape, kF1, kF2 bool

	//DEBUG
	DEBUG    bool
	debugRec REC

	//FPS TIMERS
	nowDelta, lastDelta uint64
	Delta               float32
	//FPS
	targetFPS          float64
	frameDelay         = float64(1000 / targetFPS) // Delay in milliseconds
	frameStart         time.Time
	frameTime          float64
	FPS                float64
	frameCount, frames int
	startTime          = time.Now()
)

type FONTCH struct {
	character string
	tex       *sdl.Texture
	rec       sdl.Rect
}

func INITIAL() {
	debugRec.r = sdl.Rect{0, 0, 300, winH}
	debugRec.col = DARKRED()
	debugRec.col = ChangeAlpha(debugRec.col, 100)

	ttf.Init()
	font, _ = ttf.OpenFont("Rubik-Medium.ttf", tx1)
	fontTextureSheet()

	makeimages()
	makegrid(gridNumW)

}

func debug() {
	recDRAW(debugRec)
	h := int32(font.Height())
	var x, y int32 = 4, 4
	txtXY("FPS "+fmt.Sprintf("%.0f", FPS), x, y)
	y += h

}
func EVENTS() {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch k := event.(type) {
		case sdl.QuitEvent:
			kEscape = true

		case sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				switch k.Keysym.Scancode {

				case sdl.SCANCODE_ESCAPE:
					kEscape = true
				case sdl.SCANCODE_F1:
					kF1 = true
				case sdl.SCANCODE_F2:
					kF2 = true
				}
			}
		}
	}

	//KEYS
	if kEscape {
		EXIT()
	}
	if kF2 {
		RESTART()
		kF2 = false
	}
	if kF1 {
		DEBUG = !DEBUG
		kF1 = false
	}

}

func RESTART() {
	grid = nil
	makegrid(gridNumW)
}

func B4() {
	frameStart = time.Now()
	EVENTS()
	RNDR.SetDrawColor(0, 0, 0, 0)
	RNDR.Clear()
}
func AFTER() {

	//texture, _ = RNDR.CreateTextureFromSurface(surface)
	//RNDR.Copy(texture, nil, &sdl.Rect{0, 0, int32(winW), int32(winH)})
	//surface.FillRect(&sdl.Rect{0, 0, surface.W, surface.H}, 0) //CLEAR SURFACE

	RNDR.Present()

	//texture.Destroy()

	//FPS
	frames++
	frameCount++

	frameTime = float64(time.Since(frameStart).Milliseconds())
	if frameTime < frameDelay {
		sdl.Delay(uint32(frameDelay - frameTime)) // Wait to achieve target FPS
	}
	if time.Since(startTime).Seconds() >= 1 {
		FPS = float64(frameCount) / time.Since(startTime).Seconds()
		frameCount = 0
		startTime = time.Now()
	}

}

func EXIT() {
	//surface.Free()
	//texture.Destroy()
	loadsurf.Free()
	loadtex.Destroy()
	RNDR.Destroy()
	WIN.Destroy()
	OFF = true
}

func getDelta() {
	tickT := sdl.GetTicks64()
	Delta = float32(tickT-lastDelta) * 0.001
	lastDelta = tickT
}

// TEXT
func txtXY(txt string, x, y int32) {
	t := strings.Split(txt, "")
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(fontchars); j++ {
			if t[i] == fontchars[j].character {
				RNDR.Copy(fontchars[j].tex, &fontchars[j].rec, &sdl.Rect{x, y, fontchars[j].rec.W, fontchars[j].rec.H})
				x += fontchars[j].rec.W + LetterSpace
				break
			}
		}
	}
}

func fontTexRevert(f []FONTCH) []FONTCH {
	for i := 0; i < len(f); i++ {
		fontchars[i].tex.SetColorMod(255, 255, 255)
	}
	return f
}
func fontTextureSheet() {
	t := strings.Split(standardCharacters, "")
	for i := 0; i < len(t); i++ {
		fontchars = append(fontchars, fontCharCreate(t[i], 0))
	}
}

func fontCharCreate(singleCharacter string, fontNum int) FONTCH {

	var w, h int

	switch fontNum {
	case 0: //FONT DEFAULT
		txSurf, _ = font.RenderUTF8Blended(singleCharacter, WHITEǁ3())
		defer txSurf.Free()
		w, h, _ = font.SizeUTF8(singleCharacter)
	}

	f := FONTCH{}
	f.tex, _ = RNDR.CreateTextureFromSurface(txSurf)
	f.rec = sdl.Rect{0, 0, int32(w), int32(h)}
	f.character = singleCharacter
	return f
}
