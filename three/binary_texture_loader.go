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

// BinaryTextureLoader represents a binarytextureloader.
type BinaryTextureLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (b *BinaryTextureLoader) JSObject() *js.Object { return b.p }

// BinaryTextureLoader returns a BinaryTextureLoader JavaScript class.
func (t *Three) BinaryTextureLoader() *BinaryTextureLoader {
	p := t.ctx.Get("BinaryTextureLoader")
	return BinaryTextureLoaderFromJSObject(p)
}

// BinaryTextureLoaderFromJSObject returns a wrapped BinaryTextureLoader JavaScript class.
func BinaryTextureLoaderFromJSObject(p *js.Object) *BinaryTextureLoader {
	return &BinaryTextureLoader{p: p}
}

// NewBinaryTextureLoader returns a new BinaryTextureLoader object.
func (t *Three) NewBinaryTextureLoader() *BinaryTextureLoader {
	p := t.ctx.Get("BinaryTextureLoader").New()
	return BinaryTextureLoaderFromJSObject(p)
}

// Load TODO description.
func (b *BinaryTextureLoader) Load(url, onLoad, onProgress, onError float64) *BinaryTextureLoader {
	b.p.Call("load", url, onLoad, onProgress, onError)
	return b
}
