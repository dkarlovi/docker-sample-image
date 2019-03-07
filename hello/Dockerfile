# image building the executable
FROM golang:alpine AS builder

COPY hello.go /go/src/hello/
RUN cd /go/src/hello/ && go get -d -v ./... && go install -ldflags '-s' -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/hello /usr/bin/
ENTRYPOINT ["hello"]
