// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
	return &PointsMaterial{p: p}
}

// Copy TODO description.
func (p *PointsMaterial) Copy(source *PointsMaterial) *PointsMaterial {
	p.p.Call("copy", source.p)
	return p
}
