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

// CompressedTextureLoader represents a compressedtextureloader.
type CompressedTextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CompressedTextureLoader) JSObject() *js.Object { return c.p }

// CompressedTextureLoader returns a CompressedTextureLoader JavaScript class.
func (t *Three) CompressedTextureLoader() *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader")
	return CompressedTextureLoaderFromJSObject(p)
}

// CompressedTextureLoaderFromJSObject returns a wrapped CompressedTextureLoader JavaScript class.
func CompressedTextureLoaderFromJSObject(p *js.Object) *CompressedTextureLoader {
	return &CompressedTextureLoader{p: p}
}

// NewCompressedTextureLoader returns a new CompressedTextureLoader object.
func (t *Three) NewCompressedTextureLoader(manager float64) *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader").New(manager)
	return CompressedTextureLoaderFromJSObject(p)
}

// Load TODO description.
func (c *CompressedTextureLoader) Load(url, onLoad, onProgress, onError float64) *CompressedTextureLoader {
	c.p.Call("load", url, onLoad, onProgress, onError)
	return c
}

// SetPath TODO description.
func (c *CompressedTextureLoader) SetPath(value float64) *CompressedTextureLoader {
	c.p.Call("setPath", value)
	return c
}
