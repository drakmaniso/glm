// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package mouse

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/screen"
)

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() screen.Coord {
	return core.MousePosition
}

// Delta returns the mouse position relative to the last call of Delta.
func Delta() screen.Coord {
	result := core.MouseDelta
	core.MouseDelta.X, core.MouseDelta.Y = 0, 0
	return result
}

// SetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelativeMode(enabled bool) error {
	return core.MouseSetRelativeMode(enabled)
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return core.MouseGetRelativeMode()
}

//------------------------------------------------------------------------------
