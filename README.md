# Concurrent map benchmark

```go
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
```

## Result
```
go test -cpu 1,2,4,8,16 -count  5 -benchmem -bench .
BenchmarkSyncMap                 1000000              3474 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap                 1000000              3078 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap                 1000000              2897 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap                 1000000              3019 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap                 1000000              3078 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-2               1000000              1759 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-2               1000000              1717 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-2               1000000              1685 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-2               1000000              2001 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-2               1000000              1783 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-4               1000000              1669 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-4               1000000              1762 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-4               1000000              1752 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-4               1000000              1762 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-4               1000000              1807 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-8               1000000              1711 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-8               1000000              1690 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-8               1000000              1643 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-8               1000000              1682 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-8               1000000              1756 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-16              1000000              1781 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-16              1000000              1712 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-16              1000000              1712 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-16              1000000              1701 ns/op             196 B/op          6 allocs/op
BenchmarkSyncMap-16              1000000              1642 ns/op             196 B/op          6 allocs/op
BenchmarkConcurrentMap           1000000              1265 ns/op             176 B/op          2 allocs/op
BenchmarkConcurrentMap           1000000              1053 ns/op             176 B/op          2 allocs/op
BenchmarkConcurrentMap           1000000              1274 ns/op             175 B/op          2 allocs/op
BenchmarkConcurrentMap           1000000              1140 ns/op             176 B/op          2 allocs/op
BenchmarkConcurrentMap           1000000              1051 ns/op             176 B/op          2 allocs/op
BenchmarkConcurrentMap-2         2000000               608 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-2         2000000               596 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-2         2000000               612 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-2         2000000               606 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-2         2000000               595 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-4         2000000               585 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-4         2000000               600 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-4         2000000               583 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-4         2000000               585 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-4         2000000               594 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-8         2000000               580 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-8         2000000               567 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-8         2000000               575 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-8         2000000               580 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-8         2000000               582 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-16        2000000               577 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-16        2000000               570 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-16        2000000               586 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-16        2000000               584 ns/op              96 B/op          2 allocs/op
BenchmarkConcurrentMap-16        2000000               579 ns/op              96 B/op          2 allocs/op
PASS
ok      github.com/tai-ga/mapbench      98.286s
```
