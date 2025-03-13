# mediahub

## 构建镜像
```bash
docker build -t mediahub:0.1.0 .
```

## swarm集群
```bash
docker config create mediahub-conf dev.config.yaml
docker service create --name mediahub -p 9999:9999 --replicas 1 --config src=mediahub-conf,target=/app/config.yaml --network mediahub-net mediahub:0.1.0
```