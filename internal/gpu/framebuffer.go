// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gpu

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/pixel"
)

/*
#include <stdlib.h>
#include "glad.h"

static inline GLuint CreateFramebuffer(GLsizei width, GLsizei height) {
	//TODO: delte previous FBO and textures

	GLuint fbo;
	glCreateFramebuffers(1, &fbo);

	GLuint ct; // Color texture
	glCreateTextures(GL_TEXTURE_2D, 1, &ct);
	glTextureStorage2D(ct, 1, GL_RGB8, width, height);
	glTextureParameteri(ct, GL_TEXTURE_MIN_FILTER, GL_NEAREST);
	glTextureParameteri(ct, GL_TEXTURE_MAG_FILTER, GL_NEAREST);

	GLuint dt; // Depth texture
	glCreateTextures(GL_TEXTURE_2D, 1, &dt);
	glTextureStorage2D(dt, 1, GL_DEPTH_COMPONENT16, width, height);
	glTextureParameteri(dt, GL_TEXTURE_MIN_FILTER, GL_NEAREST); //TODO: remove?
	glTextureParameteri(dt, GL_TEXTURE_MAG_FILTER, GL_NEAREST); //TODO: remove?

	glNamedFramebufferTexture(fbo, GL_COLOR_ATTACHMENT0, ct, 0);
	glNamedFramebufferTexture(fbo, GL_DEPTH_ATTACHMENT, dt, 0);

	glNamedFramebufferDrawBuffer(fbo, GL_COLOR_ATTACHMENT0);

	glViewport(0, 0, width, height);
	glBindFramebuffer(GL_FRAMEBUFFER, fbo);

	return fbo;
}

static inline void BlitFramebuffer(GLint width, GLint height) {
	glBindFramebuffer(GL_DRAW_FRAMEBUFFER, 0);
	glBlitFramebuffer(
		0, 0, width, height, // fbo
		0, 0, width, height, // screen
		GL_COLOR_BUFFER_BIT, GL_NEAREST
	);
	//TODO: bind FBO back (but where?)
}
*/
import "C"

//------------------------------------------------------------------------------

// ResizeWindow sets the size in pixels of the GL viewport.
//
// This function is automatically called each time the window is resized.
func CreateFramebuffer(size pixel.Coord) {
	framebuffer = C.CreateFramebuffer(C.GLsizei(size.X), C.GLsizei(size.Y))
}

var framebuffer C.GLuint

//------------------------------------------------------------------------------

// BlitFramebuffer swaps the double-buffer.
func BlitFramebuffer(size pixel.Coord) {
	C.BlitFramebuffer(C.GLint(size.X), C.GLint(size.Y))
}

//------------------------------------------------------------------------------