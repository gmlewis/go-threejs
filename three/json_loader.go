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

import "github.com/gopherjs/gopherjs/js"

// JSONLoader represents a jsonloader.
type JSONLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (j *JSONLoader) JSObject() *js.Object { return j.p }

// JSONLoader returns a JSONLoader JavaScript class.
func (t *Three) JSONLoader() *JSONLoader {
	p := t.ctx.Get("JSONLoader")
	return JSONLoaderFromJSObject(p)
}

// JSONLoaderFromJSObject returns a wrapped JSONLoader JavaScript class.
func JSONLoaderFromJSObject(p *js.Object) *JSONLoader {
	return &JSONLoader{p: p}
}

// NewJSONLoader returns a new JSONLoader object.
func (t *Three) NewJSONLoader() *JSONLoader {
	p := t.ctx.Get("JSONLoader").New()
	return JSONLoaderFromJSObject(p)
}

// StatusDomElement returns the statusDomElement-component of the JSONLoader.
func (j *JSONLoader) StatusDomElement() float64 {
	return j.p.Get("statusDomElement").Float()
}

// JSONLoadFunc is a callback function called by Load.
type JSONLoadFunc func(geometry *Geometry, materials []*Material)

// onJSONLoadWrapperFunc wraps the geometry and materials to a typed LoadFunc.
func onJSONLoadWrapperFunc(onLoad JSONLoadFunc) func(geometry, materials *js.Object) {
	return func(geom, materialArray *js.Object) {
		var materials []*Material
		if materialArray != nil && materialArray != js.Undefined {
			for i := 0; i < materialArray.Length(); i++ {
				materials = append(materials, MaterialFromJSObject(materialArray.Index(i)))
			}
		}
		onLoad(GeometryFromJSObject(geom), materials)
	}
}

// Load TODO description.
func (j *JSONLoader) Load(url string, onLoad JSONLoadFunc, onProgress, onError interface{}) *JSONLoader {
	var onLoadWrapper func(geometry, materials *js.Object)
	if onLoad != nil {
		onLoadWrapper = onJSONLoadWrapperFunc(onLoad)
	}
	switch {
	case onProgress != nil && onError != nil:
		j.p.Call("load", url, onLoadWrapper, onProgress, onError)
	case onProgress != nil:
		j.p.Call("load", url, onLoadWrapper, onProgress)
	default:
		j.p.Call("load", url, onLoadWrapper)
	}
	return j
}

// SetTexturePath TODO description.
func (j *JSONLoader) SetTexturePath(value float64) *JSONLoader {
	j.p.Call("setTexturePath", value)
	return j
}

// Parse TODO description.
func (j *JSONLoader) Parse(json, texturePath float64) *JSONLoader {
	j.p.Call("parse", json, texturePath)
	return j
}
