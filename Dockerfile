FROM golang:1.21.0-alpine as builder
WORKDIR /go/src
COPY . .
RUN go build -o vault-audit .

FROM alpine:3.18.3
WORKDIR /bin
COPY --from=builder /go/src .
USER nobody
EXPOSE 8080
CMD ["vault-audit"]
