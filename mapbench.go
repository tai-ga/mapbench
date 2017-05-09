package mapbench

import (
	"github.com/golang/sync/syncmap"
	"github.com/orcaman/concurrent-map"
)

func makeSyncMap(n int) int {
	var i int
	m := new(syncmap.Map)
	for i = 0; i < n; i++ {
		m.Store(string(i), i)
		m.Load(string(i))
	}
	for i = 0; i < n; i++ {
		m.Delete(string(i))
	}
	return i
}

func makeConcurrentMap(n int) int {
	var i int
	m := cmap.New()
	for i = 0; i < n; i++ {
		m.Set(string(i), i)
		m.Get(string(i))
	}
	for i = 0; i < n; i++ {
		m.Remove(string(i))
	}
	return i
}
