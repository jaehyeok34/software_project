package collection

import (
	"sync"
)

type Map[T, U any] struct {
	storage sync.Map
	length  int
}

func NewMap[T, U any]() *Map[T, U] {
	return &Map[T, U]{}
}

func (m *Map[T, U]) Store(key T, value U) {
	m.storage.Store(key, value)
	m.length++
}

func (m *Map[T, U]) Load(key T) (value U, ok bool) {
	v, ok := m.storage.Load(key)
	if !ok {
		return *new(U), false
	}

	return v.(U), ok
}

func (m *Map[T, U]) Length() int {
	return m.length
}

func (m *Map[T, U]) Delete(key T) {
	m.storage.Delete(key)
	m.length--
}

func (m *Map[T, U]) GetAll() []U {
	var values []U
	m.storage.Range(func(key, value any) bool {
		values = append(values, value.(U))
		return true
	})

	return values
}
