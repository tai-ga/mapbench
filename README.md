# Concurrent map benchmark

## Result

```
go version go1.19.5 darwin/arm64
go test -cpu 1,2,4,8 -count  5 -benchmem -bench .
goos: darwin
goarch: arm64
pkg: github.com/tai-ga/mapbench
Benchmark_Map/sync.RWMutex                400476              3048 ns/op             677 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                749340              3617 ns/op             685 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                757352              3826 ns/op             683 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                691563              3754 ns/op             691 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex                763464              3630 ns/op             683 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              405060              3718 ns/op             676 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              385776              3839 ns/op             680 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              353296              3450 ns/op             687 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              414694              4012 ns/op             675 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-2              357106              3719 ns/op             685 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              351477              3905 ns/op             687 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              376802              4542 ns/op             681 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              407166              4199 ns/op             676 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              410596              4488 ns/op             675 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-4              317052              4976 ns/op             697 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              281151              4099 ns/op             713 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              305932              4450 ns/op             702 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              283983              4295 ns/op             711 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              326572              3999 ns/op             694 B/op         52 allocs/op
Benchmark_Map/sync.RWMutex-8              320192              4338 ns/op             696 B/op         52 allocs/op
Benchmark_Map/sync.Map                    669997              4136 ns/op             857 B/op         59 allocs/op
Benchmark_Map/sync.Map                    662700              4187 ns/op             851 B/op         59 allocs/op
Benchmark_Map/sync.Map                    655581              4040 ns/op             847 B/op         59 allocs/op
Benchmark_Map/sync.Map                    685515              4268 ns/op             864 B/op         59 allocs/op
Benchmark_Map/sync.Map                    698823              4527 ns/op             903 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  695235              2899 ns/op             826 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  834380              3234 ns/op             835 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  693541              3957 ns/op             954 B/op         60 allocs/op
Benchmark_Map/sync.Map-2                  646026              3336 ns/op             896 B/op         59 allocs/op
Benchmark_Map/sync.Map-2                  733927              3379 ns/op             809 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  760915              3380 ns/op             959 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  882100              2786 ns/op             891 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  814069              2952 ns/op             855 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  888970              3111 ns/op             905 B/op         59 allocs/op
Benchmark_Map/sync.Map-4                  854244              2853 ns/op             920 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  730062              3444 ns/op             875 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  699572              2734 ns/op             865 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  768628              3026 ns/op             847 B/op         59 allocs/op
Benchmark_Map/sync.Map-8                  460105              2545 ns/op             943 B/op         60 allocs/op
Benchmark_Map/sync.Map-8                  636290              2792 ns/op             868 B/op         59 allocs/op
Benchmark_Map/concurrent-map              644538              2924 ns/op             698 B/op         52 allocs/op
Benchmark_Map/concurrent-map              701340              3166 ns/op             690 B/op         52 allocs/op
Benchmark_Map/concurrent-map              606000              2909 ns/op             706 B/op         52 allocs/op
Benchmark_Map/concurrent-map              698432              3121 ns/op             690 B/op         52 allocs/op
Benchmark_Map/concurrent-map              693860              2915 ns/op             691 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            900393              1632 ns/op             753 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              1676 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            889299              1624 ns/op             756 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2            945927              1701 ns/op             745 B/op         52 allocs/op
Benchmark_Map/concurrent-map-2           1000000              1645 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              1065 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              1084 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              1098 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              1103 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-4           1000000              1086 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8           1000000              1128 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8           1000000              1126 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8           1000000              1102 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8           1000000              1081 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map-8           1000000              1097 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2             688249              2996 ns/op             694 B/op         52 allocs/op
Benchmark_Map/concurrent-map2             695569              3132 ns/op             690 B/op         52 allocs/op
Benchmark_Map/concurrent-map2             701654              3098 ns/op             690 B/op         52 allocs/op
Benchmark_Map/concurrent-map2             656288              3099 ns/op             696 B/op         52 allocs/op
Benchmark_Map/concurrent-map2             693354              2994 ns/op             691 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-2           955591              1719 ns/op             743 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-2           956334              1659 ns/op             743 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-2           920841              1612 ns/op             749 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-2          1000000              1686 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-2           963468              1676 ns/op             742 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-4          1000000              1084 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-4          1000000              1071 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-4          1000000              1132 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-4          1000000              1084 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-4          1000000              1100 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-8          1000000              1097 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-8          1000000              1100 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-8          1000000              1095 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-8          1000000              1138 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2-8          1000000              1124 ns/op             736 B/op         52 allocs/op
Benchmark_Map/concurrent-map2APIv2        874141              2792 ns/op             439 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2        890276              2795 ns/op             435 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2        933645              2771 ns/op             427 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2        830718              2567 ns/op             356 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2        936511              2765 ns/op             427 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-2     1000000              1419 ns/op             416 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-2     1000000              1334 ns/op             416 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-2     1000000              1428 ns/op             416 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-2     1000000              1382 ns/op             416 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-2     1000000              1398 ns/op             416 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-4     1301401               924.7 ns/op           379 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-4     1349997               973.6 ns/op           375 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-4     1327069               951.7 ns/op           377 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-4     1335458               919.5 ns/op           376 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-4     1399525               939.9 ns/op           372 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-8     1308805               906.1 ns/op           378 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-8     1296559               913.4 ns/op           379 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-8     1295713               902.8 ns/op           379 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-8     1343952               904.3 ns/op           375 B/op         32 allocs/op
Benchmark_Map/concurrent-map2APIv2-8     1306411               899.7 ns/op           378 B/op         32 allocs/op
PASS
ok      github.com/tai-ga/mapbench      184.397s
```
