// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel

import (
	"github.com/drakmaniso/glam/x/gl"
)

//------------------------------------------------------------------------------

const (
	maxCommandCount = 1024
	maxParamCount   = 8
)

//------------------------------------------------------------------------------

var (
	commands   []gl.DrawIndirectCommand
	parameters []int16
)

var screenUniforms struct {
	PixelSize struct{ X, Y float32 }
}

var blitUniforms struct {
	ScreenSize struct{ X, Y float32 }
}

//------------------------------------------------------------------------------

var (
	pipeline      *gl.Pipeline
	screenUBO     gl.UniformBuffer
	commandsICBO  gl.IndirectBuffer
	parametersTBO gl.BufferTexture
	mappingsTBO   gl.BufferTexture
	picturesTA    gl.TextureArray2D
)

var (
	blitPipeline *gl.Pipeline
	blitUBO      gl.UniformBuffer
	blitTexture  gl.Sampler
)

//------------------------------------------------------------------------------
