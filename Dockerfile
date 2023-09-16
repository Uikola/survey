FROM golang:1.20-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ./ ./
RUN go mod download

RUN go build -o ./bin/survey cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/survey /
COPY internal/config/envs/dev.env internal/config/envs/dev.env

CMD ["/survey"]