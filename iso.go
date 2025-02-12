package main

import "github.com/veandco/go-sdl2/sdl"

var (

	//GRID
	gridNumW = int32(18)
	grid     []OBJ
	tileSize = int32(52)
)

type OBJ struct {
	p                 []sdl.Point
	zi, ty, num       int
	ts, ts2           TILESHEET
	i, i2, i3, ir     sdl.Rect
	wall, wall2, opac bool
}

func addinternalwalls() {

	num := RINT(7, 12)

	for num > 0 {

		choose := RINT(0, 8)

		choose2 := RINT(int(gridNumW*2), len(grid)-int(gridNumW*2)) //CHOOSE GRID BLOK

		if !grid[choose2].wall {

			grid[choose2].wall2 = true

			added := false

			switch choose {
			case 7: //3 BLOK D
				nextD := findadjblokD(grid[choose2].num)
				grid[findadjblokD(grid[choose2].num)].wall2 = true
				grid[findadjblokD(nextD)].wall2 = true
				added = true
			case 6: //3 BLOK R
				nextR := findadjblokR(grid[choose2].num)
				grid[findadjblokR(grid[choose2].num)].wall2 = true
				grid[findadjblokR(nextR)].wall2 = true
				added = true
			case 5: //2 BLOK DR
				grid[findadjblokR(grid[choose2].num)].wall2 = true
				grid[findadjblokL(grid[choose2].num)].wall2 = true
				grid[findadjblokU(grid[choose2].num)].wall2 = true
				grid[findadjblokD(grid[choose2].num)].wall2 = true
				added = true
			case 4: //2 BLOK DR
				grid[findadjblokDR(grid[choose2].num)].wall2 = true
				added = true
			case 3: //2 BLOK D
				grid[findadjblokD(grid[choose2].num)].wall2 = true
				added = true
			case 2: //2 BLOK R
				grid[findadjblokR(grid[choose2].num)].wall2 = true
				added = true
			case 1: //4 BLOK
				grid[findadjblokR(grid[choose2].num)].wall2 = true
				grid[findadjblokD(grid[choose2].num)].wall2 = true
				grid[findadjblokDR(grid[choose2].num)].wall2 = true
				added = true
			}

			if added {
				num--
			}
		}

	}

	//REDO WALLS FOR POSSIBLE INTERRUPTS

	for i := 0; i < len(grid); i++ {
		if grid[i].num > 305 {
			grid[i].wall2 = false
		}
		if (grid[i].num-(int(gridNumW)-1))%int(gridNumW) == 0 {
			grid[i].wall2 = false
		}
	}

}
func findadjblokL(c int) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		if grid[i].num == c+1 {
			num = i
			break
		}
	}
	return num
}
func findadjblokR(c int) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		if grid[i].num == c-1 {
			num = i
			break
		}
	}
	return num
}
func findadjblokU(c int) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		if grid[i].num == c+int(gridNumW) {
			num = i
			break
		}
	}
	return num
}
func findadjblokD(c int) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		if grid[i].num == c-int(gridNumW) {
			num = i
			break
		}
	}
	return num
}
func findadjblokDR(c int) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		if grid[i].num == c-int(gridNumW+1) {
			num = i
			break
		}
	}
	return num
}

func makegrid(numW int32) {

	//w := (numW * 2) * tileSize
	h := numW * tileSize

	x := CNTR.X
	y := CNTR.Y
	y += h / 2
	ox := x
	oy := y

	a := numW * numW

	z := 0
	oz := z

	c := 0
	c2 := 0

	chooseFloor := RINT(0, len(floortiles.r))
	chooseWall := RINT(0, len(walltiles.r))
	chooseWall2 := chooseWall
	for chooseWall2 == chooseWall {
		chooseWall2 = RINT(0, len(walltiles.r))
	}

	for a > 0 {

		o := OBJ{}
		o.num = c2
		o.ts = floortiles
		o.ts2 = walltiles

		o.zi = z
		//ADD WALLS
		if c2 == int(gridNumW)-1 {
			o.wall = true
			//o.opac = true
		}
		if c2 < int(gridNumW) {
			o.wall = true
			if c2 < 17 {
				o.opac = true
			}
		}
		if c2 > 305 {
			o.wall = true
		}
		if c2 == 0 {
			o.wall = true
			o.opac = true
		}
		if c2%int(gridNumW) == 0 {
			o.wall = true
			if c2 < 306 {
				o.opac = true
			}
		}
		if (c2-(int(gridNumW)-1))%int(gridNumW) == 0 {
			o.wall = true
		}

		p1 := sdl.Point{x, y}
		p2 := p1
		p2.X -= tileSize
		p2.Y -= tileSize / 2
		p3 := p1
		p3.Y -= tileSize
		p4 := p1
		p4.X += tileSize
		p4.Y -= tileSize / 2

		o.p = append(o.p, p1, p2, p3, p4)

		o.ir = sdl.Rect{p2.X, p3.Y, tileSize * 2, tileSize * 2}

		o.i = floortiles.r[chooseFloor]
		o.i2 = walltiles.r[chooseWall]
		o.i3 = walltiles.r[chooseWall2]
		/*
			if Roll6() == 6 {
				switch RINT(0, 7) {
				case 0:
					o.i = floortiles.r[15]
				case 1:
					o.i = floortiles.r[9]
				case 2:
					o.i = floortiles.r[5]
				case 3:
					o.i = floortiles.r[6]
				case 4:
					o.i = floortiles.r[7]
				case 5:
					o.i = floortiles.r[8]
				case 6:
					o.i = floortiles.r[4]
				}
			}
			if Roll6() == 6 {
				switch RINT(0, 3) {
				case 0:
					o.i2 = walltiles.r[0]
				case 1:
					o.i2 = walltiles.r[1]
				case 2:
					o.i2 = walltiles.r[3]
				}
			}
			if Roll6() == 6 {
				switch RINT(0, 3) {
				case 0:
					o.i3 = walltiles.r[0]
				case 1:
					o.i3 = walltiles.r[1]
				case 2:
					o.i3 = walltiles.r[3]
				}
			}
		*/
		grid = append(grid, o)

		x -= tileSize
		y -= tileSize / 2

		c++
		c2++
		z++
		a--

		if c == int(numW) {
			c = 0
			z = oz
			z++
			oz = z
			x = ox
			x += tileSize
			ox = x
			y = oy
			y -= tileSize / 2
			oy = y
		}
	}

	grid = gridSORT(grid)
	addinternalwalls()
}
