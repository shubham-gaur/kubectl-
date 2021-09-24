export GOBIN=${PWD}/bin
export CGO_ENABLED=0

.PHONY: all
all: install

build:
	go build -o bin/kubectl++  cmd/main.go

install:
	# Generating binary at $(PWD)/bin
	# Installing kubectl++ ...
	# Generated binary can be found in bin folder
	go install cmd/main.go
	@mv bin/main bin/kubectl++
	# Kubectl++ installed successfully in bin folder!

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: RESET
RESET: clean

test:
	# Executing tests...
	go test -v -timeout 30s -run TestExecKubectlCmd github.com/shubham-gaur/kubectl++/helper
	# All test completed !

gen-doc:
	@mkdir -p docs; golds -wdpkgs-listing=solo -source-code-reading=plain -nouses -allow-network-connection -gen -dir=docs/ cmd/main.go
	# Docs generation completed!
