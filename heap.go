package heap

type node struct {
	value int
	index int
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

func parent(tree *tree, index int) *node {
	if index == 0 {
		return nil
	}

	return tree.nodes[index-1/2]
}

func leftChild(tree *tree, index int) *node {
	l := (2 * index) + 1
	if l > len(tree.nodes) {
		return nil
	}

	return tree.nodes[l]
}

func rightChild(tree *tree, index int) *node {
	r := (2 * index) + 2
	if r > len(tree.nodes) {
		return nil
	}

	return tree.nodes[r]
}

func swap(n1 *node, n2 *node) {
	v := n1.value
	n1.value = n2.value
	n2.value = v
}

func (h *MinHeap) Insert(value int) {
	n := &node{
		value: value,
		index: len(h.tree.nodes),
	}

	h.tree.nodes = append(h.tree.nodes, n)

	for {
		p := parent(h.tree, n.index)
		if p == nil {
			break
		}

		if n.value < p.value {
			swap(n, p)
		} else {
			break
		}
	}
}

func (h *MaxHeap) Insert(value int) {
	n := &node{
		value: value,
		index: len(h.tree.nodes),
	}

	h.tree.nodes = append(h.tree.nodes, n)

	for {
		p := parent(h.tree, n.index)
		if p == n {
			break
		}

		if n.value > p.value {
			swap(n, p)
		} else {
			break
		}
	}
}

func (h *MinHeap) Delete() {
	swap(h.tree.nodes[0], h.tree.nodes[len(h.tree.nodes)-1])
	h.tree.nodes = h.tree.nodes[:len(h.tree.nodes)-1]

	h.heapify(0)
}

func (h *MaxHeap) Delete() {
	swap(h.tree.nodes[0], h.tree.nodes[len(h.tree.nodes)-1])
	h.tree.nodes = h.tree.nodes[:len(h.tree.nodes)-1]

	h.heapify(0)
}

func (h *MinHeap) heapify(index int) {
	for {
		current := h.tree.nodes[index]

		left := leftChild(h.tree, index)
		right := rightChild(h.tree, index)

		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value < right.value {
				break
			} else {
				swap(current, right)
				index = right.index
			}
		} else if right == nil && left != nil {
			if current.value < left.value {
				break
			} else {
				swap(current, left)
				index = left.index
			}
		} else {
			if left.value >= current.value && right.value >= current.value {
				break
			} else if left.value < right.value && left.value < current.value {
				swap(current, left)
				index = left.index
			} else {
				swap(current, right)
				index = right.index
			}
		}
	}
}

func (h *MaxHeap) heapify(index int) {
	for {
		current := h.tree.nodes[index]

		left := leftChild(h.tree, index)
		right := rightChild(h.tree, index)

		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value > right.value {
				break
			} else {
				swap(current, right)
				index = right.index
			}
		} else if right == nil && left != nil {
			if current.value > left.value {
				break
			} else {
				swap(current, left)
				index = left.index
			}
		} else {
			if left.value <= current.value && right.value <= current.value {
				break
			} else if left.value > right.value && left.value > current.value {
				swap(current, left)
				index = left.index
			} else {
				swap(current, right)
				index = right.index
			}
		}
	}
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
