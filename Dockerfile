FROM golang:1.13.1

# Enable go modules support in gopath
ENV GO111MODULE=on

WORKDIR /go/src/app

# Copy modules list which required by the project
COPY go.mod .
COPY go.sum .

# Download modules
RUN go mod download

COPY . .

RUN go build -o $GOPATH/bin/app