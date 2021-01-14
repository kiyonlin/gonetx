This project is archived, please move to [gonetx/ipconv](https://github.com/gonetx/ipconv).

# Benchmark
Here is the benchmark with `2.4 GHz 8-Core Intel Core i9`:
```hs
$ GOMAXPROCS=1 go test -bench='Ipconv' -benchmem -benchtime=10s 
Benchmark_Ipconv_Long2V4        433216750               27.7 ns/op            16 B/op          1 allocs/op
Benchmark_Ipconv_V42Long        652684755               18.5 ns/op             0 B/op          0 allocs/op
```
