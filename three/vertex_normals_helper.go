// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VertexNormalsHelper represents a vertexnormalshelper.
type VertexNormalsHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *VertexNormalsHelper) JSObject() *js.Object { return v.p }

// VertexNormalsHelper returns a VertexNormalsHelper object.
func (t *Three) VertexNormalsHelper() *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper")
	return &VertexNormalsHelper{p: p}
}

// New returns a new VertexNormalsHelper object.
func (v *VertexNormalsHelper) New(object, size, hex, linewidth float64) *VertexNormalsHelper {
	p := v.p.New(object, size, hex, linewidth)
	return &VertexNormalsHelper{p: p}
}
