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

// LoadingManager represents a loadingmanager.
type LoadingManager struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LoadingManager) JSObject() *js.Object { return l.p }

// LoadingManager returns a LoadingManager JavaScript class.
func (t *Three) LoadingManager() *LoadingManager {
	p := t.ctx.Get("LoadingManager")
	return LoadingManagerFromJSObject(p)
}

// LoadingManagerFromJSObject returns a wrapped LoadingManager JavaScript class.
func LoadingManagerFromJSObject(p *js.Object) *LoadingManager {
	return &LoadingManager{p: p}
}

// NewLoadingManager returns a new LoadingManager object.
func (t *Three) NewLoadingManager(onLoad, onProgress, onError float64) *LoadingManager {
	p := t.ctx.Get("LoadingManager").New(onLoad, onProgress, onError)
	return LoadingManagerFromJSObject(p)
}
