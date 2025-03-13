# shorturl-crontab

## 构建镜像
```bash
docker build -t shorturl-crontab:0.1.0 .
```

## swarm集群
```bash
docker config create shorturl-crontab-conf dev.config.yaml
docker service create --name shorturl-crontab --replicas 1 --config src=shorturl-crontab-conf,target=/app/config.yaml  shorturl-crontab:0.1.0
```