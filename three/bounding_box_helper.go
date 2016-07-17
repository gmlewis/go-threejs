// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoundingBoxHelper is a helper object to show the
// world-axis-aligned bounding box for an object.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/BoundingBoxHelper
type BoundingBoxHelper struct{ *Mesh }

// JSObject returns the underlying *js.Object.
func (b *BoundingBoxHelper) JSObject() *js.Object { return b.p }

// BoundingBoxHelper returns a BoundingBoxHelper JavaScript class.
func (t *Three) BoundingBoxHelper() *BoundingBoxHelper {
	p := t.ctx.Get("BoundingBoxHelper")
	return BoundingBoxHelperFromJSObject(p)
}

// BoundingBoxHelperFromJSObject returns a wrapped BoundingBoxHelper JavaScript class.
func BoundingBoxHelperFromJSObject(p *js.Object) *BoundingBoxHelper {
	return &BoundingBoxHelper{MeshFromJSObject(p)}
}

// NewBoundingBoxHelper returns a new BoundingBoxHelper object.
func (t *Three) NewBoundingBoxHelper(object JSObject, hex int) *BoundingBoxHelper {
	p := t.ctx.Get("BoundingBoxHelper").New(object.JSObject(), hex)
	return BoundingBoxHelperFromJSObject(p)
}

// Update updates the BoundingBoxHelper based on the object property.
func (b *BoundingBoxHelper) Update() *BoundingBoxHelper {
	b.p.Call("update")
	return b
}
