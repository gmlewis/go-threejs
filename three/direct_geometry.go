// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectGeometry represents a directgeometry.
type DirectGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DirectGeometry) JSObject() *js.Object { return d.p }

// DirectGeometry returns a DirectGeometry JavaScript class.
func (t *Three) DirectGeometry() *DirectGeometry {
	p := t.ctx.Get("DirectGeometry")
	return &DirectGeometry{p: p}
}

// NewDirectGeometry returns a new DirectGeometry object.
func (t *Three) NewDirectGeometry() *DirectGeometry {
	p := t.ctx.Get("DirectGeometry").New()
	return &DirectGeometry{p: p}
}

// ComputeFaceNormals TODO description.
func (d *DirectGeometry) ComputeFaceNormals() *DirectGeometry {
	d.p.Call("computeFaceNormals")
	return d
}

// ComputeVertexNormals TODO description.
func (d *DirectGeometry) ComputeVertexNormals() *DirectGeometry {
	d.p.Call("computeVertexNormals")
	return d
}

// ComputeGroups TODO description.
func (d *DirectGeometry) ComputeGroups(geometry float64) *DirectGeometry {
	d.p.Call("computeGroups", geometry)
	return d
}

// FromGeometry TODO description.
func (d *DirectGeometry) FromGeometry(geometry float64) *DirectGeometry {
	d.p.Call("fromGeometry", geometry)
	return d
}

// Dispose TODO description.
func (d *DirectGeometry) Dispose() *DirectGeometry {
	d.p.Call("dispose")
	return d
}
