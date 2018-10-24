# image building the executable
FROM golang:alpine AS builder

COPY entry.go /go/src/entry/
RUN cd /go/src/entry/ && go get -d -v ./... && go install -v ./...

COPY env.go /go/src/env/
RUN cd /go/src/env/ && go get -d -v ./... && go install -v ./...

COPY exec.go /go/src/exec/
RUN cd /go/src/exec/ && go get -d -v ./... && go install -v ./...

COPY signal.go /go/src/signal/
RUN cd /go/src/signal/ && go get -d -v ./... && go install -v ./...

COPY sleep.go /go/src/sleep/
RUN cd /go/src/sleep/ && go get -d -v ./... && go install -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/entry /go/bin/env /go/bin/exec /go/bin/signal /go/bin/sleep /usr/bin/
ENTRYPOINT ["exec"]
CMD ["hello"]
