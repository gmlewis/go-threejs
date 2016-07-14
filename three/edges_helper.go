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

// EdgesHelper returns an EdgesHelper JavaScript class.
func (t *Three) EdgesHelper() *EdgesHelper {
	p := t.ctx.Get("EdgesHelper")
	return EdgesHelperFromJSObject(p)
}

// EdgesHelperFromJSObject returns a wrapped EdgesHelper JavaScript class.
func EdgesHelperFromJSObject(p *js.Object) *EdgesHelper {
	return &EdgesHelper{p: p}
}

// NewEdgesHelper returns a new EdgesHelper object.
func (t *Three) NewEdgesHelper(object, hex, thresholdAngle float64) *EdgesHelper {
	p := t.ctx.Get("EdgesHelper").New(object, hex, thresholdAngle)
	return EdgesHelperFromJSObject(p)
}
