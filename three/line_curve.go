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

// LineCurve represents a linecurve.
type LineCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LineCurve) JSObject() *js.Object { return l.p }

// LineCurve returns a LineCurve JavaScript class.
func (t *Three) LineCurve() *LineCurve {
	p := t.ctx.Get("LineCurve")
	return LineCurveFromJSObject(p)
}

// LineCurveFromJSObject returns a wrapped LineCurve JavaScript class.
func LineCurveFromJSObject(p *js.Object) *LineCurve {
	return &LineCurve{p: p}
}

// NewLineCurve returns a new LineCurve object.
func (t *Three) NewLineCurve(v1, v2 float64) *LineCurve {
	p := t.ctx.Get("LineCurve").New(v1, v2)
	return LineCurveFromJSObject(p)
}

// GetPoint TODO description.
func (l *LineCurve) GetPoint(t float64) *LineCurve {
	l.p.Call("getPoint", t)
	return l
}

// GetPointAt TODO description.
func (l *LineCurve) GetPointAt(u float64) *LineCurve {
	l.p.Call("getPointAt", u)
	return l
}

// GetTangent TODO description.
func (l *LineCurve) GetTangent(t float64) *LineCurve {
	l.p.Call("getTangent", t)
	return l
}
