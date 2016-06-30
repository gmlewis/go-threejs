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
func (t *SkeletonHelper) JSObject() *js.Object { return t.p }

// SkeletonHelper returns a SkeletonHelper object.
func (t *Three) SkeletonHelper() *SkeletonHelper {
	p := t.ctx.Get("SkeletonHelper")
	return &SkeletonHelper{p: p}
}

// New returns a new SkeletonHelper object.
func (t *SkeletonHelper) New(object float64) *SkeletonHelper {
	p := t.p.New(object)
	return &SkeletonHelper{p: p}
}

// GetBoneList TODO description.
func (s *SkeletonHelper) GetBoneList(object float64) *SkeletonHelper {
	s.p.Call("getBoneList", object)
	return s
}
