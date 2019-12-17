FROM golang:1.12.9-alpine3.10 as builder
COPY web.go .
RUN go build -o /web .

FROM alpine:3.10
RUN apk update
CMD ["./web"]
COPY --from=builder /web .