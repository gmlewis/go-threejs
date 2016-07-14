// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LUT represents a lookup table.
type LUT struct{ p *js.Object }

// LUTFromJSObject returns a wrapped LUT JavaScript class.
func LUTFromJSObject(p *js.Object) *LUT {
	return &LUT{p: p}
}

// NewLUT returns a new LUT object.
func (t *Three) NewLUT(colormap string, numberofcolors int) *LUT {
	p := t.ctx.Get("Lut").New(colormap, numberofcolors)
	return LUTFromJSObject(p)
}

// Set TODO description.
func (l *LUT) Set(value float64) *LUT {
	l.p.Call("set", value)
	return l
}

// SetMin TODO description.
func (l *LUT) SetMin(min float64) *LUT {
	l.p.Call("setMin", min)
	return l
}

// SetMax TODO description.
func (l *LUT) SetMax(max float64) *LUT {
	l.p.Call("setMax", max)
	return l
}

// ChangeNumberOfColors TODO description.
func (l *LUT) ChangeNumberOfColors(numberofcolors int) *LUT {
	l.p.Call("changeNumberOfColors", numberofcolors)
	return l
}

// ChangeColorMap TODO description.
func (l *LUT) ChangeColorMap(colormap string) *LUT {
	l.p.Call("changeColorMap", colormap)
	return l
}

// Copy TODO description.
func (l *LUT) Copy(lut *LUT) *LUT {
	l.p.Call("copy", lut)
	return l
}

// GetColor TODO description.
func (l *LUT) GetColor(alpha float64) *Color {
	p := l.p.Call("getColor", alpha)
	if p == nil || p == js.Undefined {
		return nil
	}
	return ColorFromJSObject(p)
}

// AddColorMap TODO description.
func (l *LUT) AddColorMap(colormapName string, arrayOfColors *js.Object) *LUT {
	l.p.Call("addColorMap", colormapName, arrayOfColors)
	return l
}

// LUTLegendOpts defines options for LUT properties.
type LUTLegendOpts map[string]interface{}

// SetLegendOn TODO description.
func (l *LUT) SetLegendOn(parameters LUTLegendOpts) *Mesh {
	return MeshFromJSObject(l.p.Call("setLegendOn", parameters))
}

// SetLegendOff TODO description.
func (l *LUT) SetLegendOff() string {
	l.p.Call("setLegendOff")
	return ""
}

// SetLegendLayout TODO description.
func (l *LUT) SetLegendLayout(layout bool) bool {
	return l.p.Call("setLegendLayout", layout).Bool()
}

// SetLegendPosition TODO description.
func (l *LUT) SetLegendPosition(position *Vector3) *js.Object {
	return l.p.Call("setLegendPosition", position.JSObject())
}

// SetLegendLabels TODO description.
func (l *LUT) SetLegendLabels(parameters, callback interface{}) *js.Object {
	return l.p.Call("setLegendLabels", parameters, callback)
}
