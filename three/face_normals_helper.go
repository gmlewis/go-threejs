// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FaceNormalsHelper renders arrows to visualize an object's
// face normals. Requires that the object's geometry be an
// instance of Geometry (does not work with BufferGeometry),
// and that face normals have been specified on all faces or
// calculated with ComputeFaceNormals.
type FaceNormalsHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *FaceNormalsHelper) JSObject() *js.Object { return f.p }

// FaceNormalsHelper returns a FaceNormalsHelper JavaScript class.
func (t *Three) FaceNormalsHelper() *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper")
	return FaceNormalsHelperFromJSObject(p)
}

// FaceNormalsHelperFromJSObject returns a wrapped FaceNormalsHelper JavaScript class.
func FaceNormalsHelperFromJSObject(p *js.Object) *FaceNormalsHelper {
	return &FaceNormalsHelper{p: p}
}

// NewFaceNormalsHelper returns a new FaceNormalsHelper object.
//
//     object -- object for which to render face normals
//     size -- size (length) of the arrows
//     color -- color of the arrows
//     linewidth -- width of the arrow lines
func (t *Three) NewFaceNormalsHelper(object JSObject, size float64, hex int, linewidth float64) *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper").New(object.JSObject(), size, hex, linewidth)
	return FaceNormalsHelperFromJSObject(p)
}
