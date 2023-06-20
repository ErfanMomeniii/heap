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
	m, _ := h.GetMax()
	assert.Equal(t, m, 12)

	h.Delete()
	m, _ = h.GetMax()
	assert.Equal(t, m, 10)

	h.Delete()
	m, _ = h.GetMax()
	assert.Equal(t, m, 4)

	h.Delete()
	m, _ = h.GetMax()
	assert.Equal(t, m, 0)
}

func Test_MinHeap(t *testing.T) {
	h := NewMin()

	h.Insert(4)
	h.Insert(12)
	h.Insert(10)

	m, _ := h.GetMin()
	assert.Equal(t, m, 4)

	h.Delete()
	m, _ = h.GetMin()
	assert.Equal(t, m, 10)

	h.Delete()
	m, _ = h.GetMin()
	assert.Equal(t, m, 12)

	h.Delete()
	m, _ = h.GetMin()
	assert.Equal(t, m, 0)
}

func Test_Update_MinHeap(t *testing.T) {
	h := NewMin()

	h.Insert(4, 1)
	h.Insert(12, 2)
	h.Insert(10, 3)

	m, _ := h.GetMin()
	assert.Equal(t, m, 4)

	h.Update(1, 13)
	m, _ = h.GetMin()
	assert.Equal(t, m, 10)

	h.Update(1, 2)
	m, _ = h.GetMin()
	assert.Equal(t, m, 2)
}

func Test_Update_MaxHeap(t *testing.T) {
	h := NewMax()

	h.Insert(4, 1)
	h.Insert(12, 2)
	h.Insert(10, 3)

	m, _ := h.GetMax()
	assert.Equal(t, m, 12)

	h.Update(1, 13)
	m, _ = h.GetMax()
	assert.Equal(t, m, 13)

	h.Update(1, 2)
	m, _ = h.GetMax()
	assert.Equal(t, m, 12)
}
