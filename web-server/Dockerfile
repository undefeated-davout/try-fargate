FROM golang:1.12 as builder
WORKDIR /go/api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest as prod
EXPOSE 80
WORKDIR /api
COPY --from=builder /go/api/ .
RUN pwd
CMD ["./webapp-api"]

FROM golang:1.12 as dev
EXPOSE 80
WORKDIR /go/src/whoisapp
COPY . .
ENV GO111MODULE=on
RUN go get github.com/pilu/fresh
CMD ["fresh"]
