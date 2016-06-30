// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LoadingManager represents a loadingmanager.
type LoadingManager struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (l *LoadingManager) JSObject() *js.Object { return l.p }

// LoadingManager returns a LoadingManager object.
func (t *Three) LoadingManager() *LoadingManager {
	p := t.ctx.Get("LoadingManager")
	return &LoadingManager{p: p}
}

// New returns a new LoadingManager object.
func (l *LoadingManager) New(onLoad, onProgress, onError float64) *LoadingManager {
	p := l.p.New(onLoad, onProgress, onError)
	return &LoadingManager{p: p}
}
