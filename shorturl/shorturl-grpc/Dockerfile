FROM golang:1.23 as stage0
RUN go env -w GOPROXY='https://goproxy.cn,direct'
ADD ./ /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shorturl-grpc ./shorturl-server


FROM alpine:latest as stage1
WORKDIR /app/
ADD ./grpc_health_probe-linux-amd64 /app/grpc_health_probe
RUN chmod +x /app/grpc_health_probe
ADD ./dev.config.yaml /app/config.yaml
COPY --from=stage0 /src/shorturl-grpc ./

HEALTHCHECK --interval=5s --retries=3 \
CMD ./grpc_health_probe -addr=:50059

ENTRYPOINT [ "./shorturl-grpc" ]
CMD [ "--config=config.yaml" ]