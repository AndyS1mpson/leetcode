package structs

// Stack реализация дженерик стэка
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

// Pop достает верхний элемент из стэка
func (s *Stack[T]) Pop() (T, bool) {
	var x T
	if len(s.data) != 0 {
		el := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]

		return el, true
	}

	return x, false
}

// Push добавляет элемент на верх стэка
func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

// Len возвращает длину стэка
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// IsEmpty пустой ли стэк
func (s *Stack[T]) IsEmpty() bool {
	return s.Len() == 0
}
