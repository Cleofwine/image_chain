# shorturl-grpc

## 生成grpc存根
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/shorturl.proto
```

## 构建镜像
```bash
docker build -t shorturl-grpc:0.1.0 .
```

## 创建一个网络
```bash
docker network create -d overlay mediahub-net
```

## swarm集群
```bash
docker config create shorturl-grpc-conf dev.config.yaml
docker service create --name shorturl-grpc --replicas 1 --config src=shorturl-grpc-conf,target=/app/config.yaml --network mediahub-net shorturl-grpc:0.1.0
```