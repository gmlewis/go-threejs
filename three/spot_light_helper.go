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
func (t *SpotLightHelper) JSObject() *js.Object { return t.p }

// SpotLightHelper returns a SpotLightHelper object.
func (t *Three) SpotLightHelper() *SpotLightHelper {
	p := t.ctx.Get("SpotLightHelper")
	return &SpotLightHelper{p: p}
}

// New returns a new SpotLightHelper object.
func (t *SpotLightHelper) New(light float64) *SpotLightHelper {
	p := t.p.New(light)
	return &SpotLightHelper{p: p}
}
