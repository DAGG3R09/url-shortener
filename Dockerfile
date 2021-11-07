FROM golang:1.17.2-alpine as builder

ENV GOFLAG="-mod=vendor"

RUN apk --no-cache add make bash

WORKDIR /shortener
COPY . /shortener/.

ENV CGO_ENABLED=0

RUN make clean build

FROM alpine:3.14.2

COPY --from=builder /shortener/build/urlshortener /urlshortener

EXPOSE 8080

CMD ["/urlshortener"]

