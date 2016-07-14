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
func (f *FaceNormalsHelper) JSObject() *js.Object { return f.p }

// FaceNormalsHelper returns a FaceNormalsHelper JavaScript class.
func (t *Three) FaceNormalsHelper() *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper")
	return FaceNormalsHelperFromJSObject(p)
}

// FaceNormalsHelperFromJSObject returns a wrapped FaceNormalsHelper JavaScript class.
func FaceNormalsHelperFromJSObject(p *js.Object) *FaceNormalsHelper {
	return &FaceNormalsHelper{p: p}
}

// NewFaceNormalsHelper returns a new FaceNormalsHelper object.
func (t *Three) NewFaceNormalsHelper(object, size, hex, linewidth float64) *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper").New(object, size, hex, linewidth)
	return FaceNormalsHelperFromJSObject(p)
}
