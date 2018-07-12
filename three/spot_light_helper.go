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

// SpotLightHelper represents a spotlighthelper.
type SpotLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SpotLightHelper) JSObject() *js.Object { return s.p }

// SpotLightHelper returns a SpotLightHelper JavaScript class.
func (t *Three) SpotLightHelper() *SpotLightHelper {
	p := t.ctx.Get("SpotLightHelper")
	return SpotLightHelperFromJSObject(p)
}

// SpotLightHelperFromJSObject returns a wrapped SpotLightHelper JavaScript class.
func SpotLightHelperFromJSObject(p *js.Object) *SpotLightHelper {
	return &SpotLightHelper{p: p}
}

// NewSpotLightHelper returns a new SpotLightHelper object.
func (t *Three) NewSpotLightHelper(light float64) *SpotLightHelper {
	p := t.ctx.Get("SpotLightHelper").New(light)
	return SpotLightHelperFromJSObject(p)
}
