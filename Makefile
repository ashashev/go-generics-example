.PHONY: bench bench_long

bench:
	go test -bench . ./example

bench_long:
	go test -bench . -benchtime 10s ./example
