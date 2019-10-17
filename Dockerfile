# Alpine as mininal image
FROM golang:alpine
# Git for go get
RUN apk update && apk add --no-cache git

# Enable go modules support in gopath
ENV GO111MODULE=on

WORKDIR /go/src/app

# Copy modules list which required by the project
COPY go.mod .
COPY go.sum .

# Download modules
RUN go mod download

COPY . .
# Build
RUN go build -o $GOPATH/bin/app
# Set default entrypoint
CMD ["/go/bin/app"]