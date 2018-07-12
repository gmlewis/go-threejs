// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// WireframeGeometry represents a wireframegeometry.
type WireframeGeometry struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (w *WireframeGeometry) JSObject() *js.Object { return w.p }

// WireframeGeometry returns a WireframeGeometry JavaScript class.
func (t *Three) WireframeGeometry() *WireframeGeometry {
	p := t.ctx.Get("WireframeGeometry")
	return WireframeGeometryFromJSObject(p)
}

// WireframeGeometryFromJSObject returns a wrapped WireframeGeometry JavaScript class.
func WireframeGeometryFromJSObject(p *js.Object) *WireframeGeometry {
	return &WireframeGeometry{p: p}
}

// NewWireframeGeometry returns a new WireframeGeometry object.
func (t *Three) NewWireframeGeometry(geometry float64) *WireframeGeometry {
	p := t.ctx.Get("WireframeGeometry").New(geometry)
	return WireframeGeometryFromJSObject(p)
}
