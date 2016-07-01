// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Font represents a font.
type Font struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (f *Font) JSObject() *js.Object { return f.p }

// Font returns a Font JavaScript class.
func (t *Three) Font() *Font {
	p := t.ctx.Get("Font")
	return &Font{p: p}
}

// NewFont returns a new Font object.
func (t *Three) NewFont(data float64) *Font {
	p := t.ctx.Get("Font").New(data)
	return &Font{p: p}
}

// GenerateShapes TODO description.
func (f *Font) GenerateShapes(text, size, divisions float64) *Font {
	f.p.Call("generateShapes", text, size, divisions)
	return f
}
