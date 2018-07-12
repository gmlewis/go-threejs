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

// LensFlare represents a simulated lens flare that tracks a light.
//
// http://threejs.org/docs/index.html#Reference/Objects/LensFlare
type LensFlare struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LensFlare) JSObject() *js.Object { return l.p }

// LensFlare returns a LensFlare JavaScript class.
func (t *Three) LensFlare() *LensFlare {
	p := t.ctx.Get("LensFlare")
	return LensFlareFromJSObject(p)
}

// LensFlareFromJSObject returns a wrapped LensFlare JavaScript class.
func LensFlareFromJSObject(p *js.Object) *LensFlare {
	return &LensFlare{p: p}
}

// NewLensFlare returns a new LensFlare object.
//
//     texture -- THREE.Texture (optional)
//     size -- size in pixels (-1 = use texture.width)
//     distance -- (0-1) from light source (0 = at light source)
//     blending -- Blending Mode - Defaults to THREE.NormalBlending
//     color -- The color of the lens flare
func (t *Three) NewLensFlare(texture *js.Object, size, distance float64, blending, color int) *LensFlare {
	p := t.ctx.Get("LensFlare").New(texture, size, distance, blending, color)
	return LensFlareFromJSObject(p)
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
