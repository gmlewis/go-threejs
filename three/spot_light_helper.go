// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
