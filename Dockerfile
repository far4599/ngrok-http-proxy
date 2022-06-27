FROM ginuerzh/gost as gost

FROM ngrok/ngrok:alpine as ngrok

FROM golang:1.18-alpine as builder
WORKDIR /src
COPY main.go go.mod go.sum ./
RUN go mod download && go build -o /app/run ./main.go

FROM alpine

WORKDIR /app

COPY --from=gost /bin/gost .
COPY --from=ngrok /bin/ngrok .
COPY --from=builder /app/run .

RUN chmod +x /app/*

CMD ["/app/run"]