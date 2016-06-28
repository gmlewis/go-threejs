package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// QuaternionKeyframeTrack represents a quaternionkeyframetrack.
type QuaternionKeyframeTrack struct{ p *js.Object }

// QuaternionKeyframeTrack returns a QuaternionKeyframeTrack object.
func (t *Three) QuaternionKeyframeTrack() *QuaternionKeyframeTrack {
	p := t.ctx.Get("QuaternionKeyframeTrack")
	return &QuaternionKeyframeTrack{p: p}
}

// New returns a new QuaternionKeyframeTrack object.
func (t *QuaternionKeyframeTrack) New(name, times, values, interpolation float64) *QuaternionKeyframeTrack {
	p := t.p.New(name, times, values, interpolation)
	return &QuaternionKeyframeTrack{p: p}
}

// InterpolantFactoryMethodLinear TODO description.
func (q *QuaternionKeyframeTrack) InterpolantFactoryMethodLinear(result float64) *QuaternionKeyframeTrack {
	q.p.Call("InterpolantFactoryMethodLinear", result)
	return q
}

