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

// InstancedInterleavedBuffer represents an instancedinterleavedbuffer.
type InstancedInterleavedBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InstancedInterleavedBuffer) JSObject() *js.Object { return i.p }

// InstancedInterleavedBuffer returns an InstancedInterleavedBuffer JavaScript class.
func (t *Three) InstancedInterleavedBuffer() *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer")
	return InstancedInterleavedBufferFromJSObject(p)
}

// InstancedInterleavedBufferFromJSObject returns a wrapped InstancedInterleavedBuffer JavaScript class.
func InstancedInterleavedBufferFromJSObject(p *js.Object) *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{p: p}
}

// NewInstancedInterleavedBuffer returns a new InstancedInterleavedBuffer object.
func (t *Three) NewInstancedInterleavedBuffer(array, stride, meshPerAttribute float64) *InstancedInterleavedBuffer {
	p := t.ctx.Get("InstancedInterleavedBuffer").New(array, stride, meshPerAttribute)
	return InstancedInterleavedBufferFromJSObject(p)
}

// Copy TODO description.
func (i *InstancedInterleavedBuffer) Copy(source *InstancedInterleavedBuffer) *InstancedInterleavedBuffer {
	i.p.Call("copy", source.p)
	return i
}
