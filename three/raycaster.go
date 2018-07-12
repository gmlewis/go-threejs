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

// Raycaster is the JavaScript class that makes raycasting easier.
// Raycasting is used for picking and more.
//
// http://threejs.org/docs/index.html#Reference/Core/Raycaster
type Raycaster struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (r *Raycaster) JSObject() *js.Object { return r.p }

// Raycaster returns a Raycaster JavaScript class.
func (t *Three) Raycaster() *Raycaster {
	p := t.ctx.Get("Raycaster")
	return RaycasterFromJSObject(p)
}

// RaycasterFromJSObject returns a wrapped Raycaster JavaScript class.
func RaycasterFromJSObject(p *js.Object) *Raycaster {
	return &Raycaster{p: p}
}

// NewRaycaster returns a new Raycaster object.
func (t *Three) NewRaycaster(origin, direction, near, far float64) *Raycaster {
	p := t.ctx.Get("Raycaster").New(origin, direction, near, far)
	return RaycasterFromJSObject(p)
}

// Get TODO description.
func (r *Raycaster) Get() *Raycaster {
	r.p.Call("get")
	return r
}

// Set TODO description.
func (r *Raycaster) Set(origin, direction float64) *Raycaster {
	r.p.Call("set", origin, direction)
	return r
}

// SetFromCamera TODO description.
func (r *Raycaster) SetFromCamera(coords, camera float64) *Raycaster {
	r.p.Call("setFromCamera", coords, camera)
	return r
}

// IntersectObject TODO description.
func (r *Raycaster) IntersectObject(object, recursive float64) *Raycaster {
	r.p.Call("intersectObject", object, recursive)
	return r
}

// IntersectObjects TODO description.
func (r *Raycaster) IntersectObjects(objects, recursive float64) *Raycaster {
	r.p.Call("intersectObjects", objects, recursive)
	return r
}
