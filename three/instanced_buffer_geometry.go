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

// InstancedBufferGeometry represents an instancedbuffergeometry.
type InstancedBufferGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InstancedBufferGeometry) JSObject() *js.Object { return i.p }

// InstancedBufferGeometry returns an InstancedBufferGeometry JavaScript class.
func (t *Three) InstancedBufferGeometry() *InstancedBufferGeometry {
	p := t.ctx.Get("InstancedBufferGeometry")
	return InstancedBufferGeometryFromJSObject(p)
}

// InstancedBufferGeometryFromJSObject returns a wrapped InstancedBufferGeometry JavaScript class.
func InstancedBufferGeometryFromJSObject(p *js.Object) *InstancedBufferGeometry {
	return &InstancedBufferGeometry{p: p}
}

// NewInstancedBufferGeometry returns a new InstancedBufferGeometry object.
func (t *Three) NewInstancedBufferGeometry() *InstancedBufferGeometry {
	p := t.ctx.Get("InstancedBufferGeometry").New()
	return InstancedBufferGeometryFromJSObject(p)
}

// AddGroup TODO description.
func (i *InstancedBufferGeometry) AddGroup(start, count, instances float64) *InstancedBufferGeometry {
	i.p.Call("addGroup", start, count, instances)
	return i
}

// Copy TODO description.
func (i *InstancedBufferGeometry) Copy(source *InstancedBufferGeometry) *InstancedBufferGeometry {
	i.p.Call("copy", source.p)
	return i
}
