package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectionalLightHelper represents a directionallighthelper.
type DirectionalLightHelper struct{ p *js.Object }

// DirectionalLightHelper returns a DirectionalLightHelper object.
func (t *Three) DirectionalLightHelper() *DirectionalLightHelper {
	p := t.ctx.Get("DirectionalLightHelper")
	return &DirectionalLightHelper{p: p}
}

// New returns a new DirectionalLightHelper object.
func (t *DirectionalLightHelper) New(light, size float64) *DirectionalLightHelper {
	p := t.p.New(light, size)
	return &DirectionalLightHelper{p: p}
}

