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
func (l *LOD) JSObject() *js.Object { return l.p }

// LOD returns a LOD JavaScript class.
func (t *Three) LOD() *LOD {
	p := t.ctx.Get("LOD")
	return LODFromJSObject(p)
}

// LODFromJSObject returns a wrapped LOD JavaScript class.
func LODFromJSObject(p *js.Object) *LOD {
	return &LOD{p: p}
}

// NewLOD returns a new LOD object.
func (t *Three) NewLOD() *LOD {
	p := t.ctx.Get("LOD").New()
	return LODFromJSObject(p)
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
func (l *LOD) Copy(source *LOD) *LOD {
	l.p.Call("copy", source.p)
	return l
}

// ToJSON TODO description.
func (l *LOD) ToJSON(meta float64) *LOD {
	l.p.Call("toJSON", meta)
	return l
}
