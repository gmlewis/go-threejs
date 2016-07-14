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

// VertexNormalsHelper returns a VertexNormalsHelper JavaScript class.
func (t *Three) VertexNormalsHelper() *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper")
	return VertexNormalsHelperFromJSObject(p)
}

// VertexNormalsHelperFromJSObject returns a wrapped VertexNormalsHelper JavaScript class.
func VertexNormalsHelperFromJSObject(p *js.Object) *VertexNormalsHelper {
	return &VertexNormalsHelper{p: p}
}

// NewVertexNormalsHelper returns a new VertexNormalsHelper object.
func (t *Three) NewVertexNormalsHelper(object, size, hex, linewidth float64) *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper").New(object, size, hex, linewidth)
	return VertexNormalsHelperFromJSObject(p)
}
