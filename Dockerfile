FROM ginuerzh/gost as gost

FROM ngrok/ngrok:alpine as ngrok

FROM golang:1.19-alpine as builder
WORKDIR /src
COPY main.go go.mod go.sum ./
RUN go mod download && go build -o /app/run ./main.go

FROM alpine

WORKDIR /app

COPY --from=gost /bin/gost /usr/local/bin/gost
COPY --from=ngrok /bin/ngrok /usr/local/bin/ngrok
COPY --from=builder /app/run .

RUN chmod +x /app/run /usr/local/bin/gost /usr/local/bin/ngrok

CMD ["/app/run"]