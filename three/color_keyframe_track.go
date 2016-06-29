package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ColorKeyframeTrack represents a colorkeyframetrack.
type ColorKeyframeTrack struct{ p *js.Object }

// ColorKeyframeTrack returns a ColorKeyframeTrack object.
func (t *Three) ColorKeyframeTrack() *ColorKeyframeTrack {
	p := t.ctx.Get("ColorKeyframeTrack")
	return &ColorKeyframeTrack{p: p}
}

// New returns a new ColorKeyframeTrack object.
func (t *ColorKeyframeTrack) New(name, times, values, interpolation float64) *ColorKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &ColorKeyframeTrack{p: p}
}

