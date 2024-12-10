package collection

import "sync"

type Map[T comparable, U any] struct {
	mu      sync.RWMutex
	storage map[T]U
}

func NewMap[T comparable, U any]() *Map[T, U] {
	return &Map[T, U]{
		storage: make(map[T]U),
	}
}

func (m *Map[T, U]) Write(write func(storage map[T]U)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	write(m.storage)
}

func (m *Map[T, U]) Read(read func(storage map[T]U) any) any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return read(m.storage)
}
