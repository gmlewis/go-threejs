package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Math represents a math.
type Math struct{ p *js.Object }

// Math returns a math object.
func (t *Three) Math() *Math {
	p := t.ctx.Get("Math")
	return &Math{p: p}
}

// GenerateUUID TODO description.
func (m *Math) GenerateUUID() *Math {
	m.p.Call("generateUUID")
	return m
}

// Clamp TODO description.
func (m *Math) Clamp(value, min, max float64) *Math {
	m.p.Call("clamp", value, min, max)
	return m
}

// EuclideanModulo TODO description.
func (m *Math) EuclideanModulo(n, m float64) *Math {
	m.p.Call("euclideanModulo", n, m)
	return m
}

// MapLinear TODO description.
func (m *Math) MapLinear(x, a1, a2, b1, b2 float64) *Math {
	m.p.Call("mapLinear", x, a1, a2, b1, b2)
	return m
}

// Smoothstep TODO description.
func (m *Math) Smoothstep(x, min, max float64) *Math {
	m.p.Call("smoothstep", x, min, max)
	return m
}

// Smootherstep TODO description.
func (m *Math) Smootherstep(x, min, max float64) *Math {
	m.p.Call("smootherstep", x, min, max)
	return m
}

// Random16 TODO description.
func (m *Math) Random16() *Math {
	m.p.Call("random16")
	return m
}

// RandInt TODO description.
func (m *Math) RandInt(low, high float64) *Math {
	m.p.Call("randInt", low, high)
	return m
}

// RandFloat TODO description.
func (m *Math) RandFloat(low, high float64) *Math {
	m.p.Call("randFloat", low, high)
	return m
}

// RandFloatSpread TODO description.
func (m *Math) RandFloatSpread(rnge float64) *Math {
	m.p.Call("randFloatSpread", rnge)
	return m
}

// DegToRad TODO description.
func (m *Math) DegToRad() *Math {
	m.p.Call("degToRad")
	return m
}

// RadToDeg TODO description.
func (m *Math) RadToDeg() *Math {
	m.p.Call("radToDeg")
	return m
}

// IsPowerOfTwo TODO description.
func (m *Math) IsPowerOfTwo(value float64) *Math {
	m.p.Call("isPowerOfTwo", value)
	return m
}

// NearestPowerOfTwo TODO description.
func (m *Math) NearestPowerOfTwo(value float64) *Math {
	m.p.Call("nearestPowerOfTwo", value)
	return m
}

// NextPowerOfTwo TODO description.
func (m *Math) NextPowerOfTwo(value float64) *Math {
	m.p.Call("nextPowerOfTwo", value)
	return m
}
