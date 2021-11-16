# Result

## 500MB File

```
$ go test -bench="^BenchmarkCheckSumReaderWithBufferSize.*KB_500MB$" -benchtime=5x
goos: linux
goarch: amd64
pkg: github.com/CooperLi/hash/pkg/reader
cpu: Common KVM processor
BenchmarkCheckSumReaderWithBufferSize_1KB_500MB-8              5        4315557304 ns/op            5352 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_4KB_500MB-8              5        4289598441 ns/op            8424 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_8KB_500MB-8              5        4131381234 ns/op           12520 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_16KB_500MB-8             5        4079235861 ns/op           20712 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_32KB_500MB-8             5        3998029503 ns/op           37096 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_64KB_500MB-8             5        3945533353 ns/op           69864 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_128KB_500MB-8            5        3984920781 ns/op          135400 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_256KB_500MB-8            5        3962856217 ns/op          266472 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_512KB_500MB-8            5        4003625136 ns/op          528616 B/op          6 allocs/op
PASS
ok      github.com/CooperLi/hash/pkg/reader     285.235s
```

## 1GB File

```
$ go test -bench="^BenchmarkCheckSumReaderWithBufferSize.*KB_1GB$" -benchtime=5x
goos: linux
goarch: amd64
pkg: github.com/CooperLi/hash/pkg/reader
cpu: Common KVM processor
BenchmarkCheckSumReaderWithBufferSize_1KB_1GB-8                5        8850595397 ns/op            5332 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_4KB_1GB-8                5        8782100158 ns/op            8404 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_8KB_1GB-8                5        8352574262 ns/op           12500 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_16KB_1GB-8               5        8375086235 ns/op           20692 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_32KB_1GB-8               5        8172201278 ns/op           37076 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_64KB_1GB-8               5        8271554308 ns/op           69844 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_128KB_1GB-8              5        8125593621 ns/op          135380 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_256KB_1GB-8              5        8145840583 ns/op          266452 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_512KB_1GB-8              5        8118965891 ns/op          528596 B/op          6 allocs/op
PASS
ok      github.com/CooperLi/hash/pkg/reader     584.704s
```

## 2GB File

```
$ go test -bench="^BenchmarkCheckSumReaderWithBufferSize.*KB_2GB$" -benchtime=5x
goos: linux
goarch: amd64
pkg: github.com/CooperLi/hash/pkg/reader
cpu: Common KVM processor
BenchmarkCheckSumReaderWithBufferSize_1KB_2GB-8                5        8856545032 ns/op            5332 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_4KB_2GB-8                5        8960816857 ns/op            8404 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_8KB_2GB-8                5        8412535145 ns/op           12500 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_16KB_2GB-8               5        8330972357 ns/op           20692 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_32KB_2GB-8               5        8125097017 ns/op           37076 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_64KB_2GB-8               5        8232215227 ns/op           69844 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_128KB_2GB-8              5        8131110497 ns/op          135380 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_256KB_2GB-8              5        8110579013 ns/op          266452 B/op          6 allocs/op
BenchmarkCheckSumReaderWithBufferSize_512KB_2GB-8              5        8061413229 ns/op          528596 B/op          6 allocs/op
PASS
ok      github.com/CooperLi/hash/pkg/reader     585.393s
```

