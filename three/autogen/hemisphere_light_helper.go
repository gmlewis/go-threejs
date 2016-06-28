package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// HemisphereLightHelper represents a hemispherelighthelper.
type HemisphereLightHelper struct{ p *js.Object }

// HemisphereLightHelper returns a hemispherelighthelper object.
func (t *Three) HemisphereLightHelper() *HemisphereLightHelper {
	p := t.ctx.Get("HemisphereLightHelper")
	return &HemisphereLightHelper{p: p}
}

// NewHemisphereLightHelper returns a new hemispherelighthelper object.
func (t *HemisphereLightHelper) New(light, sphereSize float64) *HemisphereLightHelper {
	p := t.p.New(light, sphereSize)
	return &HemisphereLightHelper{p: p}
}

