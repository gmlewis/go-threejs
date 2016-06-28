package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PointLightHelper represents a pointlighthelper.
type PointLightHelper struct{ p *js.Object }

// PointLightHelper returns a pointlighthelper object.
func (t *Three) PointLightHelper() *PointLightHelper {
	p := t.ctx.Get("PointLightHelper")
	return &PointLightHelper{p: p}
}

// NewPointLightHelper returns a new pointlighthelper object.
func (t *PointLightHelper) New(light, sphereSize float64) *PointLightHelper {
	p := t.p.New(light, sphereSize)
	return &PointLightHelper{p: p}
}

