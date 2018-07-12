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

// LightShadow represents a light shadow.
//
// http://threejs.org/docs/index.html#Reference/Lights/LightShadow
type LightShadow struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LightShadow) JSObject() *js.Object { return l.p }

// LightShadow returns a LightShadow JavaScript class.
func (t *Three) LightShadow() *LightShadow {
	p := t.ctx.Get("LightShadow")
	return LightShadowFromJSObject(p)
}

// LightShadowFromJSObject returns a wrapped LightShadow JavaScript class.
func LightShadowFromJSObject(p *js.Object) *LightShadow {
	return &LightShadow{p: p}
}

// NewLightShadow returns a new LightShadow object.
func (t *Three) NewLightShadow(camera *Camera) *LightShadow {
	p := t.ctx.Get("LightShadow").New(camera.JSObject())
	return LightShadowFromJSObject(p)
}

// Copy TODO description.
func (l *LightShadow) Copy(source *LightShadow) *LightShadow {
	l.p.Call("copy", source.p)
	return l
}

// Clone TODO description.
func (l *LightShadow) Clone() *LightShadow {
	l.p.Call("clone")
	return l
}

// Camera returns the property of the same name.
func (l *LightShadow) Camera() *Camera {
	return &Camera{&Object3D{p: l.p.Get("camera")}}
}

// MapSize returns the property of the same name.
func (l *LightShadow) MapSize() *Vector2 {
	return &Vector2{p: l.p.Get("mapSize")}
}

// SetMapSize sets the mapSize property.
func (l *LightShadow) SetMapSize(value *Vector2) *LightShadow {
	l.p.Set("mapSize", value.p)
	return l
}

// Matrix returns the property of the same name.
func (l *LightShadow) Matrix() *Matrix4 {
	return &Matrix4{p: l.p.Get("matrix")}
}

// SetMatrix sets the matrix property.
func (l *LightShadow) SetMatrix(value *Matrix4) *LightShadow {
	l.p.Set("matrix", value.p)
	return l
}

// Bias returns the property of the same name.
func (l *LightShadow) Bias() float64 {
	return l.p.Get("bias").Float()
}

// SetBias sets the bias property.
func (l *LightShadow) SetBias(value float64) *LightShadow {
	l.p.Set("bias", value)
	return l
}

// Radius returns the property of the same name.
func (l *LightShadow) Radius() float64 {
	return l.p.Get("radius").Float()
}

// SetRadius sets the radius property.
func (l *LightShadow) SetRadius(value float64) *LightShadow {
	l.p.Set("radius", value)
	return l
}
