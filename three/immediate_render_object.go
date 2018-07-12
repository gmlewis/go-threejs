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

// ImmediateRenderObject represents an immediaterenderobject.
type ImmediateRenderObject struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *ImmediateRenderObject) JSObject() *js.Object { return i.p }

// ImmediateRenderObject returns an ImmediateRenderObject JavaScript class.
func (t *Three) ImmediateRenderObject() *ImmediateRenderObject {
	p := t.ctx.Get("ImmediateRenderObject")
	return ImmediateRenderObjectFromJSObject(p)
}

// ImmediateRenderObjectFromJSObject returns a wrapped ImmediateRenderObject JavaScript class.
func ImmediateRenderObjectFromJSObject(p *js.Object) *ImmediateRenderObject {
	return &ImmediateRenderObject{p: p}
}

// NewImmediateRenderObject returns a new ImmediateRenderObject object.
func (t *Three) NewImmediateRenderObject(material float64) *ImmediateRenderObject {
	p := t.ctx.Get("ImmediateRenderObject").New(material)
	return ImmediateRenderObjectFromJSObject(p)
}
