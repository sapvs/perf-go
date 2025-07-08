.DEFAULT_GOAL := all
gen: 
	docker run --rm -v ${PWD}:/proto protoc:0.1 -I=/proto --go_out=/proto /proto/proto/user.proto
all: gen
	# Runs ALL benchmark
	go test -benchmem -coverprofile=coverage  -bench . -race