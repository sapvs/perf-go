.DEFAULT_GOAL := all
all:
	# Runs ALL benchmark
	go test -benchmem -coverprofile=coverage  -bench . -race