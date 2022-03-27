# Results

## Slice Contains

### Run #1
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

### Run #2
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

### Run #3
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

### Run #4
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

### Run #5
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

## Get Keys

### Run #1
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

### Run #2
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

### Run #3
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

### Run #4
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

### Run #5
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

## Error Handing

### Run #1
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

### Run #2
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

### Run #3
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

### Run #4
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

### Run #5
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

