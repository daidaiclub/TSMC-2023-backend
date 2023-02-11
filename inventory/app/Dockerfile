FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY main.go /app/main.go
RUN go mod init business
RUN go mod tidy
RUN go build -o main .
EXPOSE 8200
ENTRYPOINT "/app/main"