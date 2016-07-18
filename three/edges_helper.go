// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesHelper represents a wireframe object that shows
// the "hard" edges of another object's geometry.
// To draw a full wireframe image of an object, see WireframeHelper.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/EdgesHelper
type EdgesHelper struct{ *Line }

// JSObject returns the underlying *js.Object.
func (e *EdgesHelper) JSObject() *js.Object { return e.p }

// EdgesHelper returns an EdgesHelper JavaScript class.
func (t *Three) EdgesHelper() *EdgesHelper {
	p := t.ctx.Get("EdgesHelper")
	return EdgesHelperFromJSObject(p)
}

// EdgesHelperFromJSObject returns a wrapped EdgesHelper JavaScript class.
func EdgesHelperFromJSObject(p *js.Object) *EdgesHelper {
	return &EdgesHelper{LineFromJSObject(p)}
}

// NewEdgesHelper creates a Line, showing only the "hard" edges of
// the passed object; specifically, no edge will be drawn between
// faces which are adjacent and coplanar (or nearly coplanar).
//
//     object -- Object of which to draw edges
//     color -- Color of the edges.
//     thresholdAngle -- the minimim angle (in degrees), between the
//         face normals of adjacent faces, that is required to render
//         an edge. Default is 0.1.
func (t *Three) NewEdgesHelper(object JSObject, hex int, thresholdAngle float64) *EdgesHelper {
	p := t.ctx.Get("EdgesHelper").New(object.JSObject(), hex, thresholdAngle)
	return EdgesHelperFromJSObject(p)
}
