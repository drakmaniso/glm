// Adapted by Laurent Moussault (2018) from C code:
// http://www.ic.unicamp.br/~stolfi/EXPORT/software/c/2000-05-04/libquad/
//
// See
//
//   "Primitives for the Manipulation of General Subdivisions
//   and the Computation of Voronoi Diagrams"
//
//   p. Guibas, J. Stolfi, ACM TOG, April 1985
//
// Originally written by Jim Roth (DEC CADM Advanced Group) on May 1986.
// Modified by J. Stolfi on April 1993.
//
// Original Copyright notice:
//
// Copyright 1996 Institute of Computing, Unicamp.
//
// Permission to use this software for any purpose is hereby granted,
// provided that any substantial copy or mechanically derived version
// of this file that is made available to other parties is accompanied
// by this copyright notice in full, and is distributed under these same
// terms.
//
// DISCLAIMER: This software is provided "as is" with no explicit or
// implicit warranty of any kind.  Neither the authors nor their
// employers can be held responsible for any losses or damages
// that might be attributed to its use.
//
// End of copyright notice.

package quadedge

//------------------------------------------------------------------------------

// A Pool is an allocator for Edge objects.
type Pool struct {
	next     []edgeID
	data     []uint32
	marks    []uint32
	free     []edgeID
	capacity uint32
	nextMark uint32
}

//------------------------------------------------------------------------------

func NewPool(capacity uint32) *Pool {
	return &Pool{
		next:     make([]edgeID, 0, 4*capacity),
		data:     make([]uint32, 0, 4*capacity),
		marks:    make([]uint32, 0, capacity),
		capacity: capacity,
		nextMark: 1,
	}
}

//------------------------------------------------------------------------------

func (p *Pool) NewEdge() Edge {

	//TODO: implement the free list

	sz := uint32(len(p.next) + 4)
	if sz > p.capacity {
		//TODO: grow the slices
		panic("growing QuadEdges not implemented")
	}

	// Allocate the new quad
	e := edgeID(len(p.next))
	p.next = p.next[:sz]
	p.data = p.data[:sz]

	// Initialize the quad
	p.next[e] = e
	p.data[e] = 0xFFFFFFFF
	p.next[e.sym()] = e.sym()
	p.data[e.sym()] = 0xFFFFFFFF
	p.next[e.rot()] = e.tor()
	p.data[e.rot()] = 0xFFFFFFFF
	p.next[e.tor()] = e.rot()
	p.data[e.tor()] = 0xFFFFFFFF

	return Edge{pool: p, id: e}
}

//------------------------------------------------------------------------------

func Splice(a, b Edge) {
	p := a.pool
	alpha := p.next[a.id].rot()
	beta := p.next[b.id].rot()

	p.next[a.id], p.next[b.id] = p.next[b.id], p.next[a.id]

	p.next[alpha], p.next[beta] = p.next[beta], p.next[alpha]
}

//------------------------------------------------------------------------------

func (p *Pool) Free(e Edge) {
	f := e.Sym()
	if p.next[e.id] != e.id {
		Splice(e, e.OrigPrev())
	}
	if p.next[f.id] != f.id {
		Splice(f, f.OrigPrev())
	}

	p.next[e.id] = noEdge
	p.data[e.id] = 0xFFFFFFFF
	p.next[e.id.rot()] = noEdge
	p.data[e.id.rot()] = 0xFFFFFFFF
	p.next[e.id.sym()] = noEdge
	p.data[e.id.sym()] = 0xFFFFFFFF
	p.next[e.id.tor()] = noEdge
	p.data[e.id.tor()] = 0xFFFFFFFF

	p.marks[e.id>>2] = 0

	p.free = append(p.free, e.id)
}

//------------------------------------------------------------------------------
