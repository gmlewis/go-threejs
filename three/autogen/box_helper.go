package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxHelper represents a boxhelper.
type BoxHelper struct{ p *js.Object }

// BoxHelper returns a boxhelper object.
func (t *Three) BoxHelper() *BoxHelper {
	p := t.ctx.Get("BoxHelper")
	return &BoxHelper{p: p}
}

// NewBoxHelper returns a new boxhelper object.
func (t *BoxHelper) New(object float64) *BoxHelper {
	p := t.p.New(object)
	return &BoxHelper{p: p}
}

