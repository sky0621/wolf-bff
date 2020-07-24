FROM golang:1.14 as builder

ENV PROJECT github.com/sky0621/wolf-bff
WORKDIR $GOPATH/src/$PROJECT/src

ENV GO111MODULE=on
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download

COPY ./src .
RUN go build -o /wolf ./cmd

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /wolf /wolf

EXPOSE 8080
ENTRYPOINT ["/wolf"]
