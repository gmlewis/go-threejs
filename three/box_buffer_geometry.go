// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxBufferGeometry represents a BufferGeometry port of BoxGeometry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/BoxBufferGeometry
type BoxBufferGeometry struct{ *BufferGeometry }

// JSObject returns the underlying *js.Object.
func (b *BoxBufferGeometry) JSObject() *js.Object { return b.p }

// BoxBufferGeometry returns a BoxBufferGeometry JavaScript class.
func (t *Three) BoxBufferGeometry() *BoxBufferGeometry {
	p := t.ctx.Get("BoxBufferGeometry")
	return BoxBufferGeometryFromJSObject(p)
}

// BoxBufferGeometryFromJSObject returns a wrapped BoxBufferGeometry JavaScript class.
func BoxBufferGeometryFromJSObject(p *js.Object) *BoxBufferGeometry {
	return &BoxBufferGeometry{BufferGeometryFromJSObject(p)}
}

// BoxBufferGeometryOpts represents optional arguments that can be
// passed to NewBoxBufferGeometry.
type BoxBufferGeometryOpts struct {
	widthSegments, heightSegments, depthSegments int // all default to 1
}

// NewBoxBufferGeometry returns a new BoxBufferGeometry object.
func (t *Three) NewBoxBufferGeometry(width, height, depth int, opts *BoxBufferGeometryOpts) *BoxBufferGeometry {
	var p *js.Object
	if opts != nil {
		p = t.ctx.Get("BoxBufferGeometry").New(width, height, depth, opts.widthSegments, opts.heightSegments, opts.depthSegments)
	} else {
		p = t.ctx.Get("BoxBufferGeometry").New(width, height, depth)
	}
	return BoxBufferGeometryFromJSObject(p)
}
