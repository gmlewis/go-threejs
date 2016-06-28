package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineBasicMaterial represents a linebasicmaterial.
type LineBasicMaterial struct{ p *js.Object }

// LineBasicMaterial returns a linebasicmaterial object.
func (t *Three) LineBasicMaterial() *LineBasicMaterial {
	p := t.ctx.Get("LineBasicMaterial")
	return &LineBasicMaterial{p: p}
}

// NewLineBasicMaterial returns a new linebasicmaterial object.
func (t *LineBasicMaterial) New(parameters float64) *LineBasicMaterial {
	p := t.p.New(parameters)
	return &LineBasicMaterial{p: p}
}

// Copy TODO description.
func (l *LineBasicMaterial) Copy(source float64) *LineBasicMaterial {
	l.p.Call("copy", source)
	return l
}

