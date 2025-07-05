.DEFAULT_GOAL := all
gen: 
	protoc -I=. --go_out=. --go_opt=paths=source_relative ./user.proto
all:
	# Runs ALL benchmark
	go test -benchmem -coverprofile=coverage  -bench . -race
