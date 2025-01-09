package custom_structs

/*
Задача:

	Сделать структуру, которая реализует LRU cache.
	LRU cache - это кэш с ограниченным размером, который автоматически удаляет наименее недавно использованные
	элементы при необходимости освободить место для новых данных.

	Каждая операция должна выполняться за O(1)
*/
type listNode struct {
	key   int
	value int

	prev *listNode
	next *listNode
}

type LRUCache struct {
	data     map[int]*listNode
	top      *listNode
	bottom   *listNode
	capacity int
}

func LRUConstructor(capacity int) LRUCache {
	top := &listNode{}
	bottom := &listNode{}

	bottom.next = top
	top.prev = bottom

	return LRUCache{
		data:     make(map[int]*listNode, capacity),
		top:      top,
		bottom:   bottom,
		capacity: capacity,
	}
}

// Get получение элемента из кэша
func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		this.moveNodeToTop(v)

		return v.value
	}

	return -1
}

// Put вставка элемента в кэш
func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.data[key]; ok {
		v.value = value
		this.moveNodeToTop(v)
	} else {
		// если переполнено, удаляем самый старый элемент
		if len(this.data) == this.capacity {
			bottom := this.bottom.next
			this.bottom.next = bottom.next
			bottom.next.prev = this.bottom
			delete(this.data, bottom.key)
		}
		newNode := &listNode{
			key:   key,
			value: value,
		}

		this.addNodeToTop(newNode)
		this.data[key] = newNode
	}
}

// moveNodeToTop перенос элемента на верх кэша
func (this *LRUCache) moveNodeToTop(node *listNode) {
	// соединяем предыдущий и следующий за нодой элементы
	node.prev.next = node.next
	node.next.prev = node.prev

	// ставим ноду на самый верх (кроме вершины top, она техническая и никогда не меняется как и bottom, т е будет top -> node -> ... -> bottom)
	node.prev = this.top.prev
	this.top.prev.next = node
	this.top.prev = node
	node.next = this.top
}

// addNodeToTop добавление элемента на верх кэша
func (this *LRUCache) addNodeToTop(node *listNode) {
	node.prev = this.top.prev
	node.next = this.top
	node.prev.next = node
	this.top.prev = node
}
