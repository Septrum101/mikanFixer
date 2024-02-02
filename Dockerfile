FROM golang:alpine AS builder

ARG VERSION

WORKDIR /build
COPY . .

RUN go mod tidy
RUN go build -trimpath -ldflags="-s -w \
    -X 'github.com/thank243/mikanFixer/version.date=$(date -Iseconds)' \
    -X 'github.com/thank243/mikanFixer/version.version=$VERSION' \
    " -v -o mikanFixer main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/mikanFixer /app/mikanFixer
ENTRYPOINT ["/app/mikanFixer"]