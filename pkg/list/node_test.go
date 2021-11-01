package list

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	n := newNode[int](10)

	if n.valid() != true {
		t.Errorf("node must be valid")
	}
}

func TestNewEndNode(t *testing.T) {
	n := newEndNode[int]()

	if n.valid() != false {
		t.Errorf("end node must be invalid")
	}
}
