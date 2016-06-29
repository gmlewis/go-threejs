package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArrowHelper represents an arrowhelper.
type ArrowHelper struct{ p *js.Object }

// ArrowHelper returns an ArrowHelper object.
func (t *Three) ArrowHelper() *ArrowHelper {
	p := t.ctx.Get("ArrowHelper")
	return &ArrowHelper{p: p}
}

// New returns a new ArrowHelper object.
func (t *ArrowHelper) New() *ArrowHelper {
	p := t.p.New()
	return &ArrowHelper{p: p}
}

// SetLength TODO description.
func (a *ArrowHelper) SetLength(length, headLength, headWidth float64) *ArrowHelper {
	a.p.Call("setLength", length, headLength, headWidth)
	return a
}

// SetColor TODO description.
func (a *ArrowHelper) SetColor(color float64) *ArrowHelper {
	a.p.Call("setColor", color)
	return a
}

