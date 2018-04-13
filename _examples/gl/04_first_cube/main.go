// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/coord/space"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

////////////////////////////////////////////////////////////////////////////////

var (
	quit   = input.Bool("Quit")
	rotate = input.Bool("Rotate")
	move   = input.Bool("Move")
	zoom   = input.Bool("Zoom")
)

var context = input.Context("Default", quit, rotate, move, zoom)

var bindings = input.Bindings{
	"Default": {
		"Quit":   {"Escape"},
		"Rotate": {"Mouse Left"},
		"Move":   {"Mouse Right"},
		"Zoom":   {"Mouse Middle"},
	},
}

////////////////////////////////////////////////////////////////////////////////

// OpenGL objects
var (
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer
)

// Uniform buffer
var perObject struct {
	screenFromObject space.Matrix
}

// Vertex buffer
type mesh []struct {
	position coord.XYZ  `layout:"0"`
	color    color.LRGB `layout:"1"`
}

// Transformation matrices
var (
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matrix
)

// Cube state
var (
	position   coord.XYZ
	yaw, pitch float32
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	cozely.Configure(cozely.Multisample(8))
	cozely.Events.Resize = resize
	err := cozely.Run(loop{})
	if err != nil {
		cozely.ShowError(err)
		return
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

////////////////////////////////////////////////////////////////////////////////

func (loop) Enter() error {
	bindings.Load()
	context.Activate(1)

	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader.vert"),
		gl.Shader(cozely.Path()+"shader.frag"),
		gl.VertexFormat(0, mesh{}),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
		gl.DepthComparison(gl.LessOrEqual),
	)
	gl.Enable(gl.FramebufferSRGB)

	// Create the uniform buffer
	perFrameUBO = gl.NewUniformBuffer(&perObject, gl.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gl.NewVertexBuffer(cube(), 0)

	// Initialize worldFromObject and viewFromWorld matrices
	position = coord.XYZ{0, 0, 0}
	yaw = -0.6
	pitch = 0.3
	computeWorldFromObject()
	computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	pipeline.Bind()
	vbo.Bind(0, 0)
	pipeline.Unbind()

	return cozely.Error("gl", gl.Err())
}

func (loop) Leave() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() error {
	pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	perObject.screenFromObject =
		screenFromView.
			Times(viewFromWorld).
			Times(worldFromObject)
	perFrameUBO.SubData(&perObject, 0)
	perFrameUBO.Bind(0)

	gl.Draw(0, 6*2*3)

	pipeline.Unbind()

	return gl.Err()
}

////////////////////////////////////////////////////////////////////////////////

func resize() {
	s := cozely.WindowSize()
	gl.Viewport(0, 0, int32(s.C), int32(s.R))
	r := float32(s.C) / float32(s.R)
	screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
}

////////////////////////////////////////////////////////////////////////////////
