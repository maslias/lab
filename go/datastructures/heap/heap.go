package heap

import "fmt"

type heap struct {
	maxHeap []int
}

func (h *heap) addMaxHeap(v int) {
	h.maxHeap = append(h.maxHeap, v)
	h.heapUp(len(h.maxHeap) - 1)
}

func (h *heap) heapUp(idx int) {
	if idx == 0 {
		return
	}

	cv := h.maxHeap[idx]
	pidx, pv := h.getParent(idx)

	if cv <= pv {
		return
	}

	h.swap(idx, pidx)
	h.heapUp(pidx)
}

func (h *heap) removeMaxHeap() {
	l := len(h.maxHeap)
	if l == 0 {
		return
	}
	h.maxHeap[0] = h.maxHeap[l-1]
	h.maxHeap = h.maxHeap[:l-1]
	h.heapDown(0)
}

func (h *heap) heapDown(idx int) {
	pv := h.maxHeap[idx]
	clidx, clv, clerr := h.getLeftChild(idx)
	cridx, crv, _ := h.getRightChild(idx)

	if idx >= len(h.maxHeap) || clerr != nil {
		return
	}

	if pv >= clv && pv >= crv {
		return
	}

	if clv <= crv && pv <= crv {
		h.swap(idx, cridx)
		h.heapDown(cridx)
	} else {
		h.swap(idx, clidx)
		h.heapDown(clidx)
	}
}

func (h *heap) getLeftChild(idx int) (int, int, error) {
	cidx := (idx * 2) + 1
	if cidx >= len(h.maxHeap) {
		return 0, 0, fmt.Errorf("index is out of range")
	}
	return cidx, h.maxHeap[cidx], nil
}

func (h *heap) getRightChild(idx int) (int, int, error) {
	cidx := (idx * 2) + 2
	if cidx >= len(h.maxHeap) {
		return 0, 0, fmt.Errorf("index is out of range")
	}
	return cidx, h.maxHeap[cidx], nil
}

func (h *heap) getParent(idx int) (int, int) {
	pidx := (idx - 1) / 2
	return pidx, h.maxHeap[pidx]
}

func (h *heap) swap(idx1 int, idx2 int) {
	h.maxHeap[idx1], h.maxHeap[idx2] = h.maxHeap[idx2], h.maxHeap[idx1]
}

func (h *heap) listStruct() {
	fmt.Println(h.maxHeap)
}

func TestStructHeap() {
	h := heap{}
	h.addMaxHeap(50)
	h.addMaxHeap(100)
	h.addMaxHeap(25)
	h.addMaxHeap(12)
	h.addMaxHeap(75)
	h.addMaxHeap(150)
	h.listStruct()
	h.removeMaxHeap()
	h.listStruct()
	h.removeMaxHeap()
	h.listStruct()
	h.removeMaxHeap()
	h.listStruct()
}
