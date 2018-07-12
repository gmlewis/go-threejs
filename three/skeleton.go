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

// Skeleton uses an array of bones to create a skeleton that can be used by a SkinnedMesh.
//
// http://threejs.org/docs/index.html#Reference/Objects/Skeleton
type Skeleton struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (s *Skeleton) JSObject() *js.Object { return s.p }

// Skeleton returns a Skeleton JavaScript class.
func (t *Three) Skeleton() *Skeleton {
	p := t.ctx.Get("Skeleton")
	return SkeletonFromJSObject(p)
}

// SkeletonFromJSObject returns a wrapped Skeleton JavaScript class.
func SkeletonFromJSObject(p *js.Object) *Skeleton {
	return &Skeleton{p: p}
}

// NewSkeleton returns a new Skeleton object.
//
//     bones — The array of bones
//     boneInverses — (optional) An array of Matrix4s
//     useVertexTexture — (optional) Whether or not to use a vertex texture in the shader.
func (t *Three) NewSkeleton(bones, boneInverses []*js.Object, useVertexTexture bool) *Skeleton {
	p := t.ctx.Get("Skeleton").New(bones, boneInverses, useVertexTexture)
	return SkeletonFromJSObject(p)
}
