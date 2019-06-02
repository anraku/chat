FROM golang:1.11.10-alpine3.9 AS builder

RUN apk add git

WORKDIR /go/src/github.com/anraku/chat
ENV GO111MODULE=on
RUN addgroup -g 10001 -S app \
    && adduser -u 10001 -G app -S app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER app

ENTRYPOINT ["/go/bin/app"]
