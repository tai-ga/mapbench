all: deps test

setup:
	go get -u github.com/golang/dep/...

deps: setup
	dep ensure

dep: ## dep ensure
	dep ensure

depup: ## dep ensure -update
	dep ensure -update

test: ## Test
	go test -cpu 1,2,4 -count  5 -benchmem -bench .

#stat: ## Benchmark statistics
#	benchstat 1.log 2.log

clean: ## clean up
	@rm -rf vendor

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all setup deps dep depup test help
