package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// NumberKeyframeTrack represents a numberkeyframetrack.
type NumberKeyframeTrack struct{ p *js.Object }

// NumberKeyframeTrack returns a numberkeyframetrack object.
func (t *Three) NumberKeyframeTrack() *NumberKeyframeTrack {
	p := t.ctx.Get("NumberKeyframeTrack")
	return &NumberKeyframeTrack{p: p}
}

// NewNumberKeyframeTrack returns a new numberkeyframetrack object.
func (t *NumberKeyframeTrack) New(name, times, values, interpolation float64) *NumberKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &NumberKeyframeTrack{p: p}
}

