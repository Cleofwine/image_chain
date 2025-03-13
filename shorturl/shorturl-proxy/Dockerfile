FROM golang:1.23 as stage0
RUN go env -w GOPROXY='https://goproxy.cn,direct'
ADD ./ /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shorturl-proxy .


FROM alpine:latest as stage1
WORKDIR /app/

ADD ./curl_x86_64_glibc_8.11.1.tar.xz /app/
HEALTHCHECK --interval=5s --retries=3 \
CMD ./curl -f http://localhost:9998/health || exit 1

ADD ./dev.config.yaml /app/config.yaml
COPY --from=stage0 /src/shorturl-proxy ./

ENTRYPOINT [ "./shorturl-proxy" ]
CMD [ "--config=config.yaml" ]