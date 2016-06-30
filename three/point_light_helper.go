// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PointLightHelper represents a pointlighthelper.
type PointLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PointLightHelper) JSObject() *js.Object { return p.p }

// PointLightHelper returns a PointLightHelper object.
func (t *Three) PointLightHelper() *PointLightHelper {
	p := t.ctx.Get("PointLightHelper")
	return &PointLightHelper{p: p}
}

// New returns a new PointLightHelper object.
func (p *PointLightHelper) New(light, sphereSize float64) *PointLightHelper {
	t := p.p.New(light, sphereSize)
	return &PointLightHelper{p: t}
}
