package tasks

// /*
// Задача:

// 	Сделать структуру, которая реализует LRU cache.
// 	LRU cache - это кэш с ограниченным размером, который автоматически удаляет наименее недавно использованные
// 	элементы при необходимости освободить место для новых данных.

// 	Каждая операция должна выполняться за O(1)
// */
// type listNode struct {
// 	key   int
// 	value int

// 	prev *listNode
// 	next     *listNode
// }

// type LRUCache struct {
// 	data     map[int]*listNode
// 	top      *listNode
// 	bottom   *listNode
// 	capacity int
// }

// func LRUConstructor(capacity int) LRUCache {
// 	top := &listNode{}
// 	bottom := &listNode{}

// 	bottom.next = top
// 	top.prev = bottom

// 	return LRUCache{
// 		data:     make(map[int]*listNode, capacity),
// 		top:      top,
// 		bottom:   bottom,
// 		capacity: capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {

// }

// func (this *LRUCache) Put(key int, value int) {

// }

// func (this *LRUCache) moveNodeToTop(node *listNode) {
// 	// соединяем предыдущий и следующий за нодой элементы
// 	node.prev.next = node.next
// 	node.next.prev = node.prev

// 	// ставим ноду на самый верх (кроме вершины top, она техническая и никогда не меняется как и bottom, т е будет top -> node -> ... -> bottom)
// 	node.prev = this.top.prev
// 	this.top.prev.next = node
// 	this.top.prev = node
// 	node.next = this.top
// }

// func (this *LRUCache) m
