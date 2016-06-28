package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BooleanKeyframeTrack represents a booleankeyframetrack.
type BooleanKeyframeTrack struct{ p *js.Object }

// BooleanKeyframeTrack returns a booleankeyframetrack object.
func (t *Three) BooleanKeyframeTrack() *BooleanKeyframeTrack {
	p := t.ctx.Get("BooleanKeyframeTrack")
	return &BooleanKeyframeTrack{p: p}
}

// NewBooleanKeyframeTrack returns a new booleankeyframetrack object.
func (t *BooleanKeyframeTrack) New(name, times, values float64) *BooleanKeyframeTrack {
	p := t.p.New(name, times, values)
	return &BooleanKeyframeTrack{p: p}
}

