// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VertexNormalsHelper renders arrows to visualize an object's vertex normal vectors. Requires that
// normals have been specified in a custom attribute or have been calculated using computeVertexNormals.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/VertexNormalsHelper
type VertexNormalsHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (v *VertexNormalsHelper) JSObject() *js.Object { return v.p }

// VertexNormalsHelper returns a VertexNormalsHelper JavaScript class.
func (t *Three) VertexNormalsHelper() *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper")
	return &VertexNormalsHelper{p: p}
}

// NewVertexNormalsHelper returns a new VertexNormalsHelper object.
//
//     object -- object for which to render vertex normals size -- size (length) of the arrows color -- color of
//     the arrows linewidth -- width of the arrow lines
func (t *Three) NewVertexNormalsHelper(object, size, hex, linewidth float64) *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper").New(object, size, hex, linewidth)
	return &VertexNormalsHelper{p: p}
}

/* TODO:
Object is the attached object.
Update updates the vertex normal preview based on movement of the object.
*/
