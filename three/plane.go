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

// Plane represents a plane.
type Plane struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *Plane) JSObject() *js.Object { return p.p }

// Plane returns a Plane JavaScript class.
func (t *Three) Plane() *Plane {
	p := t.ctx.Get("Plane")
	return PlaneFromJSObject(p)
}

// PlaneFromJSObject returns a wrapped Plane JavaScript class.
func PlaneFromJSObject(p *js.Object) *Plane {
	return &Plane{p: p}
}

// NewPlane returns a new Plane object.
func (t *Three) NewPlane(normal, constant float64) *Plane {
	p := t.ctx.Get("Plane").New(normal, constant)
	return PlaneFromJSObject(p)
}

// Set TODO description.
func (p *Plane) Set(normal, constant float64) *Plane {
	p.p.Call("set", normal, constant)
	return p
}

// SetComponents TODO description.
func (p *Plane) SetComponents(x, y, z, w float64) *Plane {
	p.p.Call("setComponents", x, y, z, w)
	return p
}

// SetFromNormalAndCoplanarPoint TODO description.
func (p *Plane) SetFromNormalAndCoplanarPoint(normal, point float64) *Plane {
	p.p.Call("setFromNormalAndCoplanarPoint", normal, point)
	return p
}

// SetFromCoplanarPoints TODO description.
func (p *Plane) SetFromCoplanarPoints() *Plane {
	p.p.Call("setFromCoplanarPoints")
	return p
}

// Clone TODO description.
func (p *Plane) Clone() *Plane {
	p.p.Call("clone")
	return p
}

// Copy TODO description.
func (p *Plane) Copy(plane *Plane) *Plane {
	p.p.Call("copy", plane.p)
	return p
}

// Normalize TODO description.
func (p *Plane) Normalize() *Plane {
	p.p.Call("normalize")
	return p
}

// Negate TODO description.
func (p *Plane) Negate() *Plane {
	p.p.Call("negate")
	return p
}

// DistanceToPoint TODO description.
func (p *Plane) DistanceToPoint(point float64) *Plane {
	p.p.Call("distanceToPoint", point)
	return p
}

// DistanceToSphere TODO description.
func (p *Plane) DistanceToSphere(sphere float64) *Plane {
	p.p.Call("distanceToSphere", sphere)
	return p
}

// ProjectPoint TODO description.
func (p *Plane) ProjectPoint(point, optionalTarget float64) *Plane {
	p.p.Call("projectPoint", point, optionalTarget)
	return p
}

// OrthoPoint TODO description.
func (p *Plane) OrthoPoint(point, optionalTarget float64) *Plane {
	p.p.Call("orthoPoint", point, optionalTarget)
	return p
}

// IntersectLine TODO description.
func (p *Plane) IntersectLine() *Plane {
	p.p.Call("intersectLine")
	return p
}

// IntersectsLine TODO description.
func (p *Plane) IntersectsLine(line float64) *Plane {
	p.p.Call("intersectsLine", line)
	return p
}

// IntersectsBox TODO description.
func (p *Plane) IntersectsBox(box float64) *Plane {
	p.p.Call("intersectsBox", box)
	return p
}

// IntersectsSphere TODO description.
func (p *Plane) IntersectsSphere(sphere float64) *Plane {
	p.p.Call("intersectsSphere", sphere)
	return p
}

// CoplanarPoint TODO description.
func (p *Plane) CoplanarPoint(optionalTarget float64) *Plane {
	p.p.Call("coplanarPoint", optionalTarget)
	return p
}

// ApplyMatrix4 TODO description.
func (p *Plane) ApplyMatrix4() *Plane {
	p.p.Call("applyMatrix4")
	return p
}

// Translate TODO description.
func (p *Plane) Translate(offset float64) *Plane {
	p.p.Call("translate", offset)
	return p
}

// Equals TODO description.
func (p *Plane) Equals(plane float64) *Plane {
	p.p.Call("equals", plane)
	return p
}
