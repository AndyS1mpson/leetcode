package custom_structs

import "golang.org/x/exp/rand"

/**
Задача:

	Реализуйте структуру RandomizedSet, которая:
	    * insert(val) - добавляет val в множество, возвращает true если val не был уже в множестве, иначе возвращает false
		* remove(val) - удаляет val из множества, возвращает true если val был в множестве, иначе возвращает false
		* getRandom() -  Возвращает случайный элемент из текущего набора элементов (гарантируется,
			что на момент вызова этого метода существует хотя бы один элемент).
			Каждый элемент должен иметь одинаковую вероятность быть возвращенным.

		Каждая функция должна иметь сложность O(1)
*/
type RandomizedSet struct {
	arr  []int
	size int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr:  make([]int, 0),
		size: 0,
		set: make(map[int]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.set[val]
	if ok {
		return false
	}

	rs.arr = append(rs.arr, val)
	rs.set[val] = rs.size

	rs.size++
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.set[val]
	if !ok {
		return false
	}

	// Swap the element to remove with the last element
	lastElement := rs.arr[rs.size-1]
	rs.arr[index] = lastElement
	rs.set[lastElement] = index

	// Remove the last element
	rs.arr = rs.arr[:rs.size-1]
	delete(rs.set, val)
	rs.size--
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	index := rand.Intn(rs.size)
	return rs.arr[index]
}
