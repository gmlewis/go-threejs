package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// GridHelper represents a gridhelper.
type GridHelper struct{ p *js.Object }

// GridHelper returns a gridhelper object.
func (t *Three) GridHelper() *GridHelper {
	p := t.ctx.Get("GridHelper")
	return &GridHelper{p: p}
}

// NewGridHelper returns a new gridhelper object.
func (t *GridHelper) New(size, step float64) *GridHelper {
	p := t.p.New(size, step)
	return &GridHelper{p: p}
}

// SetColors TODO description.
func (g *GridHelper) SetColors(colorCenterLine, colorGrid float64) *GridHelper {
	g.p.Call("setColors", colorCenterLine, colorGrid)
	return g
}

