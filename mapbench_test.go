package sync

import (
	"strconv"
	"sync"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
	cmap2 "github.com/orcaman/concurrent-map/v2"
)

type apiVersion int

const (
	mapAPIv1 apiVersion = iota + 1
	mapAPIv2
)

// MapAPIv1 is the same API as sync.Map
type MapAPIv1 interface {
	Delete(key any)
	Load(key any) (value any, ok bool)
	LoadOrStore(key, value any) (actual any, loaded bool)
	Range(f func(key, value any) bool)
	Store(key, value any)
}

// MapAPIv2 is implementation using Generics
type MapAPIv2 interface {
	Delete2(key string)
	Load2(key string) (value any, ok bool)
	LoadOrStore2(key string, value any) (actual any, loaded bool)
	Range2(f func(key string, value any) bool)
	Store2(key string, value any)
}

var _ MapAPIv1 = (*sync.Map)(nil)
var _ MapAPIv1 = (*RWMutexMap)(nil)
var _ MapAPIv1 = (*CMap)(nil)
var _ MapAPIv1 = (*CMap2)(nil)
var _ MapAPIv2 = (*CMap2)(nil)

type CMap struct {
	m cmap.ConcurrentMap
}

func NewCMap() *CMap {
	c := new(CMap)
	c.m = cmap.New()
	return c
}

func toString(key any) string {
	k, ok := key.(string)
	if ok {
		return k
	}
	return ""
}

func (c *CMap) Delete(key any) {
	c.m.Remove(toString(key))
}

func (c *CMap) Load(key any) (value any, ok bool) {
	return c.m.Get(toString(key))
}

func (c *CMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	ok := c.m.SetIfAbsent(toString(key), value)
	return value, ok
}

func (c *CMap) Range(f func(key, value any) bool) {
	c.m.IterCb(func(k string, v any) {
		f(k, v)
	})
}

func (c *CMap) Store(key, value any) {
	c.m.Set(toString(key), value)
}

// CMap2 is Version 2 of orcaman/concurrent-map
type CMap2 struct {
	m cmap2.ConcurrentMap[string, any]
}

func NewCMap2() *CMap2 {
	c := new(CMap2)
	c.m = cmap2.New[any]()
	return c
}

func (c *CMap2) Delete(key any) {
	c.m.Remove(toString(key))
}

func (c *CMap2) Load(key any) (value any, ok bool) {
	return c.m.Get(toString(key))
}

func (c *CMap2) LoadOrStore(key, value any) (actual any, loaded bool) {
	ok := c.m.SetIfAbsent(toString(key), value)
	return value, ok
}

func (c *CMap2) Range(f func(key, value any) bool) {
	c.m.IterCb(func(k string, v any) {
		f(k, v)
	})
}

func (c *CMap2) Store(key, value any) {
	c.m.Set(toString(key), value)
}

func (c *CMap2) Delete2(key string) {
	c.m.Remove(key)
}

func (c *CMap2) Load2(key string) (value any, ok bool) {
	return c.m.Get(key)
}

func (c *CMap2) LoadOrStore2(key string, value any) (actual any, loaded bool) {
	ok := c.m.SetIfAbsent(key, value)
	return value, ok
}

func (c *CMap2) Range2(f func(key string, value any) bool) {
	c.m.IterCb(func(k string, v any) {
		f(k, v)
	})
}

func (c *CMap2) Store2(key string, value any) {
	c.m.Set(key, value)
}

type RWMutexMap struct {
	m  map[any]any
	mu sync.RWMutex
}

func (m *RWMutexMap) Delete(key any) {
	_, ok := m.Load(key)
	if !ok {
		return
	}
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *RWMutexMap) Load(key any) (value any, ok bool) {
	if m.m == nil {
		m.m = make(map[any]any)
	}

	m.mu.RLock()
	v, ok := m.m[key]
	m.mu.RUnlock()
	return v, ok
}

func (m *RWMutexMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	v, ok := m.Load(key)
	if ok {
		return v, true
	}
	m.Store(key, value)
	return value, false
}

func (m *RWMutexMap) Range(f func(key, value any) bool) {
	if m.m == nil {
		m.m = make(map[any]any)
	}

	m.mu.RLock()
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
	m.mu.RUnlock()
}

func (m *RWMutexMap) Store(key, value any) {
	if m.m == nil {
		m.m = make(map[any]any)
	}

	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

type Map interface {
	Delete(key any)
	Load(key any) (value any, ok bool)
	LoadOrStore(key, value any) (actual any, loaded bool)
	Range(f func(key, value any) bool)
	Store(key, value any)
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
