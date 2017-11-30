// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal/core"
	"github.com/drakmaniso/carol/key"
	"github.com/drakmaniso/carol/mouse"
	"github.com/drakmaniso/carol/screen"
)

//------------------------------------------------------------------------------

// Handlers implements default behavior for all events.
//
// It's an empty struct intended to be embedded in the user-defined GameLoop.
type Handlers struct{}

//------------------------------------------------------------------------------

// WindowShown does nothing.
func (h Handlers) WindowShown() {}

// WindowHidden does nothing.
func (h Handlers) WindowHidden() {}

// WindowResized does nothing.
func (h Handlers) WindowResized(s screen.Coord) {}

// WindowMinimized does nothing.
func (h Handlers) WindowMinimized() {}

// WindowMaximized does nothing.
func (h Handlers) WindowMaximized() {}

// WindowRestored does nothing.
func (h Handlers) WindowRestored() {}

// WindowMouseEnter does nothing.
func (h Handlers) WindowMouseEnter() {}

// WindowMouseLeave does nothing.
func (h Handlers) WindowMouseLeave() {}

// WindowFocusGained does nothing.
func (h Handlers) WindowFocusGained() {}

// WindowFocusLost does nothing.
func (h Handlers) WindowFocusLost() {}

// WindowQuit requests the game loop to stop.
func (h Handlers) WindowQuit() {
	core.QuitRequested = true
}

//------------------------------------------------------------------------------

// MouseMotion does nothing.
func (h Handlers) MouseMotion(rel screen.Coord, pos screen.Coord) {}

// MouseButtonDown does nothing.
func (h Handlers) MouseButtonDown(b mouse.Button, clicks int) {}

// MouseButtonUp does nothing.
func (h Handlers) MouseButtonUp(b mouse.Button, clicks int) {}

// MouseWheel does nothing.
func (h Handlers) MouseWheel(w screen.Coord) {}

//------------------------------------------------------------------------------

// KeyDown requests the game loop to stop if Escape is pressed.
func (h Handlers) KeyDown(l key.Label, p key.Position) {
	switch l {
	case key.LabelEscape:
		core.QuitRequested = true
	case key.LabelF11:
		core.ToggleFullscreen()
	}
}

// KeyUp does nothing.
func (h Handlers) KeyUp(l key.Label, p key.Position) {
}

//------------------------------------------------------------------------------
