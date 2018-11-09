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
go version go1.11.2 darwin/amd64
go test -cpu 1,2,4 -count  5 -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/tai-ga/mapbench
Benchmark_Map/sync.RWMutex                500000              8754 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              9080 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              9209 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              8386 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                500000              8837 ns/op             734 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000             11260 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              9719 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000              9509 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000             10613 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              200000             11108 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              8648 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              9124 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000             10775 ns/op             673 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000             10268 ns/op             671 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              200000              7574 ns/op             672 B/op         52 allocs/op
Benchmark_Map/sync.Map                    300000              8241 ns/op             802 B/op         59 allocs/op
Benchmark_Map/sync.Map                    500000              8349 ns/op             803 B/op         59 allocs/op
Benchmark_Map/sync.Map                    300000              8801 ns/op             851 B/op         60 allocs/op
Benchmark_Map/sync.Map                    300000              9675 ns/op             912 B/op         60 allocs/op
Benchmark_Map/sync.Map                    500000              9891 ns/op             910 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  300000              6559 ns/op             881 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  500000              6132 ns/op             819 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  500000              6615 ns/op             910 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  300000              6921 ns/op             949 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  500000              6079 ns/op             883 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              6376 ns/op             933 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  500000              5414 ns/op             833 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  500000              5846 ns/op             879 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  500000              6939 ns/op             950 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  300000              6550 ns/op             952 B/op         60 allocs/op
Benchmark_Map/concurrent-map              300000              5791 ns/op             704 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5968 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5894 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5838 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map              500000              5732 ns/op             734 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3332 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3201 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3211 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3286 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              3288 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              3020 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              2886 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              3031 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              2943 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              3065 ns/op             736 B/op         52 allocs/op
PASS
ok      github.com/tai-ga/mapbench      132.888s
```
