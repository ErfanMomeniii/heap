package heap

// node is an instantiation of heap's elements that contains value and alias.
type node struct {
	value int
	alias any
}

// tree is an instantiation of a binary heap tree.
type tree struct {
	nodes []*node
}

// MinHeap is an instantiation of a min binary heap .
type MinHeap struct {
	tree *tree
}

// MaxHeap is an instantiation of a max binary heap .
type MaxHeap struct {
	tree *tree
}

// parent returns index of the parent node of inputted node.
func parent(nodeIndex int) int {
	if nodeIndex == 0 {
		return -1
	}

	return (nodeIndex - 1) / 2
}

// leftChild returns left child node and index of the inputted node.
func leftChild(tree *tree, nodeIndex int) (*node, int) {
	l := (2 * nodeIndex) + 1
	if l >= len(tree.nodes) {
		return nil, -1
	}

	return tree.nodes[l], l
}

// rightChild returns left child node and index of the inputted node.
func rightChild(tree *tree, nodeIndex int) (*node, int) {
	r := (2 * nodeIndex) + 2
	if r >= len(tree.nodes) {
		return nil, -1
	}

	return tree.nodes[r], r
}

// findNodeWithAlias returns node object and corresponding index of a node that matches with the inputted alias.
func findNodeWithAlias(tree *tree, alias any) (*node, int) {
	for i, n := range tree.nodes {
		if n.alias == alias {
			return n, i
		}
	}
	return nil, -1
}

// swap changes position of two node in the inputted tree.
func swap(tree *tree, index1 int, index2 int) {
	tree.nodes[index1], tree.nodes[index2] = tree.nodes[index2], tree.nodes[index1]
	return
}

// Insert add new node to the tree.
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

// Insert add new node to the tree.
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

// Delete removes root and updates (heapify) tree.
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

// Delete removes root and updates (heapify) tree.
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

// Update change value of the specific node that matches with alias and heapify tree.
func (h *MaxHeap) Update(alias any, value int) {
	n, i := findNodeWithAlias(h.tree, alias)
	if n == nil {
		h.Insert(value, alias)

		return
	}

	if value > n.value {
		n.value = value
		for {
			p := parent(i)
			if p == -1 {
				break
			}

			if n.value > h.tree.nodes[p].value {
				swap(h.tree, i, p)
				i = p
			} else {
				break
			}
		}
	} else {
		n.value = value
		h.heapify(i)
	}
}

// Update change value of the specific node that matches with alias and heapify tree.
func (h *MinHeap) Update(alias any, value int) {
	n, i := findNodeWithAlias(h.tree, alias)
	if n == nil {
		h.Insert(value, alias)

		return
	}

	if value < n.value {
		n.value = value
		for {
			p := parent(i)
			if p == -1 {
				break
			}

			if n.value < h.tree.nodes[p].value {
				swap(h.tree, i, p)
				i = p
			} else {
				break
			}
		}
	} else {
		n.value = value
		h.heapify(i)
	}
}

// heapify rearranged the tree's nodes to maintain the heap.
func (h *MinHeap) heapify(rootIndex int) {
	for {
		current := h.tree.nodes[rootIndex]

		left, leftIndex := leftChild(h.tree, rootIndex)
		right, rightIndex := rightChild(h.tree, rootIndex)
		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value < right.value {
				break
			} else {
				swap(h.tree, rootIndex, rightIndex)
				rootIndex = rightIndex
			}
		} else if right == nil && left != nil {
			if current.value < left.value {
				break
			} else {
				swap(h.tree, rootIndex, leftIndex)
				rootIndex = leftIndex
			}
		} else {
			if left.value >= current.value && right.value >= current.value {
				break
			} else if left.value < right.value && left.value < current.value {
				swap(h.tree, rootIndex, leftIndex)
				rootIndex = leftIndex
			} else {
				swap(h.tree, rootIndex, rightIndex)
				rootIndex = rightIndex
			}
		}
	}
}

// heapify rearranged the tree's nodes to maintain the heap.
func (h *MaxHeap) heapify(rootIndex int) {
	for {
		current := h.tree.nodes[rootIndex]

		left, leftIndex := leftChild(h.tree, rootIndex)
		right, rightIndex := rightChild(h.tree, rootIndex)

		if left == nil && right == nil {
			break
		} else if left == nil && right != nil {
			if current.value > right.value {
				break
			} else {
				swap(h.tree, rootIndex, rightIndex)
				rootIndex = rightIndex
			}
		} else if right == nil && left != nil {
			if current.value > left.value {
				break
			} else {
				swap(h.tree, rootIndex, leftIndex)
				rootIndex = leftIndex
			}
		} else {
			if left.value <= current.value && right.value <= current.value {
				break
			} else if left.value > right.value && left.value > current.value {
				swap(h.tree, rootIndex, leftIndex)
				rootIndex = leftIndex
			} else {
				swap(h.tree, rootIndex, rightIndex)
				rootIndex = rightIndex
			}
		}
	}
}

// GetMin returns min value (root) from heap tree with alias.
func (h *MinHeap) GetMin() (int, any) {
	if len(h.tree.nodes) == 0 {
		return 0, nil
	}

	return h.tree.nodes[0].value, h.tree.nodes[0].alias
}

// GetMax returns max value (root) from heap tree with alias.
func (h *MaxHeap) GetMax() (int, any) {
	if len(h.tree.nodes) == 0 {
		return 0, nil
	}

	return h.tree.nodes[0].value, h.tree.nodes[0].alias
}

// NewMin creates a new instance of a min binary heap.
func NewMin() *MinHeap {
	return &MinHeap{
		tree: &tree{
			nodes: make([]*node, 0),
		},
	}
}

// NewMax creates a new instance of a max binary heap.
func NewMax() *MaxHeap {
	return &MaxHeap{
		tree: &tree{
			nodes: make([]*node, 0),
		},
	}
}
