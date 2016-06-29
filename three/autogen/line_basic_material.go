package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineBasicMaterial represents a material for drawing wireframe-style geometries.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineBasicMaterial
type LineBasicMaterial struct{ p *js.Object }

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
func (t *LineBasicMaterial) New(parameters map[string]interface{}) *LineBasicMaterial {
	p := t.p.New(parameters)
	return &LineBasicMaterial{p: p}
}

// Copy TODO description.
func (l *LineBasicMaterial) Copy(source float64) *LineBasicMaterial {
	l.p.Call("copy", source)
	return l
}
