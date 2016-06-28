package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ImageLoader represents an imageloader.
type ImageLoader struct{ p *js.Object }

// ImageLoader returns an ImageLoader object.
func (t *Three) ImageLoader() *ImageLoader {
	p := t.ctx.Get("ImageLoader")
	return &ImageLoader{p: p}
}

// New returns a new ImageLoader object.
func (t *ImageLoader) New(manager float64) *ImageLoader {
	p := t.p.New(manager)
	return &ImageLoader{p: p}
}

// Load TODO description.
func (i *ImageLoader) Load(url, onLoad, onProgress, onError float64) *ImageLoader {
	i.p.Call("load", url, onLoad, onProgress, onError)
	return i
}

// SetCrossOrigin TODO description.
func (i *ImageLoader) SetCrossOrigin(value float64) *ImageLoader {
	i.p.Call("setCrossOrigin", value)
	return i
}

// SetPath TODO description.
func (i *ImageLoader) SetPath(value float64) *ImageLoader {
	i.p.Call("setPath", value)
	return i
}

