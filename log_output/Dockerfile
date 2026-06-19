FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main .

# Run the binary in a clean image
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

CMD [ "./main" ]