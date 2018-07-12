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

// PropertyBinding represents a propertybinding.
type PropertyBinding struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (p *PropertyBinding) JSObject() *js.Object { return p.p }

// PropertyBinding returns a PropertyBinding JavaScript class.
func (t *Three) PropertyBinding() *PropertyBinding {
	p := t.ctx.Get("PropertyBinding")
	return PropertyBindingFromJSObject(p)
}

// PropertyBindingFromJSObject returns a wrapped PropertyBinding JavaScript class.
func PropertyBindingFromJSObject(p *js.Object) *PropertyBinding {
	return &PropertyBinding{p: p}
}

// NewPropertyBinding returns a new PropertyBinding object.
func (t *Three) NewPropertyBinding(rootNode, path, parsedPath float64) *PropertyBinding {
	p := t.ctx.Get("PropertyBinding").New(rootNode, path, parsedPath)
	return PropertyBindingFromJSObject(p)
}

// Bind TODO description.
func (p *PropertyBinding) Bind() *PropertyBinding {
	p.p.Call("bind")
	return p
}

// Unbind TODO description.
func (p *PropertyBinding) Unbind() *PropertyBinding {
	p.p.Call("unbind")
	return p
}

// GetValue TODO description.
func (p *PropertyBinding) GetValue(array, offset float64) *PropertyBinding {
	p.p.Call("getValue", array, offset)
	return p
}

// SetValue TODO description.
func (p *PropertyBinding) SetValue(array, offset float64) *PropertyBinding {
	p.p.Call("setValue", array, offset)
	return p
}
