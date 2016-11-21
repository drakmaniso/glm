// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gfx

//------------------------------------------------------------------------------

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
#include "glad.h"

GLuint NewBuffer(GLsizeiptr size, void* data, GLenum flags) {
	GLuint b;
	glCreateBuffers(1, &b);
	glNamedBufferStorage(b, size, data, flags);
	return b;
}

static inline void UpdateBuffer(GLuint buffer, GLintptr offset, GLsizei size, void *data) {
	glNamedBufferSubData(buffer, offset, size, data);
}

static inline void BindUniform(GLuint binding, GLuint buffer) {
	glBindBufferBase(GL_UNIFORM_BUFFER, binding, buffer);
}

static inline void BindVertex(GLuint binding, GLuint buffer, GLintptr offset, GLsizei stride) {
	glBindVertexBuffer(binding, buffer, offset, stride);
}
*/
import "C"

//------------------------------------------------------------------------------

// A UniformBuffer is a block of memory owned by the GPU.
type UniformBuffer struct {
	object C.GLuint
	stride uintptr
}

// NewUniformBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a pointer
// to a struct of pure values (no nested references). In all cases the size of
// the buffer is fixed at creation.
func NewUniformBuffer(data interface{}, f bufferFlags) (UniformBuffer, error) {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		return UniformBuffer{}, err
	}
	var ub UniformBuffer
	ub.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLenum(f))
	//TODO: error handling
	ub.stride = 0
	return ub, nil
}

// Update the buffer with data, starting at a specified offset.
//
// It is your responsability to ensure that the size of data plus the offset
// does not exceed the buffer size.
func (ub *UniformBuffer) Update(data interface{}, atOffset uintptr) error {
	p, s, err := pointerAndSizeOf(data)
	if err != nil {
		return err
	}
	C.UpdateBuffer(ub.object, C.GLintptr(atOffset), C.GLsizei(s), p)
	return nil
}

// Bind to a uniform binding index.
//
// This index should correspond to one indicated by a layout qualifier in the
// shaders.
func (ub *UniformBuffer) Bind(binding uint32) {
	C.BindUniform(C.GLuint(binding), ub.object)
}

//------------------------------------------------------------------------------

// A VertexBuffer is a block of memory owned by the GPU.
type VertexBuffer struct {
	object C.GLuint
	stride uintptr
}

// NewVertexBuffer asks the GPU to allocate a new block of memory.
//
// If data is a uinptr, it is interpreted as the desired size for the buffer (in
// bytes), and the content is not initialized. Otherwise data must be a slice of
// pure values (no nested references). In all cases the size of the buffer is
// fixed at creation.
func NewVertexBuffer(data interface{}, f bufferFlags) (VertexBuffer, error) {
	p, s, st, err := pointerSizeAndStrideOf(data)
	if err != nil {
		return VertexBuffer{}, err
	}
	var vb VertexBuffer
	vb.object = C.NewBuffer(C.GLsizeiptr(s), p, C.GLenum(f))
	//TODO: error handling
	vb.stride = st
	return vb, nil
}

// Update the buffer with data, starting at a specified offset. It is your
// responsability to ensure that the size of data plus the offset does not
// exceed the buffer size.
func (vb *VertexBuffer) Update(data interface{}, atOffset uintptr) error {
	p, s, st, err := pointerSizeAndStrideOf(data)
	if err != nil {
		return err
	}
	C.UpdateBuffer(vb.object, C.GLintptr(atOffset), C.GLsizei(s), p)
	if st != 0 {
		// In case the stride was not specified at buffer creation, or the new data
		// has a different stride.
		vb.stride = st
	}
	return nil
}

// Bind to a vertex buffer binding index.
//
// The buffer should use the same struct type than the one used in the
// corresponding call to Pipeline.VertexFormat.
func (vb *VertexBuffer) Bind(binding uint32, offset uintptr) {
	C.BindVertex(C.GLuint(binding), vb.object, C.GLintptr(offset), C.GLsizei(vb.stride))
}

//------------------------------------------------------------------------------

type bufferFlags uint32

// Flags for buffer creation.
const (
	StaticStorage  bufferFlags = 0x0000 // Content will not be updated
	MapRead        bufferFlags = 0x0001 // Data store will be mapped for reading
	MapWrite       bufferFlags = 0x0002 // Data store will be mapped for writing
	MapPersistent  bufferFlags = 0x0040 // Data store will be accessed by both application and GPU while mapped
	MapCoherent    bufferFlags = 0x0080 // No synchronization needed when persistently mapped
	DynamicStorage bufferFlags = 0x0100 // Content will be updated
	ClientStorage  bufferFlags = 0x0200 // Prefer storage on application side
)

func pointerAndSizeOf(data interface{}) (ptr unsafe.Pointer, size uintptr, err error) {
	var p unsafe.Pointer
	var s uintptr
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		p = nil
		s = uintptr(v.Uint())
	case reflect.Ptr:
		p = unsafe.Pointer(v.Pointer())
		//TODO: check if pointer refer to a struct
		s = v.Elem().Type().Size()
	default:
		return nil, 0, fmt.Errorf("%s instead of point-to-struct or uinptr", reflect.TypeOf(data).Kind())
	}
	return p, s, nil
}

func pointerSizeAndStrideOf(data interface{}) (ptr unsafe.Pointer, size uintptr, stride uintptr, err error) {
	var p unsafe.Pointer
	var s uintptr
	var st uintptr
	v := reflect.ValueOf(data)
	k := v.Kind()
	switch k {
	case reflect.Uintptr:
		p = nil
		s = uintptr(v.Uint())
		st = 0
	case reflect.Slice:
		l := v.Len()
		if l == 0 {
			return nil, 0, 0, fmt.Errorf("buffer data cannot be an empty slice")
		}
		p = unsafe.Pointer(v.Pointer())
		st = v.Index(0).Type().Size()
		s = uintptr(l) * st
	default:
		return nil, 0, 0, fmt.Errorf("%s instead of slice or uintptr", reflect.TypeOf(data).Kind())
	}
	return p, s, st, nil
}

//------------------------------------------------------------------------------
