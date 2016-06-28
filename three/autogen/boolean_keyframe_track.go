package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BooleanKeyframeTrack represents a booleankeyframetrack.
type BooleanKeyframeTrack struct{ p *js.Object }

// BooleanKeyframeTrack returns a BooleanKeyframeTrack object.
func (t *Three) BooleanKeyframeTrack() *BooleanKeyframeTrack {
	p := t.ctx.Get("BooleanKeyframeTrack")
	return &BooleanKeyframeTrack{p: p}
}

// New returns a new BooleanKeyframeTrack object.
func (t *BooleanKeyframeTrack) New(name, times, values float64) *BooleanKeyframeTrack {
	p := t.p.New(name, times, values)
	return &BooleanKeyframeTrack{p: p}
}

