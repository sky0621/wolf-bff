# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.14 as builder

ENV PROJECT github.com/sky0621/wolf-bff
WORKDIR $GOPATH/src/$PROJECT

ENV GO111MODULE=on
COPY go.* ./
RUN go mod download

COPY ./src/ ./
RUN go build -o /wolf ./cmd

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /wolf /wolf

# Run the web service on container startup.
EXPOSE 8765
ENTRYPOINT ["/wolf"]
