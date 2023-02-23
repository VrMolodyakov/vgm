FROM golang:alpine

WORKDIR /app/gateway-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o gateway-service ./cmd/app/main.go

CMD ["./gateway-service"]