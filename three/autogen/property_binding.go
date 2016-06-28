package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PropertyBinding represents a propertybinding.
type PropertyBinding struct{ p *js.Object }

// PropertyBinding returns a propertybinding object.
func (t *Three) PropertyBinding() *PropertyBinding {
	p := t.ctx.Get("PropertyBinding")
	return &PropertyBinding{p: p}
}

// NewPropertyBinding returns a new propertybinding object.
func (t *PropertyBinding) New(rootNode, path, parsedPath float64) *PropertyBinding {
	p := t.p.New(rootNode, path, parsedPath)
	return &PropertyBinding{p: p}
}

// Bind TODO description.
func (p *PropertyBinding) Bind() *PropertyBinding {
	p.p.Call("bind")
	return p
}

// Unbind TODO description.
func (p *PropertyBinding) Unbind() *PropertyBinding {
	p.p.Call("unbind")
	return p
}

// _getValue_unavailable TODO description.
func (p *PropertyBinding) _getValue_unavailable() *PropertyBinding {
	p.p.Call("_getValue_unavailable")
	return p
}

// _setValue_unavailable TODO description.
func (p *PropertyBinding) _setValue_unavailable() *PropertyBinding {
	p.p.Call("_setValue_unavailable")
	return p
}

// GetValue TODO description.
func (p *PropertyBinding) GetValue(array, offset float64) *PropertyBinding {
	p.p.Call("getValue", array, offset)
	return p
}

// SetValue TODO description.
func (p *PropertyBinding) SetValue(array, offset float64) *PropertyBinding {
	p.p.Call("setValue", array, offset)
	return p
}
