FROM golang:alpine

WORKDIR /app/youtube-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o youtube-service ./cmd/app/main.go

CMD ["./youtube-service"]