package three

import (
	"log"
	"testing"
)

func TestNewQuaternion(t *testing.T) {
	tt := New()
	q := tt.NewQuaternion()
	log.Printf("q=%+v", q)
}
