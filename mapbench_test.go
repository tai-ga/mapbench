package sync

import (
	"strconv"
	"sync"
	"testing"

	"github.com/orcaman/concurrent-map"
)

type CMap struct {
	m cmap.ConcurrentMap
}

func NewCMap() *CMap {
	c := new(CMap)
	c.m = cmap.New()
	return c
}

func toString(key interface{}) string {
	k, ok := key.(string)
	if ok {
		return k
	}
	return ""
}

func (c *CMap) Delete(key interface{}) {
	c.m.Remove(toString(key))
}

func (c *CMap) Load(key interface{}) (value interface{}, ok bool) {
	return c.m.Get(toString(key))
}

func (c *CMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	ok := c.m.SetIfAbsent(toString(key), value)
	return value, ok
}

func (c *CMap) Range(f func(key, value interface{}) bool) {
	c.m.IterCb(func(k string, v interface{}) {
		f(k, v)
	})
}

func (c *CMap) Store(key, value interface{}) {
	c.m.Set(toString(key), value)
}

type RWMutexMap struct {
	m  map[interface{}]interface{}
	mu sync.RWMutex
}

func (m *RWMutexMap) Delete(key interface{}) {
	_, ok := m.Load(key)
	if !ok {
		return
	}
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *RWMutexMap) Load(key interface{}) (value interface{}, ok bool) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.RLock()
	v, ok := m.m[key]
	m.mu.RUnlock()
	return v, ok
}

func (m *RWMutexMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	v, ok := m.Load(key)
	if ok {
		return v, true
	}
	m.Store(key, value)
	return value, false
}

func (m *RWMutexMap) Range(f func(key, value interface{}) bool) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.RLock()
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

func (m *RWMutexMap) Store(key, value interface{}) {
	if m.m == nil {
		m.m = make(map[interface{}]interface{})
	}

	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

type Map interface {
	Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
}

func benchmark_Map(m Map, n int) {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				m.Store(strconv.Itoa(i), i)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				m.Load(strconv.Itoa(i))
			}
		}()
	}
	wg.Wait()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				m.Delete(strconv.Itoa(i))
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				m.LoadOrStore(strconv.Itoa(i), i)
				m.LoadOrStore(strconv.Itoa(i), i)
			}
		}()
	}
	wg.Wait()
}

func Benchmark_Map(b *testing.B) {
	b.Run("sync.RWMutex", func(b *testing.B) {
		m := new(RWMutexMap)
		b.ResetTimer()
		benchmark_Map(m, b.N)
	})

	b.Run("sync.Map", func(b *testing.B) {
		m := new(sync.Map)
		b.ResetTimer()
		benchmark_Map(m, b.N)
	})

	b.Run("concurrent-map", func(b *testing.B) {
		m := NewCMap()
		b.ResetTimer()
		benchmark_Map(m, b.N)
	})
}
