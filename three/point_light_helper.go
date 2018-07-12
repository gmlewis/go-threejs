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

// PointLightHelper represents a pointlighthelper.
type PointLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PointLightHelper) JSObject() *js.Object { return p.p }

// PointLightHelper returns a PointLightHelper JavaScript class.
func (t *Three) PointLightHelper() *PointLightHelper {
	p := t.ctx.Get("PointLightHelper")
	return PointLightHelperFromJSObject(p)
}

// PointLightHelperFromJSObject returns a wrapped PointLightHelper JavaScript class.
func PointLightHelperFromJSObject(p *js.Object) *PointLightHelper {
	return &PointLightHelper{p: p}
}

// NewPointLightHelper returns a new PointLightHelper object.
func (t *Three) NewPointLightHelper(light, sphereSize float64) *PointLightHelper {
	p := t.ctx.Get("PointLightHelper").New(light, sphereSize)
	return PointLightHelperFromJSObject(p)
}
