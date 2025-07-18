# Go 1.25 `reflect.TypeAssert` vs type assertion

```sh
go test -bench '^BenchmarkTypeAssertion$|^BenchmarkReflectTypeAssert$' -count 6
goos: linux
goarch: amd64
pkg: github.com/miyamo2/reflect_type_assert_vs_type_assertion
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkTypeAssertion-16               35453163                34.60 ns/op           24 B/op          1 allocs/op
BenchmarkTypeAssertion-16               34472107                35.62 ns/op           24 B/op          1 allocs/op
BenchmarkTypeAssertion-16               35326629                34.59 ns/op           24 B/op          1 allocs/op
BenchmarkTypeAssertion-16               35174138                34.63 ns/op           24 B/op          1 allocs/op
BenchmarkTypeAssertion-16               34724137                34.60 ns/op           24 B/op          1 allocs/op
BenchmarkTypeAssertion-16               33512199                34.64 ns/op           24 B/op          1 allocs/op
BenchmarkReflectTypeAssert-16           294901711                4.058 ns/op           0 B/op          0 allocs/op
BenchmarkReflectTypeAssert-16           293091184                4.113 ns/op           0 B/op          0 allocs/op
BenchmarkReflectTypeAssert-16           295116627                4.052 ns/op           0 B/op          0 allocs/op
BenchmarkReflectTypeAssert-16           296398741                4.044 ns/op           0 B/op          0 allocs/op
BenchmarkReflectTypeAssert-16           296427202                4.047 ns/op           0 B/op          0 allocs/op
BenchmarkReflectTypeAssert-16           296555486                4.043 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/miyamo2/reflect_type_assert_vs_type_assertion        14.460s
```

## benchstat

### sec/op

| TypeAssertion-16 | ReflectTypeAssert-16 |
|------------------|----------------------|
| 34.62n ± 3%      | 4.049n ± 3%          |

### B/op

| TypeAssertion-16 | ReflectTypeAssert-16 |
|------------------|----------------------|
| 24.00 ± 0%       | 0.000 ± 0%           |

### allocs/op

| TypeAssertion-16 | ReflectTypeAssert-16 |
|------------------|----------------------|
| 1.000 ± 0%       | 0.000 ± 0%           |
