.DEFAULT_GOAL := all
pbgen: 
	docker run --rm -v ${PWD}:/proto vsaps/pb-go -I=/proto --go_out=/proto /proto/proto/user.proto
all: pbgen
	# Runs ALL benchmark
	go test -benchmem -coverprofile=coverage  -bench . -timeout 0 -race