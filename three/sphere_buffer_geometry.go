// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SphereBufferGeometry represents a spherebuffergeometry.
type SphereBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SphereBufferGeometry) JSObject() *js.Object { return s.p }

// SphereBufferGeometry returns a SphereBufferGeometry JavaScript class.
func (t *Three) SphereBufferGeometry() *SphereBufferGeometry {
	p := t.ctx.Get("SphereBufferGeometry")
	return &SphereBufferGeometry{p: p}
}

// NewSphereBufferGeometry returns a new SphereBufferGeometry object.
func (t *Three) NewSphereBufferGeometry(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength float64) *SphereBufferGeometry {
	p := t.ctx.Get("SphereBufferGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
	return &SphereBufferGeometry{p: p}
}
