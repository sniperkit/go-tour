package concurrent

import "sync"

// A map that is safe for concurrent access.
type Map struct {
	data  map[string]interface{}
	mutex sync.Mutex
}

// Creates a new concurrent map.
func NewMap() *Map {
	return &Map{
		data: make(map[string]interface{}),
	}
}

// If the key k is absent from the map, (k,v) is put into the map.
// The inserted value and true is returned. Otherwise, the map is not mutated
// and the existing value at k and false is returned.
func (m *Map) PutIfAbsent(k string, v interface{}) (interface{}, bool) {
	m.mutex.Lock()
	if v, ok := m.data[k]; ok {
		m.mutex.Unlock()
		return v, false
	}
	m.data[k] = v
	m.mutex.Unlock()
	return v, true
}

// (k,v) are put into the map. The existing value at k (if any) is
// overwritten.
func (m *Map) Put(k string, v interface{}) {
	m.mutex.Lock()
	m.data[k] = v
	m.mutex.Unlock()
}

// If the key k exists in the map, the value at k and true is returned.
// Otherwise a non-deterministic value and false is returned.
func (m *Map) Get(k string) (interface{}, bool) {
	m.mutex.Lock()
	v, ok := m.data[k]
	m.mutex.Unlock()
	return v, ok
}
