// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// FilePath of the executable (uses os-dependant separator).
var FilePath string

// Path of the executable (uses slash separators, and ends with one).
var Path string

// Title of the game
var Title = "Cozely"

////////////////////////////////////////////////////////////////////////////////

// Config holds the initial configuration of the game.
var Config = struct {
	Debug          bool
	WindowSize     [2]int16
	Display        int
	Fullscreen     bool
	FullscreenMode string
	VSync          bool
}{
	Debug:          false,
	WindowSize:     [2]int16{1280, 720},
	Display:        0,
	Fullscreen:     false,
	FullscreenMode: "Desktop",
	VSync:          true,
}

////////////////////////////////////////////////////////////////////////////////

var (
	Log   logger = log.New(os.Stderr, "", log.Ltime|log.Lmicroseconds)
	Debug logger = nolog{}
)

////////////////////////////////////////////////////////////////////////////////

// Running is true once the game loop is started.
var Running = false

// GameTime is the current time.
var GameTime float64

// UpdateStep is the fixed time between calls to Update
var UpdateStep = float64(1.0 / 50)

var (
	// RenderDelta is the time elapsed between current and previous frames.
	RenderDelta float64
	// UpdateLag is the time accumulator used to decorrelate render frames from
	// updates.
	UpdateLag float64
)

////////////////////////////////////////////////////////////////////////////////

// QuitRequested makes the game loop stop if true.
var QuitRequested = false

////////////////////////////////////////////////////////////////////////////////

// Window is the game window.
var Window struct {
	window        *C.SDL_Window
	context       C.SDL_GLContext
	Width, Height int16
	Multisample   int32
}

// Focus state
var (
	HasFocus      bool
	HasMouseFocus bool
)

////////////////////////////////////////////////////////////////////////////////

// Loop holds the active looper.
//
// Note: The variable is set with cozely.Loop.
var Loop GameLoop

////////////////////////////////////////////////////////////////////////////////

// KeyState holds the pressed state of all keys, indexed by position.
var KeyState [512]bool //TODO: remove

////////////////////////////////////////////////////////////////////////////////
