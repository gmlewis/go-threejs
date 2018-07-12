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

// ImageLoader represents an imageloader.
type ImageLoader struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *ImageLoader) JSObject() *js.Object { return i.p }

// ImageLoader returns an ImageLoader JavaScript class.
func (t *Three) ImageLoader() *ImageLoader {
	p := t.ctx.Get("ImageLoader")
	return ImageLoaderFromJSObject(p)
}

// ImageLoaderFromJSObject returns a wrapped ImageLoader JavaScript class.
func ImageLoaderFromJSObject(p *js.Object) *ImageLoader {
	return &ImageLoader{p: p}
}

// NewImageLoader returns a new ImageLoader object.
func (t *Three) NewImageLoader(manager float64) *ImageLoader {
	p := t.ctx.Get("ImageLoader").New(manager)
	return ImageLoaderFromJSObject(p)
}

// Load TODO description.
func (i *ImageLoader) Load(url, onLoad, onProgress, onError float64) *ImageLoader {
	i.p.Call("load", url, onLoad, onProgress, onError)
	return i
}

// SetCrossOrigin TODO description.
func (i *ImageLoader) SetCrossOrigin(value float64) *ImageLoader {
	i.p.Call("setCrossOrigin", value)
	return i
}

// SetPath TODO description.
func (i *ImageLoader) SetPath(value float64) *ImageLoader {
	i.p.Call("setPath", value)
	return i
}
