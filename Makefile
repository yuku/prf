VERBOSE_FLAG = $(if $(VERBOSE),-v)

build: deps
	go build $(VERBOSE_FLAG)

deps:
	go get -d $(VERBOSE_FLAG)

install: deps
	go install $(VERBOSE_FLAG)

.PHONY: build deps install
