package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CatmullRomCurve3 represents a catmullromcurve3.
type CatmullRomCurve3 struct{ p *js.Object }

// CatmullRomCurve3 returns a CatmullRomCurve3 object.
func (t *Three) CatmullRomCurve3() *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3")
	return &CatmullRomCurve3{p: p}
}

// New returns a new CatmullRomCurve3 object.
func (t *CatmullRomCurve3) New() *CatmullRomCurve3 {
	p := t.p.New()
	return &CatmullRomCurve3{p: p}
}

