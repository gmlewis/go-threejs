package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Group represents a group.
type Group struct{ p *js.Object }

// Group returns a Group object.
func (t *Three) Group() *Group {
	p := t.ctx.Get("Group")
	return &Group{p: p}
}

// New returns a new Group object.
func (t *Group) New() *Group {
	p := t.p.New()
	return &Group{p: p}
}

