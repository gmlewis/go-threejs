package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ColorKeyframeTrack represents a colorkeyframetrack.
type ColorKeyframeTrack struct{ p *js.Object }

// ColorKeyframeTrack returns a colorkeyframetrack object.
func (t *Three) ColorKeyframeTrack() *ColorKeyframeTrack {
	p := t.ctx.Get("ColorKeyframeTrack")
	return &ColorKeyframeTrack{p: p}
}

// NewColorKeyframeTrack returns a new colorkeyframetrack object.
func (t *ColorKeyframeTrack) New(name, times, values, interpolation float64) *ColorKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &ColorKeyframeTrack{p: p}
}

