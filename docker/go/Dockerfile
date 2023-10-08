FROM --platform=linux/x86_64 golang:1.21-alpine as dev
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk update && \
    apk add curl make bash && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && \
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

