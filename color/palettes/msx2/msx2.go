// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package msx2

import (
	"github.com/cozely/cozely/color"
)

var Palette = color.Palette()

func init() {
	for i := 1; i < 256; i++ {
		g, r, b := i>>5, (i&0x1C)>>2, i&0x3
		Palette.Set(uint8(i), color.LRGBA{
			float32(r) / 7.0,
			float32(g) / 7.0,
			float32(b) / 3.0,
			1.0,
		})
	}
}