.PHONY:=all
.DEFAULT_GOAL:=run


run:
	@./scripts/run.sh

build: 
	@./scripts/build.sh

test:
	@./scripts/test.sh

cov:
	@./scripts/coverage.sh 

covHtml: cov
	go tool cover -html=reports/coverage.out

