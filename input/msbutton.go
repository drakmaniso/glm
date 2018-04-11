// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package input

import "github.com/drakmaniso/glam/internal"

type msButton struct {
	button        MouseButton
	target        Action
	just, pressed bool
}

// MouseButton identifies a mouse button
type MouseButton uint32

// MouseButton constants
const (
	MouseLeft MouseButton = 1 << iota
	MouseMiddle
	MouseRight
	MouseBack
	MouseForward
	Mouse6
	Mouse7
	Mouse8
)

func (a *msButton) bind(c Context, target Action) {
	aa := *a
	aa.target = target
	devices.bindings[kbmouse][c] =
		append(devices.bindings[kbmouse][c], &aa)
}

func (a *msButton) activate(d Device) {
	a.target.activate(d, a)
}

func (a *msButton) asBool() (just bool, value bool) {
	v := (MouseButton(internal.MouseButtons) & a.button) != 0
	a.just = (v != a.pressed) //TODO: no need to store?
	a.pressed = v
	return a.just, a.pressed
}