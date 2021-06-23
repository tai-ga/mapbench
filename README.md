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
go version go1.16.5 darwin/amd64
go test -cpu 1,2,4,8 -count  5 -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/tai-ga/mapbench
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
Benchmark_Map/sync.RWMutex                362870              8321 ns/op             684 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                380793              8070 ns/op             681 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                390444              7983 ns/op             679 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                392809              7667 ns/op             678 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                409642              8195 ns/op             676 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              185540              8040 ns/op             677 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              174224              8597 ns/op             683 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              183075              8675 ns/op             679 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              226926              8554 ns/op             745 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              182582              8563 ns/op             678 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              176948              7282 ns/op             681 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              179995              6969 ns/op             679 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              185884              6933 ns/op             677 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              190986              7282 ns/op             675 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              185503              8129 ns/op             677 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              170602              8602 ns/op             684 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              171357              8328 ns/op             684 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              174945              8106 ns/op             682 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              178005              8123 ns/op             681 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              146167              7927 ns/op             701 B/op         51 allocs/op
Benchmark_Map/sync.Map                    360312              8283 ns/op             786 B/op         59 allocs/op
Benchmark_Map/sync.Map                    352556              8750 ns/op             888 B/op         59 allocs/op
Benchmark_Map/sync.Map                    332758              8043 ns/op             783 B/op         59 allocs/op
Benchmark_Map/sync.Map                    345384              8446 ns/op             858 B/op         59 allocs/op
Benchmark_Map/sync.Map                    322189              9122 ns/op             986 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  306418              4450 ns/op             833 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  444670              6073 ns/op             984 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  363121              6159 ns/op            1013 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  347695              4313 ns/op             786 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  422526              5420 ns/op             866 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  343027              5545 ns/op            1060 B/op         60 allocs/op
Benchmark_Map/sync.Map-4                  367998              3355 ns/op             793 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  515540              5128 ns/op            1013 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  409647              3880 ns/op             840 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  320163              5538 ns/op            1058 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  509028              4226 ns/op             891 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  334224              3928 ns/op             844 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  306444              3980 ns/op             854 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  508755              5687 ns/op            1070 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  377551              6024 ns/op            1047 B/op         59 allocs/op
Benchmark_Map/concurrent-map              440830              5768 ns/op             755 B/op         52 allocs/op
Benchmark_Map/concurrent-map              392472              5648 ns/op             677 B/op         52 allocs/op
Benchmark_Map/concurrent-map              402418              5466 ns/op             676 B/op         52 allocs/op
Benchmark_Map/concurrent-map              403568              5686 ns/op             675 B/op         52 allocs/op
Benchmark_Map/concurrent-map              409058              5393 ns/op             675 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            672831              2469 ns/op             694 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            674583              2466 ns/op             693 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            679269              2446 ns/op             693 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            640879              2460 ns/op             699 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            686034              2534 ns/op             692 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            987357              1661 ns/op             738 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            729748              1781 ns/op             686 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            888522              1820 ns/op             756 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            848949              1892 ns/op             675 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4            864320              1864 ns/op             761 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8            841507              1779 ns/op             676 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8            772796              1842 ns/op             682 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8            726567              1876 ns/op             687 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8            717614              1858 ns/op             688 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8            804320              1797 ns/op             679 B/op         52 allocs/op
PASS
ok      github.com/tai-ga/mapbench      118.578s
```
