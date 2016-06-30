// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// XHRLoader represents a xhrloader.
type XHRLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *XHRLoader) JSObject() *js.Object { return t.p }

// XHRLoader returns a XHRLoader object.
func (t *Three) XHRLoader() *XHRLoader {
	p := t.ctx.Get("XHRLoader")
	return &XHRLoader{p: p}
}

// New returns a new XHRLoader object.
func (t *XHRLoader) New(manager float64) *XHRLoader {
	p := t.p.New(manager)
	return &XHRLoader{p: p}
}

// Load TODO description.
func (x *XHRLoader) Load(url, onLoad, onProgress, onError float64) *XHRLoader {
	x.p.Call("load", url, onLoad, onProgress, onError)
	return x
}

// SetPath TODO description.
func (x *XHRLoader) SetPath(value float64) *XHRLoader {
	x.p.Call("setPath", value)
	return x
}

// SetResponseType TODO description.
func (x *XHRLoader) SetResponseType(value float64) *XHRLoader {
	x.p.Call("setResponseType", value)
	return x
}

// SetWithCredentials TODO description.
func (x *XHRLoader) SetWithCredentials(value float64) *XHRLoader {
	x.p.Call("setWithCredentials", value)
	return x
}
