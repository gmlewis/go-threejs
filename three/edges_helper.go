// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesHelper represents an edgeshelper.
type EdgesHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (e *EdgesHelper) JSObject() *js.Object { return e.p }

// EdgesHelper returns an EdgesHelper object.
func (t *Three) EdgesHelper() *EdgesHelper {
	p := t.ctx.Get("EdgesHelper")
	return &EdgesHelper{p: p}
}

// New returns a new EdgesHelper object.
func (e *EdgesHelper) New(object, hex, thresholdAngle float64) *EdgesHelper {
	p := e.p.New(object, hex, thresholdAngle)
	return &EdgesHelper{p: p}
}
