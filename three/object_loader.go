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

// ObjectLoader represents an objectloader.
type ObjectLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (o *ObjectLoader) JSObject() *js.Object { return o.p }

// ObjectLoader returns an ObjectLoader JavaScript class.
func (t *Three) ObjectLoader() *ObjectLoader {
	p := t.ctx.Get("ObjectLoader")
	return ObjectLoaderFromJSObject(p)
}

// ObjectLoaderFromJSObject returns a wrapped ObjectLoader JavaScript class.
func ObjectLoaderFromJSObject(p *js.Object) *ObjectLoader {
	return &ObjectLoader{p: p}
}

// NewObjectLoader returns a new ObjectLoader object.
func (t *Three) NewObjectLoader(manager float64) *ObjectLoader {
	p := t.ctx.Get("ObjectLoader").New(manager)
	return ObjectLoaderFromJSObject(p)
}

// Load TODO description.
func (o *ObjectLoader) Load(url, onLoad, onProgress, onError float64) *ObjectLoader {
	o.p.Call("load", url, onLoad, onProgress, onError)
	return o
}

// SetTexturePath TODO description.
func (o *ObjectLoader) SetTexturePath(value float64) *ObjectLoader {
	o.p.Call("setTexturePath", value)
	return o
}

// SetCrossOrigin TODO description.
func (o *ObjectLoader) SetCrossOrigin(value float64) *ObjectLoader {
	o.p.Call("setCrossOrigin", value)
	return o
}

// Parse TODO description.
func (o *ObjectLoader) Parse(json, onLoad float64) *ObjectLoader {
	o.p.Call("parse", json, onLoad)
	return o
}

// ParseGeometries TODO description.
func (o *ObjectLoader) ParseGeometries(json float64) *ObjectLoader {
	o.p.Call("parseGeometries", json)
	return o
}

// ParseMaterials TODO description.
func (o *ObjectLoader) ParseMaterials(json, textures float64) *ObjectLoader {
	o.p.Call("parseMaterials", json, textures)
	return o
}

// ParseAnimations TODO description.
func (o *ObjectLoader) ParseAnimations(json float64) *ObjectLoader {
	o.p.Call("parseAnimations", json)
	return o
}

// ParseImages TODO description.
func (o *ObjectLoader) ParseImages(json, onLoad float64) *ObjectLoader {
	o.p.Call("parseImages", json, onLoad)
	return o
}

// ParseTextures TODO description.
func (o *ObjectLoader) ParseTextures(json, images float64) *ObjectLoader {
	o.p.Call("parseTextures", json, images)
	return o
}

// ParseObject TODO description.
func (o *ObjectLoader) ParseObject() *ObjectLoader {
	o.p.Call("parseObject")
	return o
}
