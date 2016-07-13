// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
