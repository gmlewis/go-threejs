package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LOD represents a lod.
type LOD struct{ p *js.Object }

// LOD returns a LOD object.
func (t *Three) LOD() *LOD {
	p := t.ctx.Get("LOD")
	return &LOD{p: p}
}

// New returns a new LOD object.
func (t *LOD) New() *LOD {
	p := t.p.New()
	return &LOD{p: p}
}

// Get TODO description.
func (l *LOD) Get() *LOD {
	l.p.Call("get")
	return l
}

// AddLevel TODO description.
func (l *LOD) AddLevel(object, distance float64) *LOD {
	l.p.Call("addLevel", object, distance)
	return l
}

// GetObjectForDistance TODO description.
func (l *LOD) GetObjectForDistance(distance float64) *LOD {
	l.p.Call("getObjectForDistance", distance)
	return l
}

// Copy TODO description.
func (l *LOD) Copy(source float64) *LOD {
	l.p.Call("copy", source)
	return l
}

// ToJSON TODO description.
func (l *LOD) ToJSON(meta float64) *LOD {
	l.p.Call("toJSON", meta)
	return l
}

