// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

import (
	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/colour"
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

func main() {
	err := glam.Run(loop{})
	if err != nil {
		glam.ShowError(err)
		return
	}
}

//------------------------------------------------------------------------------

// OpenGL objects
var (
	pipeline *gl.Pipeline
)

//------------------------------------------------------------------------------

type loop struct {
	glam.EmptyLoop
}

//------------------------------------------------------------------------------

func (loop) Enter() error {
	// Create and configure the pipeline
	pipeline = gl.NewPipeline(
		gl.Shader(glam.Path()+"shader.vert"),
		gl.Shader(glam.Path()+"shader.frag"),
		gl.Topology(gl.Triangles),
	)

	return glam.Error("gfx", gl.Err())
}

//------------------------------------------------------------------------------

func (l loop) WindowResized(w, h int32) {
	gl.Viewport(0, 0, w, h)
}

//------------------------------------------------------------------------------

func (loop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (loop) Draw() error {
	pipeline.Bind()
	gl.ClearColorBuffer(colour.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	pipeline.Unbind()

	return gl.Err()
}

//------------------------------------------------------------------------------
