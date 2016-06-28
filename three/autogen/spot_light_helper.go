package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SpotLightHelper represents a spotlighthelper.
type SpotLightHelper struct{ p *js.Object }

// SpotLightHelper returns a spotlighthelper object.
func (t *Three) SpotLightHelper() *SpotLightHelper {
	p := t.ctx.Get("SpotLightHelper")
	return &SpotLightHelper{p: p}
}

// NewSpotLightHelper returns a new spotlighthelper object.
func (t *SpotLightHelper) New(light float64) *SpotLightHelper {
	p := t.p.New(light)
	return &SpotLightHelper{p: p}
}

