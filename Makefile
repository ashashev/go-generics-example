.PHONY: bench bench_long

bench:
	go test -bench . ./example

bench_get_keys_long:
	go test -bench . -benchtime 10s ./example/get_keys_test.go

bench_slice_contains_long:
	go test -bench . -benchtime 10s ./example/contains_test.go

bench_result:
	go test -bench . -benchtime 10s ./example/result_test.go
