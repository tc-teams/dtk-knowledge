FROM golang:1.13 as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go get github.com/gorilla/mux && go build -o main .

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main"]