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

// Frustum represents a frustum.
type Frustum struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Frustum) JSObject() *js.Object { return f.p }

// Frustum returns a Frustum JavaScript class.
func (t *Three) Frustum() *Frustum {
	p := t.ctx.Get("Frustum")
	return FrustumFromJSObject(p)
}

// FrustumFromJSObject returns a wrapped Frustum JavaScript class.
func FrustumFromJSObject(p *js.Object) *Frustum {
	return &Frustum{p: p}
}

// NewFrustum returns a new Frustum object.
func (t *Three) NewFrustum(p0, p1, p2, p3, p4, p5 float64) *Frustum {
	p := t.ctx.Get("Frustum").New(p0, p1, p2, p3, p4, p5)
	return FrustumFromJSObject(p)
}

// Set TODO description.
func (f *Frustum) Set(p0, p1, p2, p3, p4, p5 float64) *Frustum {
	f.p.Call("set", p0, p1, p2, p3, p4, p5)
	return f
}

// Clone TODO description.
func (f *Frustum) Clone() *Frustum {
	f.p.Call("clone")
	return f
}

// Copy TODO description.
func (f *Frustum) Copy(frustum *Frustum) *Frustum {
	f.p.Call("copy", frustum.p)
	return f
}

// SetFromMatrix TODO description.
func (f *Frustum) SetFromMatrix(m float64) *Frustum {
	f.p.Call("setFromMatrix", m)
	return f
}

// IntersectsObject TODO description.
func (f *Frustum) IntersectsObject() *Frustum {
	f.p.Call("intersectsObject")
	return f
}

// IntersectsSphere TODO description.
func (f *Frustum) IntersectsSphere(sphere float64) *Frustum {
	f.p.Call("intersectsSphere", sphere)
	return f
}

// IntersectsBox TODO description.
func (f *Frustum) IntersectsBox() *Frustum {
	f.p.Call("intersectsBox")
	return f
}

// ContainsPoint TODO description.
func (f *Frustum) ContainsPoint(point float64) *Frustum {
	f.p.Call("containsPoint", point)
	return f
}
