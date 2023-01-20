FROM golang:alpine

WORKDIR /app/user-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o user-service ./cmd/app/main.go

CMD ["./user-service"]