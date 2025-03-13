# mediahub

## 构建镜像
```bash
docker build -t shorturl-proxy:0.1.0 .
```

## swarm集群
```bash
docker config create shorturl-proxy-conf dev.config.yaml
docker service create --name shorturl-proxy -p 9998:9998 --replicas 1 --config src=shorturl-proxy-conf,target=/app/config.yaml --network mediahub-net shorturl-proxy:0.1.0
```