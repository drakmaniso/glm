// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package internal

import (
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

/*
#include "sdl.h"

#define PEEP_SIZE 128

SDL_Event Events[PEEP_SIZE];

int PeepEvents()
{
  SDL_PumpEvents();
  int n = SDL_PeepEvents(Events, PEEP_SIZE, SDL_GETEVENT, SDL_FIRSTEVENT, SDL_LASTEVENT);
  return n;
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// GameLoop (identic to cozely.GameLoop).
type GameLoop interface {
	Enter()
	Leave()

	React()
	Update()
	Render()
}

////////////////////////////////////////////////////////////////////////////////

// ProcessEvents processes and dispatches all events.
func ProcessEvents(win struct {
	Resize  func()
	Hide    func()
	Show    func()
	Focus   func()
	Unfocus func()
	Quit    func()
}) {
	more := true
	for more && !QuitRequested {
		n := peepEvents()

		for i := 0; i < n && !QuitRequested; i++ {
			e := eventAt(i)
			dispatch(e, win)
		}
		more = n >= C.PEEP_SIZE
	}

	var mx, my C.int
	C.SDL_GetRelativeMouseState(&mx, &my)
	MouseDeltaX += int16(mx)
	MouseDeltaY += int16(my)
	btn := C.SDL_GetMouseState(&mx, &my)
	MousePositionX = int16(mx)
	MousePositionY = int16(my)
	MouseButtons = uint32(btn)
}

func dispatch(e unsafe.Pointer, win struct {
	Resize  func()
	Hide    func()
	Show    func()
	Focus   func()
	Unfocus func()
	Quit    func()
}) {
	switch ((*C.SDL_CommonEvent)(e))._type {
	case C.SDL_QUIT:
		win.Quit()
	// Window Events
	case C.SDL_WINDOWEVENT:
		e := (*C.SDL_WindowEvent)(e)
		switch e.event {
		case C.SDL_WINDOWEVENT_NONE:
			// Ignore
		case C.SDL_WINDOWEVENT_SHOWN:
			win.Show()
		case C.SDL_WINDOWEVENT_HIDDEN:
			win.Hide()
		case C.SDL_WINDOWEVENT_EXPOSED:
			// Ignore
		case C.SDL_WINDOWEVENT_MOVED:
			// Ignore
		case C.SDL_WINDOWEVENT_RESIZED:
			Window.Width = int16(e.data1)
			Window.Height = int16(e.data2)
			PixelResize()
			win.Resize()
		case C.SDL_WINDOWEVENT_SIZE_CHANGED:
			//TODO
		case C.SDL_WINDOWEVENT_MINIMIZED:
			//TODO: check that Hide is enough
		case C.SDL_WINDOWEVENT_MAXIMIZED:
			// Ingnore
		case C.SDL_WINDOWEVENT_RESTORED:
			//TODO: check that Show is enough
		case C.SDL_WINDOWEVENT_ENTER:
			HasMouseFocus = true
			// C.SDL_ShowCursor(C.SDL_DISABLE)
		case C.SDL_WINDOWEVENT_LEAVE:
			HasMouseFocus = false
			// C.SDL_ShowCursor(C.SDL_ENABLE)
		case C.SDL_WINDOWEVENT_FOCUS_GAINED:
			HasFocus = true
			win.Focus()
		case C.SDL_WINDOWEVENT_FOCUS_LOST:
			HasFocus = false
			win.Unfocus()
		case C.SDL_WINDOWEVENT_CLOSE:
			// Ignore
		default:
			//TODO: log.Print("unknown window event")
		}
	// Mouse Events
	case C.SDL_MOUSEWHEEL:
		e := (*C.SDL_MouseWheelEvent)(e)
		var d int16 = 1
		if e.direction == C.SDL_MOUSEWHEEL_FLIPPED {
			d = -1
		}
		MouseWheelX += int16(e.x) * d
		MouseWheelY += int16(e.y) * d
	//TODO: Joystick Events
	case C.SDL_JOYAXISMOTION:
	case C.SDL_JOYBALLMOTION:
	case C.SDL_JOYHATMOTION:
	case C.SDL_JOYBUTTONDOWN:
	case C.SDL_JOYBUTTONUP:
	case C.SDL_JOYDEVICEADDED:
	case C.SDL_JOYDEVICEREMOVED:
	//TODO: Controller Events
	case C.SDL_CONTROLLERAXISMOTION:
	case C.SDL_CONTROLLERBUTTONDOWN:
	case C.SDL_CONTROLLERBUTTONUP:
	case C.SDL_CONTROLLERDEVICEADDED:
	case C.SDL_CONTROLLERDEVICEREMOVED:
	case C.SDL_CONTROLLERDEVICEREMAPPED:
	//TODO: Audio Device Events
	case C.SDL_AUDIODEVICEADDED:
	case C.SDL_AUDIODEVICEREMOVED:
	default:
		//TODO: log.Print("unknown SDL event:", ((*C.SDL_CommonEvent)(e))._type)
	}
}

// peepEvents fill the event buffer and returns the number of events fetched.
func peepEvents() int {
	return int(C.PeepEvents())
}

// EventAt returns a pointer to an event in the event buffer.
func eventAt(i int) unsafe.Pointer {
	return unsafe.Pointer(&C.Events[i])
}

////////////////////////////////////////////////////////////////////////////////

// SDLQuit is called when the game loop stops.
func SDLQuit() {
	C.SDL_Quit()
}

////////////////////////////////////////////////////////////////////////////////
