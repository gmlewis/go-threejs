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

// LensFlarePlugin represents a lensflareplugin.
type LensFlarePlugin struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LensFlarePlugin) JSObject() *js.Object { return l.p }

// LensFlarePlugin returns a LensFlarePlugin JavaScript class.
func (t *Three) LensFlarePlugin() *LensFlarePlugin {
	p := t.ctx.Get("LensFlarePlugin")
	return LensFlarePluginFromJSObject(p)
}

// LensFlarePluginFromJSObject returns a wrapped LensFlarePlugin JavaScript class.
func LensFlarePluginFromJSObject(p *js.Object) *LensFlarePlugin {
	return &LensFlarePlugin{p: p}
}

// NewLensFlarePlugin returns a new LensFlarePlugin object.
func (t *Three) NewLensFlarePlugin(renderer, flares float64) *LensFlarePlugin {
	p := t.ctx.Get("LensFlarePlugin").New(renderer, flares)
	return LensFlarePluginFromJSObject(p)
}
