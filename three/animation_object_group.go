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

// AnimationObjectGroup represents an animationobjectgroup.
type AnimationObjectGroup struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AnimationObjectGroup) JSObject() *js.Object { return a.p }

// AnimationObjectGroup returns an AnimationObjectGroup JavaScript class.
func (t *Three) AnimationObjectGroup() *AnimationObjectGroup {
	p := t.ctx.Get("AnimationObjectGroup")
	return AnimationObjectGroupFromJSObject(p)
}

// AnimationObjectGroupFromJSObject returns a wrapped AnimationObjectGroup JavaScript class.
func AnimationObjectGroupFromJSObject(p *js.Object) *AnimationObjectGroup {
	return &AnimationObjectGroup{p: p}
}

// NewAnimationObjectGroup returns a new AnimationObjectGroup object.
func (t *Three) NewAnimationObjectGroup(varArgs JSObject) *AnimationObjectGroup {
	p := t.ctx.Get("AnimationObjectGroup").New(varArgs.JSObject())
	return AnimationObjectGroupFromJSObject(p)
}

// Total returns the total-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) Total() float64 {
	return a.p.Get("total").Float()
}

// InUse returns the inUse-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) InUse() float64 {
	return a.p.Get("inUse").Float()
}

// BindingsPerObject returns the bindingsPerObject-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) BindingsPerObject() float64 {
	return a.p.Get("bindingsPerObject").Float()
}

// Add TODO description.
func (a *AnimationObjectGroup) Add(varArgs JSObject) *AnimationObjectGroup {
	a.p.Call("add", varArgs.JSObject())
	return a
}

// Remove TODO description.
func (a *AnimationObjectGroup) Remove(varArgs JSObject) *AnimationObjectGroup {
	a.p.Call("remove", varArgs.JSObject())
	return a
}

// Uncache TODO description.
func (a *AnimationObjectGroup) Uncache(varArgs JSObject) *AnimationObjectGroup {
	a.p.Call("uncache", varArgs.JSObject())
	return a
}

// Subscribe TODO description.
func (a *AnimationObjectGroup) Subscribe(path, parsedPath float64) *AnimationObjectGroup {
	a.p.Call("subscribe_", path, parsedPath)
	return a
}

// Unsubscribe TODO description.
func (a *AnimationObjectGroup) Unsubscribe(path float64) *AnimationObjectGroup {
	a.p.Call("unsubscribe_", path)
	return a
}
