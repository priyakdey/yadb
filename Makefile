.PHONY:=help
.DEFAULT_GOAL:=help


help:
	@echo "Usage hints:"
	@echo "	help		Print usage"
	@echo "	build		Builds the binary"
	@echo "	run		Run the binary"

build:
	@go build -o bin/yadb main.go

run: build
	@if [ -z ${FILE} ]; then \
		./bin/yadb; \
	else \
		./bin/yadb ${FILE}; \
	fi
