# Concurrent map benchmark

```go
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
```

## Result
```
go test -cpu 1,2,4 -count  5 -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/tai-ga/mapbench
Benchmark_Map/sync.RWMutex                300000              8890 ns/op             705 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              9224 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              9032 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000             10063 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                300000              9253 ns/op             704 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              8494 ns/op             673 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              8875 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              8451 ns/op             673 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              8570 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              9533 ns/op             673 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              8465 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              300000              6119 ns/op             704 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              7055 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              7794 ns/op             671 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              6950 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.Map                    300000              9856 ns/op             871 B/op         61 allocs/op
Benchmark_Map/sync.Map                    300000              9796 ns/op             806 B/op         59 allocs/op
Benchmark_Map/sync.Map                    300000              9895 ns/op             816 B/op         59 allocs/op
Benchmark_Map/sync.Map                    300000             10803 ns/op             795 B/op         59 allocs/op
Benchmark_Map/sync.Map                    300000              9586 ns/op             785 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  500000              9129 ns/op            1063 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  300000              7144 ns/op             876 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  500000              8256 ns/op             996 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  300000              8443 ns/op            1070 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  300000              6210 ns/op             799 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              6406 ns/op             851 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              6246 ns/op             875 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              7450 ns/op             952 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              8502 ns/op            1044 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  500000              8831 ns/op            1088 B/op         59 allocs/op
Benchmark_Map/concurrent-map              500000              5747 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5787 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              6248 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              300000              5535 ns/op             704 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5731 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3309 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            500000              3497 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            500000              3242 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3227 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3440 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            500000              3020 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              2999 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              2998 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              3169 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              2995 ns/op             736 B/op         52 allocs/op
PASS
ok      github.com/tai-ga/mapbench      122.422s
```
