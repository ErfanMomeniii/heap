package heap

type node struct {
	value int
	alias any
}

type tree struct {
	nodes []*node
}

type MinHeap struct {
	tree *tree
}

type MaxHeap struct {
	tree *tree
}

func parent(index int) int {
	if index == 0 {
		return -1
	}

	return (index - 1) / 2
}

func leftChild(tree *tree, index int) (*node, int) {
	l := (2 * index) + 1
	if l >= len(tree.nodes) {
		return nil, -1
	}

	return tree.nodes[l], l
}

func rightChild(tree *tree, index int) (*node, int) {
	r := (2 * index) + 2
	if r >= len(tree.nodes) {
		return nil, -1
	}

	return tree.nodes[r], r
}

func swap(tree *tree, index1 int, index2 int) {
	tree.nodes[index1], tree.nodes[index2] = tree.nodes[index2], tree.nodes[index1]
	return
}

func (h *MinHeap) Insert(value int, alias ...any) {
	var a any
	if len(alias) > 0 {
		a = alias[0]
	}

	n := &node{
		value: value,

		alias: a,
	}

	h.tree.nodes = append(h.tree.nodes, n)
	index := len(h.tree.nodes) - 1
	for {
		p := parent(index)
		if p == -1 {
			break
		}

		if n.value < h.tree.nodes[p].value {
			swap(h.tree, index, p)
			index = p
		} else {
			break
		}
	}
}

func (h *MaxHeap) Insert(value int, alias ...any) {
	var a any
	if len(alias) > 0 {
		a = alias[0]
	}

	n := &node{
		value: value,
		alias: a,
	}

	h.tree.nodes = append(h.tree.nodes, n)
	index := len(h.tree.nodes) - 1
	for {
		p := parent(index)
		if p == -1 {
			break
		}

		if n.value > h.tree.nodes[p].value {
			swap(h.tree, index, p)
			index = p
		} else {
			break
		}
	}
}

func (h *MinHeap) Delete() {
	if len(h.tree.nodes) == 0 {
		return
	}

	if len(h.tree.nodes) == 1 {
		h.tree.nodes = []*node{}

		return
	}

	swap(h.tree, 0, len(h.tree.nodes)-1)
	h.tree.nodes = h.tree.nodes[:len(h.tree.nodes)-1]

	h.heapify(0)
}

func (h *MaxHeap) Delete() {
	if len(h.tree.nodes) == 0 {
		return
	}

	if len(h.tree.nodes) == 1 {
		h.tree.nodes = []*node{}

		return
	}

	swap(h.tree, 0, len(h.tree.nodes)-1)
	h.tree.nodes = h.tree.nodes[:len(h.tree.nodes)-1]

	h.heapify(0)
}

func (h *MinHeap) heapify(index int) {
	for {
		current := h.tree.nodes[index]

		left, leftIndex := leftChild(h.tree, index)
		right, rightIndex := rightChild(h.tree, index)
		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value < right.value {
				break
			} else {
				swap(h.tree, index, rightIndex)
				index = rightIndex
			}
		} else if right == nil && left != nil {
			if current.value < left.value {
				break
			} else {
				swap(h.tree, index, leftIndex)
				index = leftIndex
			}
		} else {
			if left.value >= current.value && right.value >= current.value {
				break
			} else if left.value < right.value && left.value < current.value {
				swap(h.tree, index, leftIndex)
				index = leftIndex
			} else {
				swap(h.tree, index, rightIndex)
				index = rightIndex
			}
		}
	}
}

func (h *MaxHeap) heapify(index int) {
	for {
		current := h.tree.nodes[index]

		left, leftIndex := leftChild(h.tree, index)
		right, rightIndex := rightChild(h.tree, index)

		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value > right.value {
				break
			} else {
				swap(h.tree, index, rightIndex)
				index = rightIndex
			}
		} else if right == nil && left != nil {
			if current.value > left.value {
				break
			} else {
				swap(h.tree, index, leftIndex)
				index = leftIndex
			}
		} else {
			if left.value <= current.value && right.value <= current.value {
				break
			} else if left.value > right.value && left.value > current.value {
				swap(h.tree, index, leftIndex)
				index = leftIndex
			} else {
				swap(h.tree, index, rightIndex)
				index = rightIndex
			}
		}
	}
}
func (h *MinHeap) GetMin() int {
	if len(h.tree.nodes) == 0 {
		return 0
	}

	return h.tree.nodes[0].value
}

func (h *MaxHeap) GetMax() int {
	if len(h.tree.nodes) == 0 {
		return 0
	}

	return h.tree.nodes[0].value
}

func NewMin() *MinHeap {
	return &MinHeap{
		tree: &tree{
			nodes: make([]*node, 0),
		},
	}
}

func NewMax() *MaxHeap {
	return &MaxHeap{
		tree: &tree{
			nodes: make([]*node, 0),
		},
	}
}
