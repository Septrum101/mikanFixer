FROM golang:alpine AS builder

ARG VERSION=dev

WORKDIR /build
COPY . .

RUN go mod tidy
RUN go build -trimpath -ldflags="-s -w \
    -X 'main.date=$(date -Iseconds)' \
    -X 'main.version=$VERSION' \
    " -v -o mikanFixer

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/mikanFixer /app/mikanFixer
ENTRYPOINT ["/app/mikanFixer"]