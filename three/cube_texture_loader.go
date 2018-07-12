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

// CubeTextureLoader represents a cubetextureloader.
type CubeTextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CubeTextureLoader) JSObject() *js.Object { return c.p }

// CubeTextureLoader returns a CubeTextureLoader JavaScript class.
func (t *Three) CubeTextureLoader() *CubeTextureLoader {
	p := t.ctx.Get("CubeTextureLoader")
	return CubeTextureLoaderFromJSObject(p)
}

// CubeTextureLoaderFromJSObject returns a wrapped CubeTextureLoader JavaScript class.
func CubeTextureLoaderFromJSObject(p *js.Object) *CubeTextureLoader {
	return &CubeTextureLoader{p: p}
}

// NewCubeTextureLoader returns a new CubeTextureLoader object.
func (t *Three) NewCubeTextureLoader() *CubeTextureLoader {
	p := t.ctx.Get("CubeTextureLoader").New()
	return CubeTextureLoaderFromJSObject(p)
}

// CubeTextureLoadFunc is a callback function called by Load.
type CubeTextureLoadFunc func(geometry *CubeTexture, materials []*Material)

// onCubeTextureLoadWrapperFunc wraps the geometry and materials to a typed LoadFunc.
func onCubeTextureLoadWrapperFunc(onLoad CubeTextureLoadFunc) func(geometry, materials *js.Object) {
	return func(geom, materialArray *js.Object) {
		var materials []*Material
		if materialArray != nil && materialArray != js.Undefined {
			for i := 0; i < materialArray.Length(); i++ {
				materials = append(materials, MaterialFromJSObject(materialArray.Index(i)))
			}
		}
		onLoad(CubeTextureFromJSObject(geom), materials)
	}
}

// Load TODO description.
func (c *CubeTextureLoader) Load(urls []string, onLoad CubeTextureLoadFunc, onProgress, onError interface{}) *CubeTextureLoader {
	var onLoadWrapper func(geometry, materials *js.Object)
	if onLoad != nil {
		onLoadWrapper = onCubeTextureLoadWrapperFunc(onLoad)
	}
	switch {
	case onProgress != nil && onError != nil:
		c.p.Call("load", urls, onLoadWrapper, onProgress, onError)
	case onProgress != nil:
		c.p.Call("load", urls, onLoadWrapper, onProgress)
	default:
		c.p.Call("load", urls, onLoadWrapper)
	}
	return c
}

// SetCrossOrigin TODO description.
func (c *CubeTextureLoader) SetCrossOrigin(value string) *CubeTextureLoader {
	c.p.Call("setCrossOrigin", value)
	return c
}

// SetPath TODO description.
func (c *CubeTextureLoader) SetPath(value string) *CubeTextureLoader {
	c.p.Call("setPath", value)
	return c
}
