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

// Group represents a group of objects.
type Group struct{ *Object3D }

// JSObject returns the underlying *js.Object.
func (g *Group) JSObject() *js.Object { return g.p }

// Group returns a Group JavaScript class.
func (t *Three) Group() *Group {
	p := t.ctx.Get("Group")
	return GroupFromJSObject(p)
}

// GroupFromJSObject returns a wrapped Group JavaScript class.
func GroupFromJSObject(p *js.Object) *Group {
	return &Group{Object3DFromJSObject(p)}
}

// NewGroup returns a new Group object.
func (t *Three) NewGroup() *Group {
	p := t.ctx.Get("Group").New()
	return GroupFromJSObject(p)
}
