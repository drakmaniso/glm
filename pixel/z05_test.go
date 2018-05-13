// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

type loop5 struct {
	canvas  pixel.CanvasID
	scene pixel.SceneID
	palette color.PaletteID

	points                    []coord.CR
	pointshidden, lineshidden bool
}

////////////////////////////////////////////////////////////////////////////////

func TestTest5(t *testing.T) {
	do(func() {
		defer cozely.Recover()

		l := loop5{}
		l.declare()

		input.Load(bindings)
		err := cozely.Run(&l)
		if err != nil {
			t.Error(err)
		}
	})
}

func (a *loop5) declare() {
	a.canvas = pixel.Canvas(pixel.Resolution(128, 128))
	a.scene = pixel.Scene()
	a.palette = color.PaletteFrom("graphics/shape1")

	a.points = []coord.CR{
		{4, 4},
		{4 + 1, 4 + 20},
		{4 + 1 + 20, 4 + 20 - 1},
		{16, 32},
	}
}

func (a *loop5) Enter() {
	input.ShowMouse(false)
	a.palette.Activate()
}

func (loop5) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (a *loop5) React() {
	if quit.Started(0) {
		cozely.Stop(nil)
	}

	if next.Started(0) {
		m := a.canvas.FromWindow(cursor.XY(0).CR())
		a.points = append(a.points, m)
	}

	if previous.Started(0) {
		if len(a.points) > 0 {
			a.points = a.points[:len(a.points)-1]
		}
	}

	a.pointshidden = scene1.Ongoing(0)
	a.lineshidden = scene2.Ongoing(0)
}

func (loop5) Update() {
}

func (a *loop5) Render() {
	a.canvas.Clear(1)
	m := a.canvas.FromWindow(cursor.XY(0).CR())
	a.scene.Triangles(2, a.points...)
	if !a.lineshidden {
		a.scene.Lines(5, a.points...)
		a.scene.Lines(13, a.points[len(a.points)-1], m)
	}
	if !a.pointshidden {
		for _, p := range a.points {
			a.scene.Point(8, p)
		}
		a.scene.Point(18, m)
	}
	a.canvas.Display(a.scene)
}
