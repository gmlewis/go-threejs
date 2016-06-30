// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpritePlugin represents a spriteplugin.
type SpritePlugin struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *SpritePlugin) JSObject() *js.Object { return s.p }

// SpritePlugin returns a SpritePlugin object.
func (t *Three) SpritePlugin() *SpritePlugin {
	p := t.ctx.Get("SpritePlugin")
	return &SpritePlugin{p: p}
}

// New returns a new SpritePlugin object.
func (s *SpritePlugin) New(renderer, sprites float64) *SpritePlugin {
	p := s.p.New(renderer, sprites)
	return &SpritePlugin{p: p}
}
