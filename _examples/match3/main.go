// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"math/rand"
	"time"

	"github.com/drakmaniso/glam/palette"

	"github.com/drakmaniso/glam/key"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"

	"github.com/drakmaniso/glam/_examples/match3/ecs"
	"github.com/drakmaniso/glam/_examples/match3/grid"
)

//------------------------------------------------------------------------------

var tilesPict [8]struct {
	normal, big pixel.Picture
}

var current grid.Position

var screen = pixel.NewCanvas(pixel.TargetResolution(180, 180))

//------------------------------------------------------------------------------

func main() {
	setup()

	glam.Configure(
		glam.Title("Match 3"),
	)

	err := glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
	}
}

//------------------------------------------------------------------------------

func setup() error {
	err := palette.Load("graphics/blue")
	if err != nil {
		return err
	}

	// palette.Change("MSX2")

	for i, n := range []string{
		"red",
		"yellow",
		"green",
		"blue",
		"violet",
		"pink",
		"dark",
		"multi",
	} {
		tilesPict[i].normal = pixel.NewPicture("graphics/" + n)
		tilesPict[i].big = pixel.NewPicture("graphics/" + n + "_big")
	}

	current = grid.Nowhere()

	grid.Setup(8, 8)
	grid.Fill(newTile)

	return nil
}

//------------------------------------------------------------------------------

type loop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func newTile() ecs.Entity {
	e := ecs.New(ecs.Color)
	c := color(rand.Int31n(7))
	// if rand.Int31n(16) == 0 {
	// 	c = 7
	// }
	colors[e] = c

	return e
}

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) MouseButtonDown(_ mouse.Button, _ int) {
	m := screen.Mouse()
	current = grid.PositionAt(m)
	if current != grid.Nowhere() {
		e := grid.At(current)
		n := 0
		f := func(e ecs.Entity) {
			print(grid.PositionOf(e).String(), " ")
			n++
		}
		grid.PositionOf(e).TestAndMark(testMatch, f)
		println("-> ", n)
	}
}

func testMatch(e1, e2 ecs.Entity) bool {
	if !e1.Has(ecs.Color) || !e2.Has(ecs.Color) {
		return false
	}
	c1 := colors[e1]
	c2 := colors[e2]
	return c1 == c2
}

func (loop) MouseButtonUp(_ mouse.Button, _ int) {
	current = grid.Nowhere()
}

//------------------------------------------------------------------------------

func (lp loop) KeyDown(l key.Label, p key.Position) {
	switch l {
	case key.LabelSpace:
		f := func(e ecs.Entity) {
			if !e.Has(ecs.MatchFlag) {
				print(grid.PositionOf(e).String(), " ")
				e.Add(ecs.MatchFlag)
			}
		}
		grid.TestAndMark(testMatch, f)
		println()

	default:
		lp.Handlers.KeyDown(l, p)
	}
}

//------------------------------------------------------------------------------

func (loop) WindowResized(_, _ int32) {
	w, h := screen.Size().X, screen.Size().Y
	grid.ScreenResized(w, h)
}

//------------------------------------------------------------------------------
