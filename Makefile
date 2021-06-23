all: test

test: ## Test
	@go version
	go test -cpu 1,2,4 -count  5 -benchmem -bench .

#stat: ## Benchmark statistics
#	benchstat 1.log 2.log

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all test help
