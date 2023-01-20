FROM golang:alpine

WORKDIR /app/music-service

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o music-service ./cmd/app/main.go

CMD ["./music-service"]