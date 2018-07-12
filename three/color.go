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

// Color represents a color.
//
// http://threejs.org/docs/index.html#Reference/Math/Color
type Color struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *Color) JSObject() *js.Object { return c.p }

// Color returns a Color JavaScript class.
func (t *Three) Color() *Color {
	p := t.ctx.Get("Color")
	return ColorFromJSObject(p)
}

// ColorFromJSObject returns a wrapped Color JavaScript object.
func ColorFromJSObject(p *js.Object) *Color {
	return &Color{p: p}
}

// NewColor returns a new Color object.
func (t *Three) NewColor(color ...interface{}) *Color {
	p := t.ctx.Get("Color").New(color...)
	return ColorFromJSObject(p)
}

// Set TODO description.
func (c *Color) Set(value float64) *Color {
	c.p.Call("set", value)
	return c
}

// SetScalar TODO description.
func (c *Color) SetScalar(scalar float64) *Color {
	c.p.Call("setScalar", scalar)
	return c
}

// SetHex TODO description.
func (c *Color) SetHex(hex int) *Color {
	c.p.Call("setHex", hex)
	return c
}

// SetRGB TODO description.
func (c *Color) SetRGB(red, green, blue float64) *Color {
	c.p.Call("setRGB", red, green, blue)
	return c
}

// SetHSL TODO description.
func (c *Color) SetHSL(hue, saturation, luminance float64) *Color {
	c.p.Call("setHSL", hue, saturation, luminance)
	return c
}

// SetStyle TODO description.
func (c *Color) SetStyle(style float64) *Color {
	c.p.Call("setStyle", style)
	return c
}

// Clone TODO description.
func (c *Color) Clone() *Color {
	c.p.Call("clone")
	return c
}

// Copy TODO description.
func (c *Color) Copy(color *Color) *Color {
	c.p.Call("copy", color.p)
	return c
}

// CopyGammaToLinear TODO description.
func (c *Color) CopyGammaToLinear(color, gammaFactor float64) *Color {
	c.p.Call("copyGammaToLinear", color, gammaFactor)
	return c
}

// CopyLinearToGamma TODO description.
func (c *Color) CopyLinearToGamma(color, gammaFactor float64) *Color {
	c.p.Call("copyLinearToGamma", color, gammaFactor)
	return c
}

// ConvertGammaToLinear TODO description.
func (c *Color) ConvertGammaToLinear() *Color {
	c.p.Call("convertGammaToLinear")
	return c
}

// ConvertLinearToGamma TODO description.
func (c *Color) ConvertLinearToGamma() *Color {
	c.p.Call("convertLinearToGamma")
	return c
}

// GetHex TODO description.
func (c *Color) GetHex() *Color {
	c.p.Call("getHex")
	return c
}

// GetHexString TODO description.
func (c *Color) GetHexString() *Color {
	c.p.Call("getHexString")
	return c
}

// GetHSL TODO description.
func (c *Color) GetHSL(optionalTarget float64) *Color {
	c.p.Call("getHSL", optionalTarget)
	return c
}

// GetStyle TODO description.
func (c *Color) GetStyle() *Color {
	c.p.Call("getStyle")
	return c
}

// OffsetHSL TODO description.
func (c *Color) OffsetHSL(h, s, l float64) *Color {
	c.p.Call("offsetHSL", h, s, l)
	return c
}

// Add TODO description.
func (c *Color) Add(color float64) *Color {
	c.p.Call("add", color)
	return c
}

// AddColors TODO description.
func (c *Color) AddColors(color1, color2 float64) *Color {
	c.p.Call("addColors", color1, color2)
	return c
}

// AddScalar TODO description.
func (c *Color) AddScalar(s float64) *Color {
	c.p.Call("addScalar", s)
	return c
}

// Multiply TODO description.
func (c *Color) Multiply(color float64) *Color {
	c.p.Call("multiply", color)
	return c
}

// MultiplyScalar TODO description.
func (c *Color) MultiplyScalar(s float64) *Color {
	c.p.Call("multiplyScalar", s)
	return c
}

// Lerp TODO description.
func (c *Color) Lerp(color, alpha float64) *Color {
	c.p.Call("lerp", color, alpha)
	return c
}

// Equals TODO description.
func (c *Color) Equals(color *Color) bool {
	return c.p.Call("equals", color).Bool()
}

// FromArray TODO description.
func (c *Color) FromArray(array, offset float64) *Color {
	c.p.Call("fromArray", array, offset)
	return c
}

// ToArray TODO description.
func (c *Color) ToArray(array, offset float64) *Color {
	c.p.Call("toArray", array, offset)
	return c
}

// R returns the r-component of the Color.
func (c *Color) R() float64 {
	return c.p.Get("r").Float()
}

// G returns the g-component of the Color.
func (c *Color) G() float64 {
	return c.p.Get("g").Float()
}

// B returns the b-component of the Color.
func (c *Color) B() float64 {
	return c.p.Get("b").Float()
}
