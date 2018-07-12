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

// HemisphereLightHelper represents a hemispherelighthelper.
type HemisphereLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (h *HemisphereLightHelper) JSObject() *js.Object { return h.p }

// HemisphereLightHelper returns a HemisphereLightHelper JavaScript class.
func (t *Three) HemisphereLightHelper() *HemisphereLightHelper {
	p := t.ctx.Get("HemisphereLightHelper")
	return HemisphereLightHelperFromJSObject(p)
}

// HemisphereLightHelperFromJSObject returns a wrapped HemisphereLightHelper JavaScript class.
func HemisphereLightHelperFromJSObject(p *js.Object) *HemisphereLightHelper {
	return &HemisphereLightHelper{p: p}
}

// NewHemisphereLightHelper returns a new HemisphereLightHelper object.
func (t *Three) NewHemisphereLightHelper(light, sphereSize float64) *HemisphereLightHelper {
	p := t.ctx.Get("HemisphereLightHelper").New(light, sphereSize)
	return HemisphereLightHelperFromJSObject(p)
}
