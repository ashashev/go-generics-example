.PHONY: bench bench_long

bench:
	go test -bench . ./foo

bench_long:
	go test -bench . -benchtime 10s ./foo
