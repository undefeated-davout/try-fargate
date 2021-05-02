# 本番用の中間ステージ
FROM golang:1.16.3 as builder
WORKDIR /go/api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# 本番環境用
FROM alpine:latest as prod
EXPOSE 8080
WORKDIR /api
COPY --from=builder /go/api/ .
RUN pwd
CMD ["./web-server"]

# 開発環境用（ホットリロード環境）
FROM golang:1.16.3 as dev
EXPOSE 8080
WORKDIR /go/src/web-server
COPY . .
ENV GO111MODULE=on
RUN go get github.com/pilu/fresh
CMD ["fresh"]
