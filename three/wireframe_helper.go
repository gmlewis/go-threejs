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

// WireframeHelper represents a wireframehelper.
type WireframeHelper struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WireframeHelper) JSObject() *js.Object { return w.p }

// WireframeHelper returns a WireframeHelper JavaScript class.
func (t *Three) WireframeHelper() *WireframeHelper {
	p := t.ctx.Get("WireframeHelper")
	return WireframeHelperFromJSObject(p)
}

// WireframeHelperFromJSObject returns a wrapped WireframeHelper JavaScript class.
func WireframeHelperFromJSObject(p *js.Object) *WireframeHelper {
	return &WireframeHelper{p: p}
}

// NewWireframeHelper returns a new WireframeHelper object.
func (t *Three) NewWireframeHelper(object, hex float64) *WireframeHelper {
	p := t.ctx.Get("WireframeHelper").New(object, hex)
	return WireframeHelperFromJSObject(p)
}
