FROM golang:1.13-alpine as builder
COPY . /code
WORKDIR /code
RUN go build -o server

FROM alpine:latest
WORKDIR /code
COPY --from=builder /code/server .
EXPOSE 8080
CMD ["./server"]
