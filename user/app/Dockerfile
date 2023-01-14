FROM golang:alpine

WORKDIR /app/vgm-server

COPY go.mod .
COPY go.sum .
ENV GOPATH=/
RUN go mod download

#build appliction
COPY . .
RUN go build -o vgm-server ./cmd/app/main.go

CMD ["./vgm-server"]