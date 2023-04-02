FROM golang:alpine

WORKDIR /app/email-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o email-service ./cmd/app/main.go

CMD ["./email-service"]