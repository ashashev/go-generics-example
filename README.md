# Results

## Slice Contains

### Description

There are three implementations of the same thing. It's a simple for-loop, with
early returning when an element is found.

The generic implementation:
```golang
func Contains[T ~[]I, I comparable](haystack T, needle I) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}
```

And two specific implementations for slice of ints and slice of strings:
```golang
func SliceContainsString(haystack []string, needle string) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}

func SliceContainsInt(haystack []int, needle int) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}
```

### Benchmars results

Slice of strings
|                Benchmark name                         |    Run #1    |    Run #2    |    Run #3    |    Run #4    |    Run #5    |
| ----------------------------------------------------- | ------------ | ------------ |------------- |------------- |------------- |
| `BenchmarkSliceContainsString/Generic_Doesn't_Has-8`  | 14.93 ns/op  | 15.18 ns/op  | 15.17 ns/op  | 15.17 ns/op  | 15.28 ns/op  |
| `BenchmarkSliceContainsString/Generic_Has-8`          | 39.87 ns/op  | 40.66 ns/op  | 40.59 ns/op  | 40.79 ns/op  | 40.85 ns/op  |
| `BenchmarkSliceContainsString/Specific_Doesn't_Has-8` | 14.94 ns/op  | 15.19 ns/op  | 15.22 ns/op  | 15.87 ns/op  | 15.21 ns/op  |
| `BenchmarkSliceContainsString/Specific_Has-8`         | 40.51 ns/op  | 40.63 ns/op  | 40.64 ns/op  | 40.63 ns/op  | 40.81 ns/op  |

Slice of ints
|                Benchmark name                         |    Run #1    |    Run #2    |    Run #3    |    Run #4    |    Run #5    |
| ----------------------------------------------------- | ------------ | ------------ |------------- |------------- |------------- |
| `BenchmarkSliceContainsInt/Generic_Doesn't_Has-8`     |  8.444 ns/op |  8.458 ns/op |  8.443 ns/op |  8.448 ns/op |  8.461 ns/op |
| `BenchmarkSliceContainsInt/Generic_Has-8`             |  4.898 ns/op |  4.913 ns/op |  4.922 ns/op |  4.905 ns/op |  4.920 ns/op |
| `BenchmarkSliceContainsInt/Specific_Doesn't_Has-8`    |  8.444 ns/op |  8.436 ns/op |  8.446 ns/op |  8.447 ns/op |  8.449 ns/op |
| `BenchmarkSliceContainsInt/Specific_Has-8`            |  4.917 ns/op |  4.915 ns/op |  4.922 ns/op |  4.905 ns/op |  4.912 ns/op |


<details>
<summary>Raw benchmarks results</summary>

**Run #1**
```
go test -bench . -benchtime 10s ./example/contains_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	793542500	        14.93 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	300902487	        39.87 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	805521026	        14.94 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	296259190	        40.51 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.444 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.898 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.444 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.917 ns/op
PASS
ok  	command-line-arguments	88.451s
```

**Run #2**
```
go test -bench . -benchtime 10s ./example/contains_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	788970160	        15.18 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	295573492	        40.66 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	762757710	        15.19 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	280926486	        40.63 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.458 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.913 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.436 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.915 ns/op
PASS
ok  	command-line-arguments	87.970s
```

**Run #3**
```
go test -bench . -benchtime 10s ./example/contains_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	791005816	        15.17 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	295192660	        40.59 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	767063816	        15.22 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	282376572	        40.64 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.443 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.922 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.446 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.922 ns/op
PASS
ok  	command-line-arguments	88.104s
```

**Run #4**
```
go test -bench . -benchtime 10s ./example/contains_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	788782045	        15.17 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	295247182	        40.79 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	758607546	        15.87 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	294588723	        40.63 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.448 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.905 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.447 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.905 ns/op
PASS
ok  	command-line-arguments	88.801s
```

**Run #5**
```
go test -bench . -benchtime 10s ./example/contains_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	774865777	        15.28 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	294912528	        40.85 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	790071013	        15.21 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	292575610	        40.81 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.461 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.920 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.449 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.912 ns/op
PASS
ok  	command-line-arguments	88.657s
```
</details>

## Get Keys

### Description

Here I benchmark extracting keys from a map. Again the generic implementation:
```golang
func GetKeys[M ~map[K]V, K comparable, V any](m M) []K {
	data := make([]K, 0, len(m))
	for k := range m {
		data = append(data, k)
	}
	return data
}
```

And the specific implementation:
```golang
func GetKeysStringData(m map[string]*Data) []string {
	data := make([]string, 0, len(m))
	for k := range m {
		data = append(data, k)
	}
	return data
}
```

Type `Data` is defined as
```golang
type Data struct {
	ID     string
	Field1 string
	Field2 float64
	Field3 uint64
}
```

### Benchmars results

Slice of strings
|      Benchmark name         |   Run #1    |   Run #2    |   Run #3    |   Run #4    |   Run #5    |
| --------------------------- | ----------- | ----------- |------------ |------------ |------------ |
| BenchmarkGetKeys/Generic-8  | 611.7 ns/op | 654.8 ns/op | 635.2 ns/op | 611.3 ns/op | 647.8 ns/op |
| BenchmarkGetKeys/Specific-8 | 614.1 ns/op | 727.5 ns/op | 685.7 ns/op | 686.9 ns/op | 599.9 ns/op |

<details>
<summary>Raw benchmarks results</summary>

**Run #1**
```
go test -bench . -benchtime 10s ./example/get_keys_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8         	18520962	       611.7 ns/op
BenchmarkGetKeys/Specific-8        	20124560	       614.1 ns/op
PASS
ok  	command-line-arguments	33.625s
```

**Run #2**
```
go test -bench . -benchtime 10s ./example/get_keys_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8         	26377960	       654.8 ns/op
BenchmarkGetKeys/Specific-8        	23120361	       727.5 ns/op
PASS
ok  	command-line-arguments	35.084s
```

**Run #3**
```
go test -bench . -benchtime 10s ./example/get_keys_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8         	28369766	       635.2 ns/op
BenchmarkGetKeys/Specific-8        	16090556	       685.7 ns/op
PASS
ok  	command-line-arguments	30.245s
```

**Run #4**
```
go test -bench . -benchtime 10s ./example/get_keys_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8         	26351716	       611.3 ns/op
BenchmarkGetKeys/Specific-8        	21451940	       686.9 ns/op
PASS
ok  	command-line-arguments	31.878s
```

**Run #5**
```
go test -bench . -benchtime 10s ./example/get_keys_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8         	25701085	       647.8 ns/op
BenchmarkGetKeys/Specific-8        	18048080	       599.9 ns/op
PASS
ok  	command-line-arguments	28.627s
```
</details>

## Error Handing

### Description

In this case, I imitate getting data for creating of Data instance from
an external source.

As a first action, it gets a slice of strings that are converted to a value for
the ID field. Then it uses the slice and the ID for getting value for the Field1
value. The slice and value of Field1 are used for getting Field2's value.
Field3's value is got from Field1 and Field2 values.

The following struct plays a role of the data source.
```golang
type S1 struct{}

func (*S1) GetPartsID(fail bool) ([]string, error) {
	if fail {
		return nil, fmt.Errorf("fail")
	}

	return []string{"1qaz", "2wsx"}, nil
}

func (*S1) JoinParts(ps []string) (string, error) {
	return "1qaz-2wsx", nil
}

func (*S1) GetField1([]string, string) (string, error) {
	return "3edc", nil
}

func (*S1) GetField2([]string, string) (float64, error) {
	return 0.34, nil
}

func (*S1) GetField3(string, float64) (uint64, error) {
	return 24, nil
}
```

The `NewData` function creates instance of `Data`.
```golang
func NewData(ID string, Field1 string, Field2 float64, Field3 uint64) *Data {
	return &Data{
		ID:     ID,
		Field1: Field1,
		Field2: Field2,
		Field3: Field3,
	}
}
```

The first benchmarked code is simple if-guard approach.
```golang
func DataBuildedWithIfGuard(firstFail bool) (*example.Data, error) {
	s := &example.S1{}

	parts, err := s.GetPartsID(firstFail)
	if err != nil {
		return nil, err
	}

	id, err := s.JoinParts(parts)
	if err != nil {
		return nil, err
	}

	field1, err := s.GetField1(parts, id)
	if err != nil {
		return nil, err
	}

	field2, err := s.GetField2(parts, field1)
	if err != nil {
		return nil, err
	}

	field3, err := s.GetField3(field1, field2)
	if err != nil {
		return nil, err
	}

	return example.NewData(id, field1, field2, field3), nil
}
```

The next variant is a structure that wraps intermediates results.
```golang
type DataConstrutor struct {
	S1     *S1
	Fail   bool
	parts  []string
	id     string
	field1 string
	field2 float64
	field3 uint64
}

func (dc *DataConstrutor) GetPartsID() (err error) {
	dc.parts, err = dc.S1.GetPartsID(dc.Fail)
	return
}

func (dc *DataConstrutor) JoinParts() (err error) {
	dc.id, err = dc.S1.JoinParts(dc.parts)
	return
}

func (dc *DataConstrutor) GetField1() (err error) {
	dc.field1, err = dc.S1.GetField1(dc.parts, dc.id)
	return
}

func (dc *DataConstrutor) GetField2() (err error) {
	dc.field2, err = dc.S1.GetField2(dc.parts, dc.field1)
	return
}

func (dc *DataConstrutor) GetField3() (err error) {
	dc.field3, err = dc.S1.GetField3(dc.field1, dc.field2)
	return
}

func (dc *DataConstrutor) GetData() *Data {
	return NewData(dc.id, dc.field1, dc.field2, dc.field3)
}

func (*DataConstrutor) Chain(ops ...func() error) error {
	for _, op := range ops {
		if err := op(); err != nil {
			return err
		}
	}
	return nil
}
```

It is used in the following way.
```golang
func DataBuildedWithChain(firstFail bool) (*example.Data, error) {
	dc := &example.DataConstrutor{
		S1:   &example.S1{},
		Fail: firstFail,
	}

	err := dc.Chain(
		dc.GetPartsID,
		dc.JoinParts,
		dc.GetField1,
		dc.GetField2,
		dc.GetField3,
	)
	if err != nil {
		return nil, err
	}

	return dc.GetData(), nil
}
```

The last variant works with `Result` interface. It is generic interface that
defined as:
```golang
type Result[T any] interface {
	Get() T
	Err() error

	// Both returns default value for T if result is empty
	Both() (T, error)
}
```

An instance of this interface could be either `OK[T]` or `Err[T]` structures.
Their definition is:
```golang
type OK[T any] struct {
	Value T
}

type Err[T any] struct {
	Value error
}
```

And we need functions which let us to manipulate a value inside `Result`.
Here they are:
```golang
func Map[A, B any](op func(A) B, r Result[A]) Result[B]

func Map2[T1, T2, R any](op func(T1, T2) R, r1 Result[T1], r2 Result[T2]) Result[R]

func Map3[T1, T2, T3, R any](op func(T1, T2, T3) R, r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R]

func Map4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R]

func Map5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R]

func FlatMap[A, B any](op func(A) Result[B], r Result[A]) Result[B]

func FlatMap2[T1, T2, R any](op func(T1, T2) Result[R], r1 Result[T1], r2 Result[T2]) Result[R]

func FlatMap3[T1, T2, T3, R any](op func(T1, T2, T3) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R]

func FlatMap4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R]

func FlatMap5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R]
```

It looks like we need a lot of auxiliary code, but we need to write it only
once. Then we can use it everywhere.

OK, here there is how we can use `Result`.
```golang
func DataBuildedWithResult(firstFail bool) (*example.Data, error) {
	s := &example.S2{}

	parts := s.GetPartsID(firstFail)
	id := result.FlatMap(s.JoinParts, parts)
	field1 := result.FlatMap2(s.GetField1, parts, id)
	field2 := result.FlatMap2(s.GetField2, parts, field1)
	field3 := result.FlatMap2(s.GetField3, field1, field2)

	return result.Map4(example.NewData, id, field1, field2, field3).Both()
}
```

Here I use `S2`, it is a modified version of `S1` where all methods return
`Result` instead of value and error.
```golang
type S2 struct{}

func (*S2) GetPartsID(fail bool) result.Result[[]string] {
	if fail {
		return result.FromErr[[]string](fmt.Errorf("fail"))
	}

	return result.FromValue([]string{"1qaz", "2wsx"})
}

func (*S2) JoinParts(ps []string) result.Result[string] {
	return result.FromValue("1qaz-2wsx")
}

func (*S2) GetField1([]string, string) result.Result[string] {
	return result.FromValue("3edc")
}

func (*S2) GetField2([]string, string) result.Result[float64] {
	return result.FromValue(0.34)
}

func (*S2) GetField3(string, float64) result.Result[uint64] {
	return result.FromValue[uint64](24)
}
```

### Benchmars results

Slice of strings
|         Benchmark name             |    Run #1       |    Run #2       |    Run #3       |    Run #4       |    Run #5       |
| ---------------------------------- | --------------- | --------------- | --------------- | --------------- | --------------- |
| BenchmarkResultSuccess/IfGuard-8   | 0.0000009 ns/op | 0.0000006 ns/op | 0.0000006 ns/op | 0.0000006 ns/op | 0.0000009 ns/op |
| BenchmarkResultSuccess/Chain-8     | 0.0000012 ns/op | 0.0000014 ns/op | 0.0000012 ns/op | 0.0000015 ns/op | 0.0000013 ns/op |
| BenchmarkResultSuccess/Result-8    | 0.0000032 ns/op | 0.0000032 ns/op | 0.0000031 ns/op | 0.0000029 ns/op | 0.0000027 ns/op |
| BenchmarkResultFirstFail/IfGuard-8 | 0.0000028 ns/op | 0.0000032 ns/op | 0.0000029 ns/op | 0.0000026 ns/op | 0.0000025 ns/op |
| BenchmarkResultFirstFail/Chain-8   | 0.0000034 ns/op | 0.0000033 ns/op | 0.0000030 ns/op | 0.0000027 ns/op | 0.0000023 ns/op |
| BenchmarkResultFirstFail/Result-8  | 0.0000042 ns/op | 0.0000035 ns/op | 0.0000032 ns/op | 0.0000036 ns/op | 0.0000042 ns/op |

<details>
<summary>Raw benchmarks results</summary>

**Run #1**
```
go test -bench . ./example/result_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkResultSuccess/IfGuard-8         	1000000000	         0.0000009 ns/op
BenchmarkResultSuccess/Chain-8           	1000000000	         0.0000012 ns/op
BenchmarkResultSuccess/Result-8          	1000000000	         0.0000032 ns/op
BenchmarkResultFirstFail/IfGuard-8       	1000000000	         0.0000028 ns/op
BenchmarkResultFirstFail/Chain-8         	1000000000	         0.0000034 ns/op
BenchmarkResultFirstFail/Result-8        	1000000000	         0.0000042 ns/op
PASS
ok  	command-line-arguments	0.010s
```

**Run #2**
```
go test -bench . ./example/result_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkResultSuccess/IfGuard-8         	1000000000	         0.0000006 ns/op
BenchmarkResultSuccess/Chain-8           	1000000000	         0.0000014 ns/op
BenchmarkResultSuccess/Result-8          	1000000000	         0.0000032 ns/op
BenchmarkResultFirstFail/IfGuard-8       	1000000000	         0.0000032 ns/op
BenchmarkResultFirstFail/Chain-8         	1000000000	         0.0000033 ns/op
BenchmarkResultFirstFail/Result-8        	1000000000	         0.0000035 ns/op
PASS
ok  	command-line-arguments	0.009s
```

**Run #3**
```
go test -bench . ./example/result_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkResultSuccess/IfGuard-8         	1000000000	         0.0000006 ns/op
BenchmarkResultSuccess/Chain-8           	1000000000	         0.0000012 ns/op
BenchmarkResultSuccess/Result-8          	1000000000	         0.0000031 ns/op
BenchmarkResultFirstFail/IfGuard-8       	1000000000	         0.0000029 ns/op
BenchmarkResultFirstFail/Chain-8         	1000000000	         0.0000030 ns/op
BenchmarkResultFirstFail/Result-8        	1000000000	         0.0000032 ns/op
PASS
ok  	command-line-arguments	0.009s
```

**Run #4**
```
go test -bench . ./example/result_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkResultSuccess/IfGuard-8         	1000000000	         0.0000006 ns/op
BenchmarkResultSuccess/Chain-8           	1000000000	         0.0000015 ns/op
BenchmarkResultSuccess/Result-8          	1000000000	         0.0000029 ns/op
BenchmarkResultFirstFail/IfGuard-8       	1000000000	         0.0000026 ns/op
BenchmarkResultFirstFail/Chain-8         	1000000000	         0.0000027 ns/op
BenchmarkResultFirstFail/Result-8        	1000000000	         0.0000036 ns/op
PASS
ok  	command-line-arguments	0.009s
```

**Run #5**
```
go test -bench . ./example/result_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkResultSuccess/IfGuard-8         	1000000000	         0.0000009 ns/op
BenchmarkResultSuccess/Chain-8           	1000000000	         0.0000013 ns/op
BenchmarkResultSuccess/Result-8          	1000000000	         0.0000027 ns/op
BenchmarkResultFirstFail/IfGuard-8       	1000000000	         0.0000025 ns/op
BenchmarkResultFirstFail/Chain-8         	1000000000	         0.0000023 ns/op
BenchmarkResultFirstFail/Result-8        	1000000000	         0.0000042 ns/op
PASS
ok  	command-line-arguments	0.009s
```
</details>
