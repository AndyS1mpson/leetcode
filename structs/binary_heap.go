package structs

/*
Бинарная куча - двоичное подвешенное дерево, для которого выполнены следующие три условия:
	- Значение в любой вершине не больше (если куча для минимума), чем значения ее потомков
	- На i-ом слое 2^i вершин, кроме последнего. Слои нумеруются с нуля
	- Последний слой заполнен слева направо
*/
type BinaryHeap struct {
	array []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}

func (h *BinaryHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.siftUp(len(h.array) - 1)
}

func (h *BinaryHeap) Delete() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	min := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.siftDown(0)

	return min
}

/*
Если значение измененного элемента уменьшается, то свойства кучи восстанавливаются функцией siftUp.
Работа:
	Если элемент больше своего отца, условие 1 соблюдено и делать ничего не нужно. 
	Иначе, мы меняем местами его с отцом. После чего выполняем siftUp для этого отца
Сложность: O(log n)
*/
func (h *BinaryHeap) siftUp(idx int) {
	parentIdx := (idx - 1) / 2

	for idx > 0 && h.array[idx] < h.array[parentIdx] {
		h.array[idx], h.array[parentIdx] = h.array[parentIdx], h.array[idx]
		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

/*
Если значение измененного элемента увеличивается, то свойства кучи восстанавлиаются функцией siftDown.
Работа:
	Если i-й элемент меньше, чем его сыновья, все поддерево уже является кучей и делать ничего не надо.
	В противном случае меняем местами i-й элемент с наименьшим из его сыновей, после чего выполняем siftDown для сына
Сложность: O(log n)
*/
func (h *BinaryHeap) siftDown(idx int) {
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := 2*idx + 2
		smallestIdx := idx

		if leftChildIdx < len(h.array) && h.array[rightChildIdx] < h.array[smallestIdx] {
			smallestIdx = rightChildIdx
		}

		if smallestIdx == idx {
			break
		}

		h.array[idx], h.array[smallestIdx] = h.array[smallestIdx], h.array[idx]
		idx = smallestIdx
	}
}
