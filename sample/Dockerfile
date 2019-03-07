# image building the executable
FROM golang:alpine AS builder

COPY hello.go /go/src/hello/
RUN cd /go/src/hello/ && go get -d -v ./... && go install -ldflags '-s' -v ./...

COPY input.go /go/src/input/
RUN cd /go/src/input/ && go get -d -v ./... && go install -ldflags '-s'  -v ./...

COPY signal.go /go/src/signal/
RUN cd /go/src/signal/ && go get -d -v ./... && go install -ldflags '-s' -v ./...

COPY sleep.go /go/src/sleep/
RUN cd /go/src/sleep/ && go get -d -v ./... && go install -ldflags '-s' -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/hello /go/bin/input /go/bin/signal /go/bin/sleep /usr/bin/
CMD ["hello"]
