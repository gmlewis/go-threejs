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

// DirectGeometry represents a directgeometry.
type DirectGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DirectGeometry) JSObject() *js.Object { return d.p }

// DirectGeometry returns a DirectGeometry JavaScript class.
func (t *Three) DirectGeometry() *DirectGeometry {
	p := t.ctx.Get("DirectGeometry")
	return DirectGeometryFromJSObject(p)
}

// DirectGeometryFromJSObject returns a wrapped DirectGeometry JavaScript class.
func DirectGeometryFromJSObject(p *js.Object) *DirectGeometry {
	return &DirectGeometry{p: p}
}

// NewDirectGeometry returns a new DirectGeometry object.
func (t *Three) NewDirectGeometry() *DirectGeometry {
	p := t.ctx.Get("DirectGeometry").New()
	return DirectGeometryFromJSObject(p)
}

// ComputeFaceNormals TODO description.
func (d *DirectGeometry) ComputeFaceNormals() *DirectGeometry {
	d.p.Call("computeFaceNormals")
	return d
}

// ComputeVertexNormals TODO description.
func (d *DirectGeometry) ComputeVertexNormals() *DirectGeometry {
	d.p.Call("computeVertexNormals")
	return d
}

// ComputeGroups TODO description.
func (d *DirectGeometry) ComputeGroups(geometry float64) *DirectGeometry {
	d.p.Call("computeGroups", geometry)
	return d
}

// FromGeometry TODO description.
func (d *DirectGeometry) FromGeometry(geometry float64) *DirectGeometry {
	d.p.Call("fromGeometry", geometry)
	return d
}

// Dispose TODO description.
func (d *DirectGeometry) Dispose() *DirectGeometry {
	d.p.Call("dispose")
	return d
}
