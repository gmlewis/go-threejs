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
	return VertexNormalsHelperFromJSObject(p)
}

// VertexNormalsHelperFromJSObject returns a wrapped VertexNormalsHelper JavaScript class.
func VertexNormalsHelperFromJSObject(p *js.Object) *VertexNormalsHelper {
	return &VertexNormalsHelper{p: p}
}

// NewVertexNormalsHelper returns a new VertexNormalsHelper object.
//
//     object -- object for which to render vertex normals size -- size (length) of the arrows color -- color of
//     the arrows linewidth -- width of the arrow lines
func (t *Three) NewVertexNormalsHelper(object, size, hex, linewidth float64) *VertexNormalsHelper {
	p := t.ctx.Get("VertexNormalsHelper").New(object, size, hex, linewidth)
	return VertexNormalsHelperFromJSObject(p)
}

/* TODO:
Object is the attached object.
Update updates the vertex normal preview based on movement of the object.
*/
