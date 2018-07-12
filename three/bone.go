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

// Bone represents a bone which is part of a skeleton.
// The skeleton in turn is used by the SkinnedMesh.
// Bones are almost identical to a blank Object3D.
//
// http://threejs.org/docs/index.html#Reference/Objects/Bone
type Bone struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (b *Bone) JSObject() *js.Object { return b.p }

// Bone returns a Bone JavaScript class.
func (t *Three) Bone() *Bone {
	p := t.ctx.Get("Bone")
	return BoneFromJSObject(p)
}

// BoneFromJSObject returns a wrapped Bone JavaScript class.
func BoneFromJSObject(p *js.Object) *Bone {
	return &Bone{Object3DFromJSObject(p)}
}

// NewBone returns a new Bone object.
func (t *Three) NewBone(skin float64) *Bone {
	p := t.ctx.Get("Bone").New(skin)
	return BoneFromJSObject(p)
}

// Copy TODO description.
func (b *Bone) Copy(source *Bone) *Bone {
	b.p.Call("copy", source.p)
	return b
}
