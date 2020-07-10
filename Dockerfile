FROM golang:alpine AS builder

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk add -U --no-cache ca-certificates tzdata && \
    \
    adduser -s /bin/true -u 1000 -D -h /app app && \
    sed -i -r "/^(app|root)/!d" /etc/group /etc/passwd && \
    sed -i -r 's#^(.*):[^:]*$#\1:/sbin/nologin#' /etc/passwd

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s -extldflags "-static"' -o /app/bin .


FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/group /etc/shadow /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/bin /app

USER app

ENTRYPOINT ["/app"]
