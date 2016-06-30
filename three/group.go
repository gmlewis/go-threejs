// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Group represents a group of objects.
type Group struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (g *Group) JSObject() *js.Object { return g.p }

// Group returns a Group object.
func (t *Three) Group() *Group {
	p := t.ctx.Get("Group")
	return &Group{p: p}
}

// New returns a new Group object.
func (g *Group) New() *Group {
	p := g.p.New()
	return &Group{p: p}
}
