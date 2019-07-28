FROM golang:1.12-alpine

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/product-service

COPY . .

RUN go mod init product-service
WORKDIR $GOPATH/src/product-service/cmd

RUN GOOS=linux go build -o product-service

ENTRYPOINT ["./product-service"]

EXPOSE 3001