// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import (
	"math/rand"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/basic"
	"github.com/drakmaniso/glam/color"
	"github.com/drakmaniso/glam/gfx"
	"github.com/drakmaniso/glam/mouse"
	"github.com/drakmaniso/glam/pixel"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/window"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Setup()
	if err != nil {
		glam.ShowError("setting up glam", err)
		return
	}

	err = setup()
	if err != nil {
		glam.ShowError("setting up the game", err)
		return
	}

	glam.Update = update
	glam.Draw = draw
	window.Handle = handler{}
	mouse.Handle = handler{}

	err = glam.Loop()
	if err != nil {
		glam.ShowError("running", err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline    *gfx.Pipeline
	perFrameUBO gfx.UniformBuffer
	rosesINBO   gfx.VertexBuffer
)

// Uniform buffer
var perFrame struct {
	ratio float32
	time  float32
}

// Instance Buffer

var roses [64]struct {
	position    plane.Coord `layout:"0" divisor:"1"`
	size        float32     `layout:"1"`
	numerator   int32       `layout:"2"`
	denominator int32       `layout:"3"`
	offset      float32     `layout:"4"`
	speed       float32     `layout:"5"`
}

//------------------------------------------------------------------------------

func setup() error {
	// Setup the pipeline
	pipeline = gfx.NewPipeline(
		gfx.Shader(glam.Path()+"shader.vert"),
		gfx.Shader(glam.Path()+"shader.frag"),
		gfx.VertexFormat(1, roses[:]),
		gfx.Topology(gfx.LineStrip),
	)
	gfx.Enable(gfx.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gfx.NewUniformBuffer(&perFrame, gfx.DynamicStorage)

	// Create the instance buffer
	randomizeRosesData()
	rosesINBO = gfx.NewVertexBuffer(roses[:], gfx.DynamicStorage)

	// Bind the instance buffer to the pipeline
	pipeline.Bind()
	rosesINBO.Bind(1, 0)
	pipeline.Unbind()

	return glam.Error("gfx", gfx.Err())
}

//------------------------------------------------------------------------------

func update(dt, _ float64) {
	perFrame.time += float32(dt)
}

func draw() {
	pipeline.Bind()
	gfx.ClearDepthBuffer(1.0)
	gfx.ClearColorBuffer(color.RGBA{0.9, 0.85, 0.80, 1.0})

	perFrameUBO.Bind(0)
	perFrameUBO.SubData(&perFrame, 0)
	gfx.DrawInstanced(0, nbPoints, int32(len(roses)))

	pipeline.Unbind()
}

//------------------------------------------------------------------------------

const nbPoints int32 = 512

func randomizeRosesData() {
	for i := 0; i < len(roses); i++ {
		roses[i].position.X = rand.Float32()*2.0 - 1.0
		roses[i].position.Y = rand.Float32()*2.0 - 1.0
		roses[i].size = rand.Float32()*0.20 + 0.1
		roses[i].numerator = rand.Int31n(16) + 1
		roses[i].denominator = rand.Int31n(16) + 1
		roses[i].offset = rand.Float32()*2.8 + 0.2
		roses[i].speed = 0.5 + 1.5*rand.Float32()
		if rand.Int31n(2) > 0 {
			roses[i].speed = -roses[i].speed
		}
	}
}

//------------------------------------------------------------------------------

// func rose(nbPoints int, num int, den int, offset float32) []perVertex {
// 	// var m = []perVertex{{plane.Coord{0.0, 0.0}, color.RGB{0.9, 0.9, 0.9}}}
// 	var m = []perVertex{}
// 	for i := den * nbPoints; i >= 0; i-- {
// 		var k = float32(num) / float32(den)
// 		var theta = float32(i) * 2 * math.Pi / float32(nbPoints)
// 		var r = (math.Cos(k*theta) + offset) / (1.0 + offset)
// 		var p = plane.Polar{r, theta}
// 		m = append(m, perVertex{p.Coord()})
// 	}
// 	return m
// }

//------------------------------------------------------------------------------

type handler struct {
	basic.WindowHandler
	basic.MouseHandler
}

func (h handler) WindowResized(s pixel.Coord, _ uint32) {
	var sx, sy = window.Size().Cartesian()
	perFrame.ratio = sy / sx
}

func (h handler) MouseButtonDown(b mouse.Button, _ int, _ uint32) {
	randomizeRosesData()
	rosesINBO.SubData(roses[:], 0)
}

//------------------------------------------------------------------------------
