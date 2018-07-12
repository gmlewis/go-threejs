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

// XHRLoader represents a xhrloader.
type XHRLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (x *XHRLoader) JSObject() *js.Object { return x.p }

// XHRLoader returns a XHRLoader JavaScript class.
func (t *Three) XHRLoader() *XHRLoader {
	p := t.ctx.Get("XHRLoader")
	return XHRLoaderFromJSObject(p)
}

// XHRLoaderFromJSObject returns a wrapped XHRLoader JavaScript class.
func XHRLoaderFromJSObject(p *js.Object) *XHRLoader {
	return &XHRLoader{p: p}
}

// NewXHRLoader returns a new XHRLoader object.
func (t *Three) NewXHRLoader(manager float64) *XHRLoader {
	p := t.ctx.Get("XHRLoader").New(manager)
	return XHRLoaderFromJSObject(p)
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
