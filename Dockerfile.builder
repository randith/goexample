# start a golang base image, version 1.8
FROM golang:1.9.2

RUN mkdir -p /go

#copy the source files
COPY . /go/src/github.com/randith/goexample
WORKDIR /go/src/github.com/randith/goexample

# premerge checks
RUN ./premerge.sh

#build the binary with debug information removed
ENV GOOS=linux
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/randith/goexample/cmd/pwhash
RUN go build -a -ldflags '-extldflags "-static"' -o pwhash