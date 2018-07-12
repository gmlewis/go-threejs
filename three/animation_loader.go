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

// AnimationLoader represents an animationloader.
type AnimationLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AnimationLoader) JSObject() *js.Object { return a.p }

// AnimationLoader returns an AnimationLoader JavaScript class.
func (t *Three) AnimationLoader() *AnimationLoader {
	p := t.ctx.Get("AnimationLoader")
	return AnimationLoaderFromJSObject(p)
}

// AnimationLoaderFromJSObject returns a wrapped AnimationLoader JavaScript class.
func AnimationLoaderFromJSObject(p *js.Object) *AnimationLoader {
	return &AnimationLoader{p: p}
}

// NewAnimationLoader returns a new AnimationLoader object.
func (t *Three) NewAnimationLoader(manager float64) *AnimationLoader {
	p := t.ctx.Get("AnimationLoader").New(manager)
	return AnimationLoaderFromJSObject(p)
}

// Load TODO description.
func (a *AnimationLoader) Load(url, onLoad, onProgress, onError float64) *AnimationLoader {
	a.p.Call("load", url, onLoad, onProgress, onError)
	return a
}

// Parse TODO description.
func (a *AnimationLoader) Parse(json, onLoad float64) *AnimationLoader {
	a.p.Call("parse", json, onLoad)
	return a
}
