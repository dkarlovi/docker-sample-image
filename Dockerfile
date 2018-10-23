# image building the executable
FROM golang:alpine AS builder

COPY entry.go /go/src/entry/
RUN cd /go/src/entry/ && go get -d -v ./... && go install -v ./...

COPY hello.go /go/src/hello/
RUN cd /go/src/hello/ && go get -d -v ./... && go install -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/entry /go/bin/hello /usr/bin/
CMD ["hello"]
