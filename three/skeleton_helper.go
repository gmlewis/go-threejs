// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
