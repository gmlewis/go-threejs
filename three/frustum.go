// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Frustum represents a frustum.
type Frustum struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Frustum) JSObject() *js.Object { return f.p }

// Frustum returns a Frustum object.
func (t *Three) Frustum() *Frustum {
	p := t.ctx.Get("Frustum")
	return &Frustum{p: p}
}

// New returns a new Frustum object.
func (f *Frustum) New(p0, p1, p2, p3, p4, p5 float64) *Frustum {
	p := f.p.New(p0, p1, p2, p3, p4, p5)
	return &Frustum{p: p}
}

// Set TODO description.
func (f *Frustum) Set(p0, p1, p2, p3, p4, p5 float64) *Frustum {
	f.p.Call("set", p0, p1, p2, p3, p4, p5)
	return f
}

// Clone TODO description.
func (f *Frustum) Clone() *Frustum {
	f.p.Call("clone")
	return f
}

// Copy TODO description.
func (f *Frustum) Copy(frustum *Frustum) *Frustum {
	f.p.Call("copy", frustum.p)
	return f
}

// SetFromMatrix TODO description.
func (f *Frustum) SetFromMatrix(m float64) *Frustum {
	f.p.Call("setFromMatrix", m)
	return f
}

// IntersectsObject TODO description.
func (f *Frustum) IntersectsObject() *Frustum {
	f.p.Call("intersectsObject")
	return f
}

// IntersectsSphere TODO description.
func (f *Frustum) IntersectsSphere(sphere float64) *Frustum {
	f.p.Call("intersectsSphere", sphere)
	return f
}

// IntersectsBox TODO description.
func (f *Frustum) IntersectsBox() *Frustum {
	f.p.Call("intersectsBox")
	return f
}

// ContainsPoint TODO description.
func (f *Frustum) ContainsPoint(point float64) *Frustum {
	f.p.Call("containsPoint", point)
	return f
}
