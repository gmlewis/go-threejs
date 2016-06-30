// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// GridHelper represents a gridhelper.
type GridHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *GridHelper) JSObject() *js.Object { return t.p }

// GridHelper returns a GridHelper object.
func (t *Three) GridHelper() *GridHelper {
	p := t.ctx.Get("GridHelper")
	return &GridHelper{p: p}
}

// New returns a new GridHelper object.
func (t *GridHelper) New(size, step float64) *GridHelper {
	p := t.p.New(size, step)
	return &GridHelper{p: p}
}

// SetColors TODO description.
func (g *GridHelper) SetColors(colorCenterLine, colorGrid float64) *GridHelper {
	g.p.Call("setColors", colorCenterLine, colorGrid)
	return g
}
