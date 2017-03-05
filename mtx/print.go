// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mtx

//------------------------------------------------------------------------------

import (
	"fmt"
	micro "github.com/drakmaniso/glam/internal/microtext"
)

//------------------------------------------------------------------------------

func Size() (x, y int) {
	return micro.Size()
}

func Clamp(x, y int) (int, int) {
	sx, sy := micro.Size()

	if x < 0 {
		x += sx
		if x < 0 {
			x = 0
		}
	}
	if x >= sx {
		x = sx - 1
	}

	if y < 0 {
		y += sy
		if y < 0 {
			y = 0
		}
	}
	if y >= sy {
		y = sy - 1
	}

	return x, y
}

func clampBR(x, y int) (int, int) {
	sx, sy := micro.Size()

	if x < 0 {
		x += sx
	}
	if x < 1 {
		x = 1
	}
	if x > sx {
		x = sx
	}

	if y < 0 {
		y += sy
	}
	if y < 1 {
		y = 1
	}
	if y > sy {
		y = sy
	}

	return x, y
}

//------------------------------------------------------------------------------

func Clear() {
	for i := range micro.Text {
		micro.Text[i] = '\x00'
	}
	micro.TextUpdated = true
}

//------------------------------------------------------------------------------

func Peek(x, y int) byte {
	x, y = Clamp(x, y)
	return micro.Peek(x, y)
}

func Poke(x, y int, value byte) {
	x, y = Clamp(x, y)
	ov := micro.Peek(x, y)
	if value != ov {
		micro.Poke(x, y, value)
		micro.TextUpdated = true
	}
}

//------------------------------------------------------------------------------

func Print(x, y int, format string, a ...interface{}) {
	var w Writer
	w.left, w.top = Clamp(x, y)
	w.right, w.bottom = micro.Size()
	w.x, w.y = w.left, w.top

	fmt.Fprintf(&w, format, a...)
}

//------------------------------------------------------------------------------