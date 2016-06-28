package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VectorKeyframeTrack represents a vectorkeyframetrack.
type VectorKeyframeTrack struct{ p *js.Object }

// VectorKeyframeTrack returns a vectorkeyframetrack object.
func (t *Three) VectorKeyframeTrack() *VectorKeyframeTrack {
	p := t.ctx.Get("VectorKeyframeTrack")
	return &VectorKeyframeTrack{p: p}
}

// NewVectorKeyframeTrack returns a new vectorkeyframetrack object.
func (t *VectorKeyframeTrack) New(name, times, values, interpolation float64) *VectorKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &VectorKeyframeTrack{p: p}
}

