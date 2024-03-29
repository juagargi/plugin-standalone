.PHONY: all build clean update

all: build

build:
	@go build -buildmode=plugin .

clean:
	@rm -f plugin-standalone.so

update:
	@echo "Updating main module ..."
	@go get -u github.com/juagargi/plugin-test@HEAD
	@echo "Done. This might have modified the source code."
