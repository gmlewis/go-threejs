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

// Fog represents linear fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/Fog
type Fog struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Fog) JSObject() *js.Object { return f.p }

// Fog returns a Fog JavaScript class.
func (t *Three) Fog() *Fog {
	p := t.ctx.Get("Fog")
	return FogFromJSObject(p)
}

// FogFromJSObject returns a wrapped Fog JavaScript class.
func FogFromJSObject(p *js.Object) *Fog {
	return &Fog{p: p}
}

// NewFog returns a new Fog object.
func (t *Three) NewFog(color, near, far float64) *Fog {
	p := t.ctx.Get("Fog").New(color, near, far)
	return FogFromJSObject(p)
}
