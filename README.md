# Results

```
go test -bench . -benchtime 10s ./foo
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/foo
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8 	23023732	       682.3 ns/op
BenchmarkGetKeys/Specific-8         	21786260	       669.2 ns/op
BenchmarkSliceContains/String_Generic_Doesn't_Has-8         	783723334	        15.37 ns/op
BenchmarkSliceContains/String_Generic_Has-8                 	259270231	        47.10 ns/op
BenchmarkSliceContains/String_Specific_Doesn't_Has-8        	761286849	        16.07 ns/op
BenchmarkSliceContains/String_Specific_Has-8                	253630765	        47.14 ns/op
BenchmarkSliceContains/Int_Generic_Doesn't_Has-8            	1000000000	         8.552 ns/op
BenchmarkSliceContains/Int_Generic_Has-8                    	1000000000	         4.931 ns/op
BenchmarkSliceContains/Int_Specific_Doesn't_Has-8           	1000000000	         8.655 ns/op
BenchmarkSliceContains/Int_Specific_Has-8                   	1000000000	         4.918 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/foo	122.227s
-------------
go test -bench . -benchtime 10s ./foo
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/foo
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8 	27532089	       739.9 ns/op
BenchmarkGetKeys/Specific-8         	22785630	       741.1 ns/op
BenchmarkSliceContains/String_Generic_Doesn't_Has-8         	776895194	        15.55 ns/op
BenchmarkSliceContains/String_Generic_Has-8                 	254971216	        45.63 ns/op
BenchmarkSliceContains/String_Specific_Doesn't_Has-8        	760942696	        16.55 ns/op
BenchmarkSliceContains/String_Specific_Has-8                	249599912	        49.34 ns/op
BenchmarkSliceContains/Int_Generic_Doesn't_Has-8            	1000000000	         8.642 ns/op
BenchmarkSliceContains/Int_Generic_Has-8                    	1000000000	         5.141 ns/op
BenchmarkSliceContains/Int_Specific_Doesn't_Has-8           	1000000000	         8.797 ns/op
BenchmarkSliceContains/Int_Specific_Has-8                   	1000000000	         5.118 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/foo	130.130s
-------------
go test -bench . -benchtime 10s ./foo
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/foo
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8 	28898805	       710.8 ns/op
BenchmarkGetKeys/Specific-8         	23732270	       745.8 ns/op
BenchmarkSliceContains/String_Generic_Doesn't_Has-8         	774048608	        15.43 ns/op
BenchmarkSliceContains/String_Generic_Has-8                 	264248152	        46.01 ns/op
BenchmarkSliceContains/String_Specific_Doesn't_Has-8        	715689662	        16.55 ns/op
BenchmarkSliceContains/String_Specific_Has-8                	257957875	        48.67 ns/op
BenchmarkSliceContains/Int_Generic_Doesn't_Has-8            	1000000000	         8.452 ns/op
BenchmarkSliceContains/Int_Generic_Has-8                    	1000000000	         5.325 ns/op
BenchmarkSliceContains/Int_Specific_Doesn't_Has-8           	1000000000	         8.764 ns/op
BenchmarkSliceContains/Int_Specific_Has-8                   	1000000000	         4.971 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/foo	130.636s
-------------
go test -bench . -benchtime 10s ./foo
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/foo
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8 	27564216	       692.0 ns/op
BenchmarkGetKeys/Specific-8         	28896255	       650.9 ns/op
BenchmarkSliceContains/String_Generic_Doesn't_Has-8         	775978351	        16.28 ns/op
BenchmarkSliceContains/String_Generic_Has-8                 	245883170	        47.97 ns/op
BenchmarkSliceContains/String_Specific_Doesn't_Has-8        	725489311	        16.94 ns/op
BenchmarkSliceContains/String_Specific_Has-8                	255854778	        47.38 ns/op
BenchmarkSliceContains/Int_Generic_Doesn't_Has-8            	1000000000	         8.510 ns/op
BenchmarkSliceContains/Int_Generic_Has-8                    	1000000000	         5.019 ns/op
BenchmarkSliceContains/Int_Specific_Doesn't_Has-8           	1000000000	         9.072 ns/op
BenchmarkSliceContains/Int_Specific_Has-8                   	1000000000	         5.381 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/foo	131.314s
-------------
go test -bench . -benchtime 10s ./foo
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/foo
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGetKeys/Generic-8 	29954083	       704.2 ns/op
BenchmarkGetKeys/Specific-8         	26493284	       751.2 ns/op
BenchmarkSliceContains/String_Generic_Doesn't_Has-8         	785687806	        15.97 ns/op
BenchmarkSliceContains/String_Generic_Has-8                 	264972188	        49.31 ns/op
BenchmarkSliceContains/String_Specific_Doesn't_Has-8        	766004475	        16.34 ns/op
BenchmarkSliceContains/String_Specific_Has-8                	240727125	        47.94 ns/op
BenchmarkSliceContains/Int_Generic_Doesn't_Has-8            	1000000000	         8.440 ns/op
BenchmarkSliceContains/Int_Generic_Has-8                    	1000000000	         4.989 ns/op
BenchmarkSliceContains/Int_Specific_Doesn't_Has-8           	1000000000	         8.843 ns/op
BenchmarkSliceContains/Int_Specific_Has-8                   	1000000000	         4.983 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/foo	134.424s
-------------
```
