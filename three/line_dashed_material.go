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

// LineDashedMaterial represents a material for drawing wireframe-style geometries with dashed lines.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineDashedMaterial
type LineDashedMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LineDashedMaterial) JSObject() *js.Object { return l.p }

// LineDashedMaterial returns a LineDashedMaterial JavaScript class.
func (t *Three) LineDashedMaterial() *LineDashedMaterial {
	p := t.ctx.Get("LineDashedMaterial")
	return LineDashedMaterialFromJSObject(p)
}

// LineDashedMaterialFromJSObject returns a wrapped LineDashedMaterial JavaScript class.
func LineDashedMaterialFromJSObject(p *js.Object) *LineDashedMaterial {
	return &LineDashedMaterial{p: p}
}

// NewLineDashedMaterial returns a new LineDashedMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     linewidth — Line thickness. Default is 1.
//     scale — The scale of the dashed part of a line. Default is 1.
//     dashSize — The size of the dash. Default is 3.
//     gapSize - The size of the gap. Default is 1.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
func (t *Three) NewLineDashedMaterial(parameters map[string]interface{}) *LineDashedMaterial {
	p := t.ctx.Get("LineDashedMaterial").New(parameters)
	return LineDashedMaterialFromJSObject(p)
}

// Copy TODO description.
func (l *LineDashedMaterial) Copy(source *LineDashedMaterial) *LineDashedMaterial {
	l.p.Call("copy", source.p)
	return l
}
