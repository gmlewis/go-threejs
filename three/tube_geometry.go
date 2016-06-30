// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TubeGeometry represents a tubegeometry.
type TubeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *TubeGeometry) JSObject() *js.Object { return t.p }

// TubeGeometry returns a TubeGeometry object.
func (t *Three) TubeGeometry() *TubeGeometry {
	p := t.ctx.Get("TubeGeometry")
	return &TubeGeometry{p: p}
}

// New returns a new TubeGeometry object.
func (t *TubeGeometry) New(path, segments, radius, radialSegments, closed, taper float64) *TubeGeometry {
	p := t.p.New(path, segments, radius, radialSegments, closed, taper)
	return &TubeGeometry{p: p}
}
