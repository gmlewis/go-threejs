package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EdgesHelper represents an edgeshelper.
type EdgesHelper struct{ p *js.Object }

// EdgesHelper returns an EdgesHelper object.
func (t *Three) EdgesHelper() *EdgesHelper {
	p := t.ctx.Get("EdgesHelper")
	return &EdgesHelper{p: p}
}

// New returns a new EdgesHelper object.
func (t *EdgesHelper) New(object, hex, thresholdAngle float64) *EdgesHelper {
	p := t.p.New(object, hex, thresholdAngle)
	return &EdgesHelper{p: p}
}

