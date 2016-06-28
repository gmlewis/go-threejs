package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SkeletonHelper represents a skeletonhelper.
type SkeletonHelper struct{ p *js.Object }

// SkeletonHelper returns a skeletonhelper object.
func (t *Three) SkeletonHelper() *SkeletonHelper {
	p := t.ctx.Get("SkeletonHelper")
	return &SkeletonHelper{p: p}
}

// NewSkeletonHelper returns a new skeletonhelper object.
func (t *SkeletonHelper) New(object float64) *SkeletonHelper {
	p := t.p.New(object)
	return &SkeletonHelper{p: p}
}

// GetBoneList TODO description.
func (s *SkeletonHelper) GetBoneList(object float64) *SkeletonHelper {
	s.p.Call("getBoneList", object)
	return s
}

