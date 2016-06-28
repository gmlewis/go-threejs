package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// StringKeyframeTrack represents a stringkeyframetrack.
type StringKeyframeTrack struct{ p *js.Object }

// StringKeyframeTrack returns a stringkeyframetrack object.
func (t *Three) StringKeyframeTrack() *StringKeyframeTrack {
	p := t.ctx.Get("StringKeyframeTrack")
	return &StringKeyframeTrack{p: p}
}

// NewStringKeyframeTrack returns a new stringkeyframetrack object.
func (t *StringKeyframeTrack) New(name, times, values, interpolation float64) *StringKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &StringKeyframeTrack{p: p}
}

