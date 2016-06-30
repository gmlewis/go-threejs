// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineDashedMaterial represents a material for drawing wireframe-style geometries with dashed lines.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineDashedMaterial
type LineDashedMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LineDashedMaterial) JSObject() *js.Object { return t.p }

// LineDashedMaterial returns a LineDashedMaterial object.
func (t *Three) LineDashedMaterial() *LineDashedMaterial {
	p := t.ctx.Get("LineDashedMaterial")
	return &LineDashedMaterial{p: p}
}

// New returns a new LineDashedMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     linewidth — Line thickness. Default is 1.
//     scale — The scale of the dashed part of a line. Default is 1.
//     dashSize — The size of the dash. Default is 3.
//     gapSize - The size of the gap. Default is 1.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
func (t *LineDashedMaterial) New(parameters map[string]interface{}) *LineDashedMaterial {
	p := t.p.New(parameters)
	return &LineDashedMaterial{p: p}
}

// Copy TODO description.
func (l *LineDashedMaterial) Copy(source *LineDashedMaterial) *LineDashedMaterial {
	l.p.Call("copy", source.p)
	return l
}