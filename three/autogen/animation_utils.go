package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AnimationUtils represents an animationutils.
type AnimationUtils struct{ p *js.Object }

// AnimationUtils returns an AnimationUtils object.
func (t *Three) AnimationUtils() *AnimationUtils {
	p := t.ctx.Get("AnimationUtils")
	return &AnimationUtils{p: p}
}

// New returns a new AnimationUtils object.
func (t *AnimationUtils) New() *AnimationUtils {
	p := t.p.New()
	return &AnimationUtils{p: p}
}

// ArraySlice TODO description.
func (a *AnimationUtils) ArraySlice(array, from, to float64) *AnimationUtils {
	a.p.Call("arraySlice", array, from, to)
	return a
}

// ConvertArray TODO description.
func (a *AnimationUtils) ConvertArray(array, typ, forceClone float64) *AnimationUtils {
	a.p.Call("convertArray", array, typ, forceClone)
	return a
}

// IsTypedArray TODO description.
func (a *AnimationUtils) IsTypedArray(object float64) *AnimationUtils {
	a.p.Call("isTypedArray", object)
	return a
}

// GetKeyframeOrder TODO description.
func (a *AnimationUtils) GetKeyframeOrder(times float64) *AnimationUtils {
	a.p.Call("getKeyframeOrder", times)
	return a
}

// SortedArray TODO description.
func (a *AnimationUtils) SortedArray(values, stride, order float64) *AnimationUtils {
	a.p.Call("sortedArray", values, stride, order)
	return a
}

// FlattenJSON TODO description.
func (a *AnimationUtils) FlattenJSON(jsonKeys, times, values, valuePropertyName float64) *AnimationUtils {
	a.p.Call("flattenJSON", jsonKeys, times, values, valuePropertyName)
	return a
}

