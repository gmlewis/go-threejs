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

// AnimationUtils represents an animationutils.
type AnimationUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AnimationUtils) JSObject() *js.Object { return a.p }

// AnimationUtils returns an AnimationUtils JavaScript class.
func (t *Three) AnimationUtils() *AnimationUtils {
	p := t.ctx.Get("AnimationUtils")
	return AnimationUtilsFromJSObject(p)
}

// AnimationUtilsFromJSObject returns a wrapped AnimationUtils JavaScript class.
func AnimationUtilsFromJSObject(p *js.Object) *AnimationUtils {
	return &AnimationUtils{p: p}
}

// NewAnimationUtils returns a new AnimationUtils object.
func (t *Three) NewAnimationUtils() *AnimationUtils {
	p := t.ctx.Get("AnimationUtils").New()
	return AnimationUtilsFromJSObject(p)
}

// ArraySlice TODO description.
func (a *AnimationUtils) ArraySlice(array, from, to float64) *AnimationUtils {
	a.p.Call("arraySlice", array, from, to)
	return a
}

// ConvertArray TODO description.
func (a *AnimationUtils) ConvertArray(array, typ, forceClone float64) *AnimationUtils {
	a.p.Call("convertArray", array, typ, forceClone)
	return a
}

// IsTypedArray TODO description.
func (a *AnimationUtils) IsTypedArray(object float64) *AnimationUtils {
	a.p.Call("isTypedArray", object)
	return a
}

// GetKeyframeOrder TODO description.
func (a *AnimationUtils) GetKeyframeOrder(times float64) *AnimationUtils {
	a.p.Call("getKeyframeOrder", times)
	return a
}

// SortedArray TODO description.
func (a *AnimationUtils) SortedArray(values, stride, order float64) *AnimationUtils {
	a.p.Call("sortedArray", values, stride, order)
	return a
}

// FlattenJSON TODO description.
func (a *AnimationUtils) FlattenJSON(jsonKeys, times, values, valuePropertyName float64) *AnimationUtils {
	a.p.Call("flattenJSON", jsonKeys, times, values, valuePropertyName)
	return a
}
