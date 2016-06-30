// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineBasicMaterial represents a material for drawing wireframe-style geometries.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineBasicMaterial
type LineBasicMaterial struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LineBasicMaterial) JSObject() *js.Object { return l.p }

// LineBasicMaterial returns a LineBasicMaterial object.
func (t *Three) LineBasicMaterial() *LineBasicMaterial {
	p := t.ctx.Get("LineBasicMaterial")
	return &LineBasicMaterial{p: p}
}

// New returns a new LineBasicMaterial object.
//
// parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     linewidth — Line thickness. Default is 1.
//     linecap — Define appearance of line ends. Default is 'round'.
//     linejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
func (l *LineBasicMaterial) New(parameters map[string]interface{}) *LineBasicMaterial {
	p := l.p.New(parameters)
	return &LineBasicMaterial{p: p}
}

// Copy TODO description.
func (l *LineBasicMaterial) Copy(source float64) *LineBasicMaterial {
	l.p.Call("copy", source)
	return l
}
