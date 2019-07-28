FROM golang:alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/order-service

COPY cmd/order-service .

ENTRYPOINT ["./order-service"]

EXPOSE 3003