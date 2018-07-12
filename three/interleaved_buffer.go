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

// InterleavedBuffer represents an interleavedbuffer.
type InterleavedBuffer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *InterleavedBuffer) JSObject() *js.Object { return i.p }

// InterleavedBuffer returns an InterleavedBuffer JavaScript class.
func (t *Three) InterleavedBuffer() *InterleavedBuffer {
	p := t.ctx.Get("InterleavedBuffer")
	return InterleavedBufferFromJSObject(p)
}

// InterleavedBufferFromJSObject returns a wrapped InterleavedBuffer JavaScript class.
func InterleavedBufferFromJSObject(p *js.Object) *InterleavedBuffer {
	return &InterleavedBuffer{p: p}
}

// NewInterleavedBuffer returns a new InterleavedBuffer object.
func (t *Three) NewInterleavedBuffer(array, stride float64) *InterleavedBuffer {
	p := t.ctx.Get("InterleavedBuffer").New(array, stride)
	return InterleavedBufferFromJSObject(p)
}

// Length returns the length-component of the InterleavedBuffer.
func (i *InterleavedBuffer) Length() float64 {
	return i.p.Get("length").Float()
}

// Count returns the count-component of the InterleavedBuffer.
func (i *InterleavedBuffer) Count() float64 {
	return i.p.Get("count").Float()
}

// SetNeedsUpdate sets the needsUpdate-component of the InterleavedBuffer.
func (i *InterleavedBuffer) SetNeedsUpdate(value float64) *InterleavedBuffer {
	i.p.Set("needsUpdate", value)
	return i
}

// SetDynamic TODO description.
func (i *InterleavedBuffer) SetDynamic(value float64) *InterleavedBuffer {
	i.p.Call("setDynamic", value)
	return i
}

// Copy TODO description.
func (i *InterleavedBuffer) Copy(source *InterleavedBuffer) *InterleavedBuffer {
	i.p.Call("copy", source.p)
	return i
}

// CopyAt TODO description.
func (i *InterleavedBuffer) CopyAt(index1, attribute, index2 float64) *InterleavedBuffer {
	i.p.Call("copyAt", index1, attribute, index2)
	return i
}

// Set TODO description.
func (i *InterleavedBuffer) Set(value, offset float64) *InterleavedBuffer {
	i.p.Call("set", value, offset)
	return i
}

// Clone TODO description.
func (i *InterleavedBuffer) Clone() *InterleavedBuffer {
	i.p.Call("clone")
	return i
}
