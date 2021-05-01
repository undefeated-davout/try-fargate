FROM alpine:latest AS build
RUN apk -U add gcc libc-dev
COPY hello.c /
RUN gcc -static -o hello hello.c

FROM scratch AS release
COPY --from=build /hello /
CMD ["/hello"]
