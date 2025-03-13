FROM golang:1.23 as stage0
RUN go env -w GOPROXY='https://goproxy.cn,direct'
ADD ./ /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shorturl-crontab .


FROM alpine:latest as stage1
WORKDIR /app
ADD ./dev.config.yaml /app/config.yaml
COPY --from=stage0 /src/shorturl-crontab ./

ENTRYPOINT [ "./shorturl-crontab" ]
CMD [ "--config=config.yaml" ]