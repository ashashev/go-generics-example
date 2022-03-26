# Results

## Run #1
```
go test -bench . -benchtime 10s ./example
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/example
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	800580374	        14.92 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	298324292	        40.41 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	791018108	        15.22 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	294893266	        40.80 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.452 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.913 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.462 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.915 ns/op
BenchmarkGetKeys/Generic-8                                 	23922086	       690.2 ns/op
BenchmarkGetKeys/Specific-8                                	27178148	       648.4 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/example	123.851s
```

## Run #2
```
go test -bench . -benchtime 10s ./example
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/example
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	795660807	        15.07 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	295378468	        40.68 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	790436913	        15.18 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	293561124	        40.86 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.445 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.913 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.454 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.914 ns/op
BenchmarkGetKeys/Generic-8                                 	26450540	       701.7 ns/op
BenchmarkGetKeys/Specific-8                                	16807503	       615.9 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/example	118.838s
```

## Run #3
```
go test -bench . -benchtime 10s ./example
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/example
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	793174506	        15.13 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	295299771	        42.19 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	774304808	        15.54 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	294594625	        40.84 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.473 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.927 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.451 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.912 ns/op
BenchmarkGetKeys/Generic-8                                 	29269405	       642.1 ns/op
BenchmarkGetKeys/Specific-8                                	26109948	       662.9 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/example	126.310s

```
## Run #4
```
go test -bench . -benchtime 10s ./example
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/example
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	802100079	        15.01 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	294808531	        40.69 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	790537789	        15.20 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	294710605	        40.71 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.457 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.913 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.450 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.918 ns/op
BenchmarkGetKeys/Generic-8                                 	28285144	       645.5 ns/op
BenchmarkGetKeys/Specific-8                                	18511918	       741.8 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/example	121.873s
```

## Run #5
```
go test -bench . -benchtime 10s ./example
goos: linux
goarch: amd64
pkg: github.com/ashashev/go-generics-example/example
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkSliceContainsString/Generic_Doesn't_Has-8         	802601623	        15.05 ns/op
BenchmarkSliceContainsString/Generic_Has-8                 	294817100	        41.97 ns/op
BenchmarkSliceContainsString/Specific_Doesn't_Has-8        	790406701	        15.14 ns/op
BenchmarkSliceContainsString/Specific_Has-8                	281663295	        40.70 ns/op
BenchmarkSliceContainsInt/Generic_Doesn't_Has-8            	1000000000	         8.432 ns/op
BenchmarkSliceContainsInt/Generic_Has-8                    	1000000000	         4.912 ns/op
BenchmarkSliceContainsInt/Specific_Doesn't_Has-8           	1000000000	         8.449 ns/op
BenchmarkSliceContainsInt/Specific_Has-8                   	1000000000	         4.905 ns/op
BenchmarkGetKeys/Generic-8                                 	25260649	       652.5 ns/op
BenchmarkGetKeys/Specific-8                                	17656322	       728.6 ns/op
PASS
ok  	github.com/ashashev/go-generics-example/example	127.595s
```
