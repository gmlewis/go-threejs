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

// BoxHelper represents an object to show a wireframe box
// (with no face diagonals) around an object.
//
// http://threejs.org/docs/index.html#Reference/Extras.Helpers/BoxHelper
type BoxHelper struct{ *Line }

// JSObject returns the underlying *js.Object.
func (b *BoxHelper) JSObject() *js.Object { return b.p }

// BoxHelper returns a BoxHelper JavaScript class.
func (t *Three) BoxHelper() *BoxHelper {
	p := t.ctx.Get("BoxHelper")
	return BoxHelperFromJSObject(p)
}

// BoxHelperFromJSObject returns a wrapped BoxHelper JavaScript class.
func BoxHelperFromJSObject(p *js.Object) *BoxHelper {
	return &BoxHelper{LineFromJSObject(p)}
}

// NewBoxHelper returns a new BoxHelper object.
func (t *Three) NewBoxHelper(object JSObject) *BoxHelper {
	p := t.ctx.Get("BoxHelper").New(object.JSObject())
	return BoxHelperFromJSObject(p)
}

// Update updates the BoxHelper based on the object property.
func (b *BoxHelper) Update() *BoxHelper {
	b.p.Call("update")
	return b
}
