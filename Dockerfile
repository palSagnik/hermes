FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download
COPY ./backend/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir ./template
COPY --from=builder /app/main .
COPY --from=builder /app/template/mail.html ./template/mail.html
COPY --from=builder /app/.env .

EXPOSE 9000
EXPOSE 587

CMD ["./main"]

