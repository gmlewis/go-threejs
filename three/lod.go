// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LOD shows meshes with more or less geometry based on distance.
//
// http://threejs.org/docs/index.html#Reference/Objects/LOD
type LOD struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LOD) JSObject() *js.Object { return t.p }

// LOD returns a LOD object.
func (t *Three) LOD() *LOD {
	p := t.ctx.Get("LOD")
	return &LOD{p: p}
}

// New returns a new LOD object.
func (t *LOD) New() *LOD {
	p := t.p.New()
	return &LOD{p: p}
}

// Get TODO description.
func (l *LOD) Get() *LOD {
	l.p.Call("get")
	return l
}

// AddLevel TODO description.
func (l *LOD) AddLevel(object, distance float64) *LOD {
	l.p.Call("addLevel", object, distance)
	return l
}

// GetObjectForDistance TODO description.
func (l *LOD) GetObjectForDistance(distance float64) *LOD {
	l.p.Call("getObjectForDistance", distance)
	return l
}

// Copy TODO description.
func (l *LOD) Copy(source float64) *LOD {
	l.p.Call("copy", source)
	return l
}

// ToJSON TODO description.
func (l *LOD) ToJSON(meta float64) *LOD {
	l.p.Call("toJSON", meta)
	return l
}
