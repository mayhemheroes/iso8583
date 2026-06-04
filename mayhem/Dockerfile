FROM golang:1.26 AS builder

WORKDIR /src
COPY . .
RUN GOTOOLCHAIN=auto go test -c -fuzz=FuzzReader ./test/fuzz-reader/ -o /reader-fuzz

FROM ubuntu:22.04
COPY --from=builder /reader-fuzz /
ENTRYPOINT ["/reader-fuzz", "-test.fuzz=FuzzReader"]
