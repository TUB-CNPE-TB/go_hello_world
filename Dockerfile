# syntax=docker/dockerfile:1
FROM golang:alpine as stage
WORKDIR /app
COPY go.mod ./
#COPY *.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /helloworld

FROM scratch
COPY --from=stage /helloworld /helloworld
EXPOSE 80
CMD [ "/helloworld" ]