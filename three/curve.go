package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Curve represents a curve.
type Curve struct{ p *js.Object }

// Curve returns a Curve object.
func (t *Three) Curve() *Curve {
	p := t.ctx.Get("Curve")
	return &Curve{p: p}
}

// New returns a new Curve object.
func (t *Curve) New() *Curve {
	p := t.p.New()
	return &Curve{p: p}
}

// GetPoint TODO description.
func (c *Curve) GetPoint(t float64) *Curve {
	c.p.Call("getPoint", t)
	return c
}

// GetPointAt TODO description.
func (c *Curve) GetPointAt(u float64) *Curve {
	c.p.Call("getPointAt", u)
	return c
}

// GetPoints TODO description.
func (c *Curve) GetPoints(divisions float64) *Curve {
	c.p.Call("getPoints", divisions)
	return c
}

// GetSpacedPoints TODO description.
func (c *Curve) GetSpacedPoints(divisions float64) *Curve {
	c.p.Call("getSpacedPoints", divisions)
	return c
}

// GetLength TODO description.
func (c *Curve) GetLength() *Curve {
	c.p.Call("getLength")
	return c
}

// GetLengths TODO description.
func (c *Curve) GetLengths(divisions float64) *Curve {
	c.p.Call("getLengths", divisions)
	return c
}

// UpdateArcLengths TODO description.
func (c *Curve) UpdateArcLengths() *Curve {
	c.p.Call("updateArcLengths")
	return c
}

// GetUtoTmapping TODO description.
func (c *Curve) GetUtoTmapping(u, distance float64) *Curve {
	c.p.Call("getUtoTmapping", u, distance)
	return c
}

// GetTangent TODO description.
func (c *Curve) GetTangent(t float64) *Curve {
	c.p.Call("getTangent", t)
	return c
}

// GetTangentAt TODO description.
func (c *Curve) GetTangentAt(u float64) *Curve {
	c.p.Call("getTangentAt", u)
	return c
}

