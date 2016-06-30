// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CatmullRomCurve3 represents a catmullromcurve3.
type CatmullRomCurve3 struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *CatmullRomCurve3) JSObject() *js.Object { return t.p }

// CatmullRomCurve3 returns a CatmullRomCurve3 object.
func (t *Three) CatmullRomCurve3() *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3")
	return &CatmullRomCurve3{p: p}
}

// New returns a new CatmullRomCurve3 object.
func (t *CatmullRomCurve3) New() *CatmullRomCurve3 {
	p := t.p.New()
	return &CatmullRomCurve3{p: p}
}