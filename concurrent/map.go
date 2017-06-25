package concurrent

import "sync"

type Map struct {
	data  map[string]interface{}
	mutex sync.Mutex
}

type MapElement struct {
	Key   string
	Value interface{}
}

func NewMap() *Map {
	return &Map{
		data: make(map[string]interface{}),
	}
}

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

func (m *Map) Put(k string, v interface{}) {
	m.mutex.Lock()
	m.data[k] = v
	m.mutex.Unlock()
}

func (m *Map) Get(k string) (interface{}, bool) {
	m.mutex.Lock()
	v, ok := m.data[k]
	m.mutex.Unlock()
	return v, ok
}
