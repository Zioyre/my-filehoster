FROM golang:1.23-alpine

RUN apk add --no-cache git
WORKDIR /app

COPY . .
RUN go build -o server .

EXPOSE 48723

CMD ["./server"]
