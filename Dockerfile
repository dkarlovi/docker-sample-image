# image building the executable
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# image holding the exeutable
FROM scratch
COPY --from=builder /go/bin/app /usr/bin/
CMD ["app"]
