// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Group represents a group of objects.
type Group struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (g *Group) JSObject() *js.Object { return g.p }

// Group returns a Group JavaScript class.
func (t *Three) Group() *Group {
	p := t.ctx.Get("Group")
	return GroupFromJSObject(p)
}

// GroupFromJSObject returns a wrapped Group JavaScript class.
func GroupFromJSObject(p *js.Object) *Group {
	return &Group{Object3DFromJSObject(p)}
}

// NewGroup returns a new Group object.
func (t *Three) NewGroup() *Group {
	p := t.ctx.Get("Group").New()
	return GroupFromJSObject(p)
}
