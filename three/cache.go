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

// Cache represents a simple cache used internally by XHRLoader.
//
// http://threejs.org/docs/index.html#Reference/Loaders/Cache
type Cache struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *Cache) JSObject() *js.Object { return c.p }

// Cache returns a Cache JavaScript class.
func (t *Three) Cache() *Cache {
	p := t.ctx.Get("Cache")
	return CacheFromJSObject(p)
}

// CacheFromJSObject returns a wrapped Cache JavaScript class.
func CacheFromJSObject(p *js.Object) *Cache {
	return &Cache{p: p}
}

// NewCache returns a new Cache object.
func (t *Three) NewCache() *Cache {
	p := t.ctx.Get("Cache").New()
	return CacheFromJSObject(p)
}

// Add TODO description.
func (c *Cache) Add(key, file float64) *Cache {
	c.p.Call("add", key, file)
	return c
}

// Get TODO description.
func (c *Cache) Get(key float64) *Cache {
	c.p.Call("get", key)
	return c
}

// Remove TODO description.
func (c *Cache) Remove(key float64) *Cache {
	c.p.Call("remove", key)
	return c
}

// Clear TODO description.
func (c *Cache) Clear() *Cache {
	c.p.Call("clear")
	return c
}
