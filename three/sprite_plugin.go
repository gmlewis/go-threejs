// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpritePlugin allows Sprites to be rendered in the WebglRenderer. This plugin is automatically
// loaded in the Webglrenderer.
//
// http://threejs.org/docs/index.html#Reference/Renderers.WebGL.Plugins/SpritePlugin
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

// NewSpritePlugin creates a new SpritePlugin object.
func (t *Three) NewSpritePlugin(renderer, sprites float64) *SpritePlugin {
	p := t.ctx.Get("SpritePlugin").New(renderer, sprites)
	return SpritePluginFromJSObject(p)
}

// TODO:
//Render renders the sprites defined in the scene. This gets automatically called as post-render function
//to draw the lensflares.
//     scene -- The scene to render.
//     camera -- The camera to render.
