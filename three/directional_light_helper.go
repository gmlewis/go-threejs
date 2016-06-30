// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectionalLightHelper represents a directionallighthelper.
type DirectionalLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *DirectionalLightHelper) JSObject() *js.Object { return t.p }

// DirectionalLightHelper returns a DirectionalLightHelper object.
func (t *Three) DirectionalLightHelper() *DirectionalLightHelper {
	p := t.ctx.Get("DirectionalLightHelper")
	return &DirectionalLightHelper{p: p}
}

// New returns a new DirectionalLightHelper object.
func (t *DirectionalLightHelper) New(light, size float64) *DirectionalLightHelper {
	p := t.p.New(light, size)
	return &DirectionalLightHelper{p: p}
}
