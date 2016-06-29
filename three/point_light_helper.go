package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PointLightHelper represents a pointlighthelper.
type PointLightHelper struct{ p *js.Object }

// PointLightHelper returns a PointLightHelper object.
func (t *Three) PointLightHelper() *PointLightHelper {
	p := t.ctx.Get("PointLightHelper")
	return &PointLightHelper{p: p}
}

// New returns a new PointLightHelper object.
func (t *PointLightHelper) New(light, sphereSize float64) *PointLightHelper {
	p := t.p.New(light, sphereSize)
	return &PointLightHelper{p: p}
}

