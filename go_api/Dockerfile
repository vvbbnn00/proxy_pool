FROM golang:1.21 AS builder

WORKDIR /app
COPY src/go.mod ./
RUN go mod download
COPY src ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./main/

FROM alpine:latest

RUN adduser -D appuser

WORKDIR /home/appuser
COPY --from=builder /app/server .
USER appuser

EXPOSE 8080

CMD ["./server"]