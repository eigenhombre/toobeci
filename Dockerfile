FROM golang:1.18

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade
RUN apt-get install -y make

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@2022.1
RUN go install -v golang.org/x/lint/golint@latest

COPY . .

RUN make
