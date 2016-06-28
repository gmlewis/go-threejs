package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AudioAnalyser represents an audioanalyser.
type AudioAnalyser struct{ p *js.Object }

// AudioAnalyser returns an audioanalyser object.
func (t *Three) AudioAnalyser() *AudioAnalyser {
	p := t.ctx.Get("AudioAnalyser")
	return &AudioAnalyser{p: p}
}

// NewAudioAnalyser returns a new audioanalyser object.
func (t *AudioAnalyser) New(audio, fftSize float64) *AudioAnalyser {
	p := t.p.New(audio, fftSize)
	return &AudioAnalyser{p: p}
}

// GetData TODO description.
func (a *AudioAnalyser) GetData() *AudioAnalyser {
	a.p.Call("getData")
	return a
}

