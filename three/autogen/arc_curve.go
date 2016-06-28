package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArcCurve represents an arccurve.
type ArcCurve struct{ p *js.Object }

// ArcCurve returns an arccurve object.
func (t *Three) ArcCurve() *ArcCurve {
	p := t.ctx.Get("ArcCurve")
	return &ArcCurve{p: p}
}

// NewArcCurve returns a new arccurve object.
func (t *ArcCurve) New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise float64) *ArcCurve {
	p := t.p.New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
	return &ArcCurve{p: p}
}

