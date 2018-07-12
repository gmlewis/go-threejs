// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
