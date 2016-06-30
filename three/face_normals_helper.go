// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FaceNormalsHelper represents a facenormalshelper.
type FaceNormalsHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *FaceNormalsHelper) JSObject() *js.Object { return t.p }

// FaceNormalsHelper returns a FaceNormalsHelper object.
func (t *Three) FaceNormalsHelper() *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper")
	return &FaceNormalsHelper{p: p}
}

// New returns a new FaceNormalsHelper object.
func (t *FaceNormalsHelper) New(object, size, hex, linewidth float64) *FaceNormalsHelper {
	p := t.p.New(object, size, hex, linewidth)
	return &FaceNormalsHelper{p: p}
}
