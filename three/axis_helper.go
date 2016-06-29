package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AxisHelper represents an axishelper.
type AxisHelper struct{ p *js.Object }

// AxisHelper returns an AxisHelper object.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return &AxisHelper{p: p}
}

// New returns a new AxisHelper object.
func (t *AxisHelper) New(size float64) *AxisHelper {
	p := t.p.New(size)
	return &AxisHelper{p: p}
}

