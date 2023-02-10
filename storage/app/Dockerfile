FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY main.go /app/main.go
RUN go mod init storage
RUN go mod tidy
RUN go build -o main .
EXPOSE 8300
ENTRYPOINT "/app/main"