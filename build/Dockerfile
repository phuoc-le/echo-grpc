FROM golang:1.16-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/grpc-echo

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go Server GRPC
RUN go build -o ./out/grpc-echo-server ./server/main.go
# Run the Go App
RUN go build -o ./out/grpc-echo-client ./client/main.go


# This container exposes port 8080 to the outside world
EXPOSE 8080, 50055

# Run the binary program produced by `go install`
CMD ["./out/grpc-echo-server", "./out/grpc-echo-client"]
