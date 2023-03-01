.PHONY: all prog docker readme deps

PROG=toobeci

all: ${PROG} test

deps:
	go get .

${PROG}: *.go
	go build .

install: ${PROG}
	go install .

test: ${PROG}
	go test -v

docker:
	docker build -t ${PROG} .

doc:
	python updatereadme.py
