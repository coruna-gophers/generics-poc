MEM_PROFILE:="tree-mem.out" 
CPU_PROFILE:="tree-cpu.out"

.PHONY: help
help:	### This screen. Keep it first target to be default
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

.PHONY: bench
bench: 	### Run benchmarks and generate cpu and memory profiles
	go test -v ./... -bench=. -benchmem -memprofile ${MEM_PROFILE} -cpuprofile ${CPU_PROFILE}

.PHONY: results-mem
results-mem: ### Show memory profile results
	pprof -http=localhost:8080 ${MEM_PROFILE}


.PHONY: results-cpu
results-cpu: ### Show cpu profile results
	pprof -http=localhost:8080 ${CPU_PROFILE}