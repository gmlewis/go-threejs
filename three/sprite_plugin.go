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

// SpritePlugin returns a SpritePlugin JavaScript class.
func (t *Three) SpritePlugin() *SpritePlugin {
	p := t.ctx.Get("SpritePlugin")
	return SpritePluginFromJSObject(p)
}

// SpritePluginFromJSObject returns a wrapped SpritePlugin JavaScript class.
func SpritePluginFromJSObject(p *js.Object) *SpritePlugin {
	return &SpritePlugin{p: p}
}

// NewSpritePlugin returns a new SpritePlugin object.
func (t *Three) NewSpritePlugin(renderer, sprites float64) *SpritePlugin {
	p := t.ctx.Get("SpritePlugin").New(renderer, sprites)
	return SpritePluginFromJSObject(p)
}
