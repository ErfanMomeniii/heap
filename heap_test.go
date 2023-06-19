package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxHeap(t *testing.T) {
	h := NewMax()

	h.Insert(4)
	h.Insert(12)
	h.Insert(10)
	assert.Equal(t, h.GetMax(), 12)

	h.Delete()
	assert.Equal(t, h.GetMax(), 10)

	h.Delete()
	assert.Equal(t, h.GetMax(), 4)

	h.Delete()
	assert.Equal(t, h.GetMax(), 0)
}

func Test_MinHeap(t *testing.T) {
	h := NewMin()

	h.Insert(4)
	h.Insert(12)
	h.Insert(10)
	assert.Equal(t, h.GetMin(), 4)

	h.Delete()
	assert.Equal(t, h.GetMin(), 10)

	h.Delete()
	assert.Equal(t, h.GetMin(), 12)

	h.Delete()
	assert.Equal(t, h.GetMin(), 0)
}
