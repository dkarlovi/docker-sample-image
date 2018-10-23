# image building the executable
FROM golang:alpine AS builder

COPY entry.go /go/src/entry/
RUN cd /go/src/entry/ && go get -d -v ./... && go install -v ./...

COPY hello.go /go/src/hello/
RUN cd /go/src/hello/ && go get -d -v ./... && go install -v ./...

COPY signal.go /go/src/signal/
RUN cd /go/src/signal/ && go get -d -v ./... && go install -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/entry /go/bin/hello /go/bin/signal /usr/bin/
CMD ["hello"]
