// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package action

var bindings = map[string]binding{
	// Gamepad
	"Left Stick":    gpStick{},
	"Left Stick X":  gpStick{},
	"Left Stick Y":  gpStick{},
	"Right Stick":   gpStick{},
	"Right Stick X": gpStick{},
	"Right Stick Y": gpStick{},
	"Left Trigger":  gpTrigger{},
	"Right Trigger": gpTrigger{},
	"Left Bumper":   gpButton{},
	"Right Bumper":  gpButton{},
	"Dpad Up":       gpButton{},
	"Dpad Left":     gpButton{},
	"Dpad Down":     gpButton{},
	"Dpad Right":    gpButton{},
	"Button Y":      gpButton{},
	"Button X":      gpButton{},
	"Button A":      gpButton{},
	"Button B":      gpButton{},
	"Button Back":   gpButton{},
	"Button Start":  gpButton{},
	// Mouse
	"Mouse":             msPosition{},
	"Mouse X":           msPosition{},
	"Mouse Y":           msPosition{},
	"Mouse Left":        msButton{},
	"Mouse Middle":      msButton{},
	"Mouse Right":       msButton{},
	"Mouse Back":        msButton{},
	"Mouse Forward":     msButton{},
	"Mouse Button 6":    msButton{},
	"Mouse Button 7":    msButton{},
	"Mouse Button 8":    msButton{},
	"Mouse Button 9":    msButton{},
	"Mouse Button 10":   msButton{},
	"Mouse Button 11":   msButton{},
	"Mouse Button 12":   msButton{},
	"Mouse Button 13":   msButton{},
	"Mouse Button 14":   msButton{},
	"Mouse Button 15":   msButton{},
	"Mouse Button 16":   msButton{},
	"Mouse Button 17":   msButton{},
	"Mouse Button 18":   msButton{},
	"Mouse Button 19":   msButton{},
	"Mouse Button 20":   msButton{},
	"Mouse Scroll Up":   msButton{},
	"Mouse Scroll down": msButton{},
	// Keyboard
	"A":      kbKey{pos: KeyA},
	"B":      kbKey{pos: KeyB},
	"C":      kbKey{pos: KeyC},
	"D":      kbKey{pos: KeyD},
	"E":      kbKey{pos: KeyE},
	"F":      kbKey{pos: KeyF},
	"G":      kbKey{pos: KeyG},
	"H":      kbKey{pos: KeyH},
	"I":      kbKey{pos: KeyI},
	"J":      kbKey{pos: KeyJ},
	"K":      kbKey{pos: KeyK},
	"L":      kbKey{pos: KeyL},
	"M":      kbKey{pos: KeyM},
	"N":      kbKey{pos: KeyN},
	"O":      kbKey{pos: KeyO},
	"P":      kbKey{pos: KeyP},
	"Q":      kbKey{pos: KeyQ},
	"R":      kbKey{pos: KeyR},
	"S":      kbKey{pos: KeyS},
	"T":      kbKey{pos: KeyT},
	"U":      kbKey{pos: KeyU},
	"V":      kbKey{pos: KeyV},
	"W":      kbKey{pos: KeyW},
	"X":      kbKey{pos: KeyX},
	"Y":      kbKey{pos: KeyY},
	"Z":      kbKey{pos: KeyZ},
	"Up":     kbKey{pos: KeyUp},
	"Left":   kbKey{pos: KeyLeft},
	"Down":   kbKey{pos: KeyDown},
	"Right":  kbKey{pos: KeyRight},
	"Enter":  kbKey{pos: KeyReturn},
	"Space":  kbKey{pos: KeySpace},
	"Escape": kbKey{pos: KeyEscape},
}
