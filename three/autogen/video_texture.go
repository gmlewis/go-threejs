package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// VideoTexture represents a videotexture.
type VideoTexture struct{ p *js.Object }

// VideoTexture returns a VideoTexture object.
func (t *Three) VideoTexture() *VideoTexture {
	p := t.ctx.Get("VideoTexture")
	return &VideoTexture{p: p}
}

// New returns a new VideoTexture object.
func (t *VideoTexture) New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *VideoTexture {
	p := t.p.New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &VideoTexture{p: p}
}

