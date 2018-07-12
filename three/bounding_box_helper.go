// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
