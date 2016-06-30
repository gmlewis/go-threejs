// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LensFlare represents a simulated lens flare that tracks a light.
//
// http://threejs.org/docs/index.html#Reference/Objects/LensFlare
type LensFlare struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LensFlare) JSObject() *js.Object { return t.p }

// LensFlare returns a LensFlare object.
func (t *Three) LensFlare() *LensFlare {
	p := t.ctx.Get("LensFlare")
	return &LensFlare{p: p}
}

// New returns a new LensFlare object.
//
//     texture -- THREE.Texture (optional)
//     size -- size in pixels (-1 = use texture.width)
//     distance -- (0-1) from light source (0 = at light source)
//     blending -- Blending Mode - Defaults to THREE.NormalBlending
//     color -- The color of the lens flare
func (t *LensFlare) New(texture *js.Object, size, distance float64, blending, color int) *LensFlare {
	p := t.p.New(texture, size, distance, blending, color)
	return &LensFlare{p: p}
}

// Add TODO description.
func (l *LensFlare) Add(texture *js.Object, size, distance float64, blending, color, opacity int) *LensFlare {
	l.p.Call("add", texture, size, distance, blending, color, opacity)
	return l
}

// Copy TODO description.
func (l *LensFlare) Copy(source *LensFlare) *LensFlare {
	l.p.Call("copy", source.p)
	return l
}
