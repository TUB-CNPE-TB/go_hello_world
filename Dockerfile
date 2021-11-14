# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
#COPY *.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /docker-go-hello
EXPOSE 9090
CMD [ "/docker-go-hello" ]