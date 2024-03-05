.PHONY: all build clean

all: build

build:
	@go build -buildmode=plugin .

clean:
	@rm -f plugin-standalone.so
