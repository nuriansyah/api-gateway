FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o api-gateway

EXPOSE 8080

CMD ./api-gateway