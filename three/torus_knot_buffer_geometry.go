// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotBufferGeometry represents a torusknotbuffergeometry.
type TorusKnotBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TorusKnotBufferGeometry) JSObject() *js.Object { return t.p }

// TorusKnotBufferGeometry returns a TorusKnotBufferGeometry JavaScript class.
func (t *Three) TorusKnotBufferGeometry() *TorusKnotBufferGeometry {
	p := t.ctx.Get("TorusKnotBufferGeometry")
	return &TorusKnotBufferGeometry{p: p}
}

// NewTorusKnotBufferGeometry returns a new TorusKnotBufferGeometry object.
func (t *Three) NewTorusKnotBufferGeometry(radius, tube, tubularSegments, radialSegments, p, q float64) *TorusKnotBufferGeometry {
	s := t.ctx.Get("TorusKnotBufferGeometry").New(radius, tube, tubularSegments, radialSegments, p, q)
	return &TorusKnotBufferGeometry{p: s}
}
