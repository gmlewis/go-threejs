// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// HemisphereLightHelper represents a hemispherelighthelper.
type HemisphereLightHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (h *HemisphereLightHelper) JSObject() *js.Object { return h.p }

// HemisphereLightHelper returns a HemisphereLightHelper object.
func (t *Three) HemisphereLightHelper() *HemisphereLightHelper {
	p := t.ctx.Get("HemisphereLightHelper")
	return &HemisphereLightHelper{p: p}
}

// New returns a new HemisphereLightHelper object.
func (h *HemisphereLightHelper) New(light, sphereSize float64) *HemisphereLightHelper {
	p := h.p.New(light, sphereSize)
	return &HemisphereLightHelper{p: p}
}
