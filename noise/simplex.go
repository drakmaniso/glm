// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package noise

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/glam/math"
	"github.com/drakmaniso/glam/plane"
	"github.com/drakmaniso/glam/space"
)

//------------------------------------------------------------------------------

// Skewing and unskewing factors for 2, 3, and 4 dimensions
const sqrt3 = 1.732050807568877
const sqrt5 = 2.23606797749979
const f2 = 0.5 * (sqrt3 - 1.0)
const g2 = (3.0 - sqrt3) / 6.0
const f3 = 1.0 / 3.0
const g3 = 1.0 / 6.0
const f4 = (sqrt5 - 1.0) / 4.0
const g4 = (5.0 - sqrt5) / 20.0

//------------------------------------------------------------------------------

// Simplex2DAtCartesian returns the 2D simplex noise at position p.
func Simplex2DCartesianAt(p plane.Coord) float32 {
	// Source: "Simplex Noise Demystified" by Stefan Gustavson
	// http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
	// and
	// http://webstaff.itn.liu.se/~stegu/simplexnoise/SimplexNoise.java

	// Noise contributions from the three corners
	var n0, n1, n2 float32

	// Skew the input space to determine which simplex cell we're in.
	var s = (p.X + p.Y) * f2
	var i = math.Floor(p.X + s)
	var j = math.Floor(p.Y + s)

	// Unskew the cell origin back to (x,y) space.
	var t = (i + j) * g2
	var X0 = i - t
	var Y0 = j - t
	var x0 = p.X - X0 // The x,y distances from the cell origin
	var y0 = p.Y - Y0

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in.
	var i1, j1 int32 // Offsets for second (middle) corner of simplex in (i,j) coords
	if x0 > y0 {
		// lower triangle, XY order: (0,0)->(1,0)->(1,1)
		i1 = 1
		j1 = 0
	} else {
		// upper triangle, YX order: (0,0)->(0,1)->(1,1)
		i1 = 0
		j1 = 1
	}

	// A step of (1,0) in (i,j) means a step of (1-c,-c) in (x,y), and
	// a step of (0,1) in (i,j) means a step of (-c,1-c) in (x,y), where
	// c = (3-sqrt(3))/6
	var x1 = x0 - float32(i1) + g2 // Offsets for middle corner in (x,y) unskewed coords
	var y1 = y0 - float32(j1) + g2
	var x2 = x0 - 1.0 + 2.0*g2 // Offsets for last corner in (x,y) unskewed coords
	var y2 = y0 - 1.0 + 2.0*g2

	// Work out the hashed gradient indices of the three simplex corners
	var ii = int32(i) & 255
	var jj = int32(j) & 255
	var gl = int32(len(Gradient3D))
	var gi0 = permutation[ii+permutation[jj]] % gl
	var gi1 = permutation[ii+i1+permutation[jj+j1]] % gl
	var gi2 = permutation[ii+1+permutation[jj+1]] % gl

	// Calculate the contribution from the three corners

	var t0 = 0.5 - x0*x0 - y0*y0
	if t0 < 0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * (Gradient3D[gi0].Dot(space.Coord{x0, y0, 0})) // (x,y) of grad3 used for 2D gradient
	}

	var t1 = 0.5 - x1*x1 - y1*y1
	if t1 < 0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * (Gradient3D[gi1].Dot(space.Coord{x1, y1, 0}))
	}

	var t2 = 0.5 - x2*x2 - y2*y2
	if t2 < 0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * (Gradient3D[gi2].Dot(space.Coord{x2, y2, 0}))
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].

	return 70.0 * (n0 + n1 + n2)
}

//------------------------------------------------------------------------------

// Simplex2DAtAxial returns the 2D simplex noise at position (q, r), expressed
// in axial coordinates.
func Simplex2DAxialAt(q, r float32, grad []plane.Coord) float32 {
	// Source: "Simplex Noise Demystified" by Stefan Gustavson
	// http://www.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
	// and
	// http://webstaff.itn.liu.se/~stegu/simplexnoise/SimplexNoise.java

	// Noise contributions from the three corners
	var n0, n1, n2 float32

	// Determine which simplex cell we're in.
	var q0 = math.Floor(q)
	var r0 = math.Floor(r)

	// Unskew the cell origin back to (x,y) space.
	var dq0 = q - float32(q0) // The x,y distances from the cell origin
	var dr0 = r - float32(r0)

	// Determine which simplex we are in.
	// var q1, r1 int32 // Offsets for second (middle) corner of simplex
	var v int32
	if dq0+dr0 < 1.0 {
		// lower triangle, XY order: (0,0)->(1,0)->(1,1)
		v = 0
		// q1 = 1
		// r1 = 0
	} else {
		// upper triangle, YX order: (0,0)->(0,1)->(1,1)
		v = 1
		// q1 = 0
		// r1 = 1
	}

	// Work out the hashed gradient indices of the three simplex corners
	var qq = int32(q0) & 255
	var rr = int32(r0) & 255
	var gl = int32(len(grad))
	var gi0 = permutation[qq+v+permutation[rr+v]] % gl
	var gi1 = permutation[qq+1+permutation[rr]] % gl
	var gi2 = permutation[qq+permutation[rr+1]] % gl

	var x0 = (dq0 - float32(v)) + 0.5*(dr0-float32(v))
	var y0 = (dr0 - float32(v)) * 0.5 * sqrt3
	var x1 = (dq0 - 1) + 0.5*(dr0)
	var y1 = (dr0) * 0.5 * sqrt3
	var x2 = (dq0) + 0.5*(dr0-1)
	var y2 = (dr0 - 1) * 0.5 * sqrt3

	// Calculate the contribution from the three corners

	var t0 = 0.5 - x0*x0 - y0*y0
	if t0 < 0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * (grad[gi0].Dot(plane.Coord{x0, y0})) // (x,y) of grad3 used for 2D gradient
	}

	var t1 = 0.5 - x1*x1 - y1*y1
	if t1 < 0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * (grad[gi1].Dot(plane.Coord{x1, y1}))
	}

	var t2 = 0.5 - x2*x2 - y2*y2
	if t2 < 0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * (grad[gi2].Dot(plane.Coord{x2, y2}))
	}

	// Add contributions from each corner to get the final noise value.
	// The result is scaled to return values in the interval [-1,1].

	return 70.0 * (n0 + n1 + n2)
}

//------------------------------------------------------------------------------