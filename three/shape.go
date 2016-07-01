// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shape represents a shape.
type Shape struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Shape) JSObject() *js.Object { return s.p }

// Shape returns a Shape JavaScript class.
func (t *Three) Shape() *Shape {
	p := t.ctx.Get("Shape")
	return &Shape{p: p}
}

// NewShape returns a new Shape object.
func (t *Three) NewShape() *Shape {
	p := t.ctx.Get("Shape").New()
	return &Shape{p: p}
}

// Extrude TODO description.
func (s *Shape) Extrude(options float64) *Shape {
	s.p.Call("extrude", options)
	return s
}

// MakeGeometry TODO description.
func (s *Shape) MakeGeometry(options float64) *Shape {
	s.p.Call("makeGeometry", options)
	return s
}

// GetPointsHoles TODO description.
func (s *Shape) GetPointsHoles(divisions float64) *Shape {
	s.p.Call("getPointsHoles", divisions)
	return s
}

// ExtractAllPoints TODO description.
func (s *Shape) ExtractAllPoints(divisions float64) *Shape {
	s.p.Call("extractAllPoints", divisions)
	return s
}

// ExtractPoints TODO description.
func (s *Shape) ExtractPoints(divisions float64) *Shape {
	s.p.Call("extractPoints", divisions)
	return s
}
