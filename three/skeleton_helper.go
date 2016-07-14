// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SkeletonHelper represents a skeletonhelper.
type SkeletonHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SkeletonHelper) JSObject() *js.Object { return s.p }

// SkeletonHelper returns a SkeletonHelper JavaScript class.
func (t *Three) SkeletonHelper() *SkeletonHelper {
	p := t.ctx.Get("SkeletonHelper")
	return SkeletonHelperFromJSObject(p)
}

// SkeletonHelperFromJSObject returns a wrapped SkeletonHelper JavaScript class.
func SkeletonHelperFromJSObject(p *js.Object) *SkeletonHelper {
	return &SkeletonHelper{p: p}
}

// NewSkeletonHelper returns a new SkeletonHelper object.
func (t *Three) NewSkeletonHelper(object float64) *SkeletonHelper {
	p := t.ctx.Get("SkeletonHelper").New(object)
	return SkeletonHelperFromJSObject(p)
}

// GetBoneList TODO description.
func (s *SkeletonHelper) GetBoneList(object float64) *SkeletonHelper {
	s.p.Call("getBoneList", object)
	return s
}
