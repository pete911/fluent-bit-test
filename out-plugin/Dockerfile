ARG name=fluent-bit-test

FROM golang:1.17 AS builder

ARG name
WORKDIR /go/src/github.com/pete911/$name

COPY . /go/src/github.com/pete911/$name/

RUN go get github.com/fluent/fluent-bit-go/output
RUN go build -buildmode=c-shared -o out_$name.so

FROM fluent/fluent-bit:1.8.9

ARG name
COPY --from=builder /go/src/github.com/pete911/$name/out_$name.so /fluent-bit/bin/out_$name.so
COPY plugins.conf /fluent-bit/etc/

CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/bin/out_fluent-bit-test.so"]
