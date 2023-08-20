#BUILD
FROM golang:1.21.0-alpine3.18 AS builder
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o main main.go

#RUN
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /build/main .
COPY .env .
COPY template template
COPY static static

EXPOSE 8000
CMD [ "/app/main" ]