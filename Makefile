GO_CMD=go
OS=$(shell uname)

all: install

install:
	${GO_CMD} build -o ./build/calc ./cmd/main.go

test:
	${GO_CMD} test -bench=. -v ./tests/*_test.go

clean:
	cd build && rm calc
	cd log && rm *
