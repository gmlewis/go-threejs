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

// Light is the JavaScript abstract base class for lights.
//
// http://threejs.org/docs/index.html#Reference/Lights/Light
type Light struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (l *Light) JSObject() *js.Object { return l.p }

// Light returns a Light JavaScript class.
func (t *Three) Light() *Light {
	return LightFromJSObject(t.ctx.Get("Light"))
}

// LightFromJSObject returns a wrapped Light JavaScript class.
func LightFromJSObject(p *js.Object) *Light {
	return &Light{&Object3D{p: p}}
}

// NewLight returns a new Light object.
func (t *Three) NewLight(color, intensity float64) *Light {
	return LightFromJSObject(t.ctx.Get("Light").New(color, intensity))
}

// Copy TODO description.
func (l *Light) Copy(source *Light) *Light {
	l.p.Call("copy", source.p)
	return l
}

// ToJSON TODO description.
func (l *Light) ToJSON(meta float64) *Light {
	l.p.Call("toJSON", meta)
	return l
}
