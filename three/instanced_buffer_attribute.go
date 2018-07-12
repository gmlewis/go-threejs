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

// InstancedBufferAttribute represents an instancedbufferattribute.
type InstancedBufferAttribute struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InstancedBufferAttribute) JSObject() *js.Object { return i.p }

// InstancedBufferAttribute returns an InstancedBufferAttribute JavaScript class.
func (t *Three) InstancedBufferAttribute() *InstancedBufferAttribute {
	p := t.ctx.Get("InstancedBufferAttribute")
	return InstancedBufferAttributeFromJSObject(p)
}

// InstancedBufferAttributeFromJSObject returns a wrapped InstancedBufferAttribute JavaScript class.
func InstancedBufferAttributeFromJSObject(p *js.Object) *InstancedBufferAttribute {
	return &InstancedBufferAttribute{p: p}
}

// NewInstancedBufferAttribute returns a new InstancedBufferAttribute object.
func (t *Three) NewInstancedBufferAttribute(array, itemSize, meshPerAttribute float64) *InstancedBufferAttribute {
	p := t.ctx.Get("InstancedBufferAttribute").New(array, itemSize, meshPerAttribute)
	return InstancedBufferAttributeFromJSObject(p)
}

// Copy TODO description.
func (i *InstancedBufferAttribute) Copy(source *InstancedBufferAttribute) *InstancedBufferAttribute {
	i.p.Call("copy", source.p)
	return i
}
