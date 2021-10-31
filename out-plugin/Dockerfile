FROM golang:1.17 AS builder

WORKDIR /go/src/github.com/pete911/fluent-bit-test

COPY Makefile out.go go.mod go.sum /go/src/github.com/pete911/fluent-bit-test/

ENV SOURCE docker
RUN go get github.com/fluent/fluent-bit-go/output

# Not using default value here due to this: https://github.com/docker/buildx/issues/510
ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}
RUN echo "Building for ${TARGETPLATFORM} architecture"
RUN make ${TARGETPLATFORM}

FROM fluent/fluent-bit:1.8.9

COPY --from=builder /go/src/github.com/pete911/fluent-bit-test/out_fluent-bit-test-*.so /fluent-bit/bin/out_fluent-bit-test.so
COPY plugins.conf /fluent-bit/etc/

CMD ["/fluent-bit/bin/fluent-bit", "-c", "/fluent-bit/etc/fluent-bit.conf", "-e", "/fluent-bit/bin/out_fluent-bit-test.so"]
