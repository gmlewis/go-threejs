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

// SpotLightHelper returns a SpotLightHelper object.
func (t *Three) SpotLightHelper() *SpotLightHelper {
	p := t.ctx.Get("SpotLightHelper")
	return &SpotLightHelper{p: p}
}

// New returns a new SpotLightHelper object.
func (s *SpotLightHelper) New(light float64) *SpotLightHelper {
	p := s.p.New(light)
	return &SpotLightHelper{p: p}
}
