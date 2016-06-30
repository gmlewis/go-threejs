// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DataTexture represents a datatexture.
type DataTexture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DataTexture) JSObject() *js.Object { return d.p }

// DataTexture returns a DataTexture object.
func (t *Three) DataTexture() *DataTexture {
	p := t.ctx.Get("DataTexture")
	return &DataTexture{p: p}
}

// New returns a new DataTexture object.
func (d *DataTexture) New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *DataTexture {
	p := d.p.New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return &DataTexture{p: p}
}
