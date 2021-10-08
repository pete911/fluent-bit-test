FROM golang:1.17.2-alpine AS build
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/app

COPY . .
RUN go build -o /bin/fluent-bit-test


FROM alpine:3.14.2
RUN apk add --no-cache ca-certificates

COPY --from=build /bin/fluent-bit-test /usr/local/bin/fluent-bit-test
CMD ["fluent-bit-test"]
