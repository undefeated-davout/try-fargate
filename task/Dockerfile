# 本番用の中間ステージ
FROM golang:1.16.3 as builder
WORKDIR /go/task
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# 本番環境用
FROM alpine:latest as prod
WORKDIR /task
COPY --from=builder /go/task/ .
RUN pwd
CMD ["./task"]

# 開発環境用
FROM golang:1.16.3 as dev
WORKDIR /go/src/task
COPY . .
CMD ["go", "run", "/go/src/task/main.go"]
