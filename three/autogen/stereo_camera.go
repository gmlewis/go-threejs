package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// StereoCamera represents a stereocamera.
type StereoCamera struct{ p *js.Object }

// StereoCamera returns a StereoCamera object.
func (t *Three) StereoCamera() *StereoCamera {
	p := t.ctx.Get("StereoCamera")
	return &StereoCamera{p: p}
}

// New returns a new StereoCamera object.
func (t *StereoCamera) New() *StereoCamera {
	p := t.p.New()
	return &StereoCamera{p: p}
}

