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

// FogExp2 represents exponential fog.
//
// http://threejs.org/docs/index.html#Reference/Scenes/FogExp2
type FogExp2 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *FogExp2) JSObject() *js.Object { return f.p }

// FogExp2 returns a FogExp2 JavaScript class.
func (t *Three) FogExp2() *FogExp2 {
	p := t.ctx.Get("FogExp2")
	return FogExp2FromJSObject(p)
}

// FogExp2FromJSObject returns a wrapped FogExp2 JavaScript class.
func FogExp2FromJSObject(p *js.Object) *FogExp2 {
	return &FogExp2{p: p}
}

// NewFogExp2 returns a new FogExp2 object.
func (t *Three) NewFogExp2(color, density float64) *FogExp2 {
	p := t.ctx.Get("FogExp2").New(color, density)
	return FogExp2FromJSObject(p)
}
