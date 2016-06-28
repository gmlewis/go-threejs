package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// XHRLoader represents a xhrloader.
type XHRLoader struct{ p *js.Object }

// XHRLoader returns a xhrloader object.
func (t *Three) XHRLoader() *XHRLoader {
	p := t.ctx.Get("XHRLoader")
	return &XHRLoader{p: p}
}

// NewXHRLoader returns a new xhrloader object.
func (t *XHRLoader) New(manager float64) *XHRLoader {
	p := t.p.New(manager)
	return &XHRLoader{p: p}
}

// Load TODO description.
func (x *XHRLoader) Load(url, onLoad, onProgress, onError float64) *XHRLoader {
	x.p.Call("load", url, onLoad, onProgress, onError)
	return x
}

// SetPath TODO description.
func (x *XHRLoader) SetPath(value float64) *XHRLoader {
	x.p.Call("setPath", value)
	return x
}

// SetResponseType TODO description.
func (x *XHRLoader) SetResponseType(value float64) *XHRLoader {
	x.p.Call("setResponseType", value)
	return x
}

// SetWithCredentials TODO description.
func (x *XHRLoader) SetWithCredentials(value float64) *XHRLoader {
	x.p.Call("setWithCredentials", value)
	return x
}

