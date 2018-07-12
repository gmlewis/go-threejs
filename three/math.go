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

// Math represents a math.
type Math struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *Math) JSObject() *js.Object { return m.p }

// Math returns a Math JavaScript class.
func (t *Three) Math() *Math {
	p := t.ctx.Get("Math")
	return MathFromJSObject(p)
}

// MathFromJSObject returns a wrapped Math JavaScript class.
func MathFromJSObject(p *js.Object) *Math {
	return &Math{p: p}
}

// NewMath returns a new Math object.
func (t *Three) NewMath() *Math {
	p := t.ctx.Get("Math").New()
	return MathFromJSObject(p)
}

// GenerateUUID TODO description.
func (m *Math) GenerateUUID() *Math {
	m.p.Call("generateUUID")
	return m
}

// Clamp TODO description.
func (m *Math) Clamp(value, min, max float64) *Math {
	m.p.Call("clamp", value, min, max)
	return m
}

// EuclideanModulo TODO description.
func (m *Math) EuclideanModulo(n, mod float64) *Math {
	m.p.Call("euclideanModulo", n, mod)
	return m
}

// MapLinear TODO description.
func (m *Math) MapLinear(x, a1, a2, b1, b2 float64) *Math {
	m.p.Call("mapLinear", x, a1, a2, b1, b2)
	return m
}

// Smoothstep TODO description.
func (m *Math) Smoothstep(x, min, max float64) *Math {
	m.p.Call("smoothstep", x, min, max)
	return m
}

// Smootherstep TODO description.
func (m *Math) Smootherstep(x, min, max float64) *Math {
	m.p.Call("smootherstep", x, min, max)
	return m
}

// Random16 TODO description.
func (m *Math) Random16() *Math {
	m.p.Call("random16")
	return m
}

// RandInt TODO description.
func (m *Math) RandInt(low, high float64) *Math {
	m.p.Call("randInt", low, high)
	return m
}

// RandFloat TODO description.
func (m *Math) RandFloat(low, high float64) *Math {
	m.p.Call("randFloat", low, high)
	return m
}

// RandFloatSpread TODO description.
func (m *Math) RandFloatSpread(rnge float64) *Math {
	m.p.Call("randFloatSpread", rnge)
	return m
}

// DegToRad TODO description.
func (m *Math) DegToRad() *Math {
	m.p.Call("degToRad")
	return m
}

// RadToDeg TODO description.
func (m *Math) RadToDeg() *Math {
	m.p.Call("radToDeg")
	return m
}

// IsPowerOfTwo TODO description.
func (m *Math) IsPowerOfTwo(value float64) *Math {
	m.p.Call("isPowerOfTwo", value)
	return m
}

// NearestPowerOfTwo TODO description.
func (m *Math) NearestPowerOfTwo(value float64) *Math {
	m.p.Call("nearestPowerOfTwo", value)
	return m
}

// NextPowerOfTwo TODO description.
func (m *Math) NextPowerOfTwo(value float64) *Math {
	m.p.Call("nextPowerOfTwo", value)
	return m
}
