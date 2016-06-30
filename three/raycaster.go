// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Raycaster represents a raycaster.
type Raycaster struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Raycaster) JSObject() *js.Object { return t.p }

// Raycaster returns a Raycaster object.
func (t *Three) Raycaster() *Raycaster {
	p := t.ctx.Get("Raycaster")
	return &Raycaster{p: p}
}

// New returns a new Raycaster object.
func (t *Raycaster) New(origin, direction, near, far float64) *Raycaster {
	p := t.p.New(origin, direction, near, far)
	return &Raycaster{p: p}
}

// Get TODO description.
func (r *Raycaster) Get() *Raycaster {
	r.p.Call("get")
	return r
}

// Set TODO description.
func (r *Raycaster) Set(origin, direction float64) *Raycaster {
	r.p.Call("set", origin, direction)
	return r
}

// SetFromCamera TODO description.
func (r *Raycaster) SetFromCamera(coords, camera float64) *Raycaster {
	r.p.Call("setFromCamera", coords, camera)
	return r
}

// IntersectObject TODO description.
func (r *Raycaster) IntersectObject(object, recursive float64) *Raycaster {
	r.p.Call("intersectObject", object, recursive)
	return r
}

// IntersectObjects TODO description.
func (r *Raycaster) IntersectObjects(objects, recursive float64) *Raycaster {
	r.p.Call("intersectObjects", objects, recursive)
	return r
}
