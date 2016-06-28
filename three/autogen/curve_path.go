package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CurvePath represents a curvepath.
type CurvePath struct{ p *js.Object }

// CurvePath returns a CurvePath object.
func (t *Three) CurvePath() *CurvePath {
	p := t.ctx.Get("CurvePath")
	return &CurvePath{p: p}
}

// New returns a new CurvePath object.
func (t *CurvePath) New() *CurvePath {
	p := t.p.New()
	return &CurvePath{p: p}
}

// Add TODO description.
func (c *CurvePath) Add(curve float64) *CurvePath {
	c.p.Call("add", curve)
	return c
}

// GetPoint TODO description.
func (c *CurvePath) GetPoint(t float64) *CurvePath {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CurvePath) GetTangent(t float64) *CurvePath {
	c.p.Call("getTangent", t)
	return c
}

// CreatePointsGeometry TODO description.
func (c *CurvePath) CreatePointsGeometry(divisions float64) *CurvePath {
	c.p.Call("createPointsGeometry", divisions)
	return c
}

// CreateSpacedPointsGeometry TODO description.
func (c *CurvePath) CreateSpacedPointsGeometry(divisions float64) *CurvePath {
	c.p.Call("createSpacedPointsGeometry", divisions)
	return c
}

// CreateGeometry TODO description.
func (c *CurvePath) CreateGeometry(points float64) *CurvePath {
	c.p.Call("createGeometry", points)
	return c
}

