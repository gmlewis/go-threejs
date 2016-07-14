// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
