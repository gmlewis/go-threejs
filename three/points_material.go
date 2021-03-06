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

// PointsMaterial represents a material used by particle systems.
//
// http://threejs.org/docs/index.html#Reference/Materials/PointsMaterial
type PointsMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PointsMaterial) JSObject() *js.Object { return p.p }

// PointsMaterial returns a PointsMaterial JavaScript class.
func (t *Three) PointsMaterial() *PointsMaterial {
	p := t.ctx.Get("PointsMaterial")
	return PointsMaterialFromJSObject(p)
}

// PointsMaterialFromJSObject returns a wrapped PointsMaterial JavaScript class.
func PointsMaterialFromJSObject(p *js.Object) *PointsMaterial {
	return &PointsMaterial{p: p}
}

// NewPointsMaterial returns a new PointsMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — Particle color in hexadecimal. Default is 0xffffff.
//     map — a texture.If defined, then a point has the data from texture as colors. Default is null.
//     size — Define size of particles. Default is 1.0.
//     sizeAttenuation — Enable/disable size attenuation with distance.
//     vertexColors — Define whether the material uses vertex colors, or not. Default is false.
//     fog — Define whether the material color is affected by global fog settings. Default is true.
func (t *Three) NewPointsMaterial(parameters map[string]interface{}) *PointsMaterial {
	p := t.ctx.Get("PointsMaterial").New(parameters)
	return PointsMaterialFromJSObject(p)
}

// Copy TODO description.
func (p *PointsMaterial) Copy(source *PointsMaterial) *PointsMaterial {
	p.p.Call("copy", source.p)
	return p
}
