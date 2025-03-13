# 简介
- 基于微服务架构构建图床服务。通过独立部署的上传服务接收图片，利用腾讯云COS（已进行接口抽象，后续可替换为ceph等持久化方案）进行持久化存储并生成可追踪短链。

# 服务结构
![服务结构](http://cleofwine.icu/p/X "结构图")

# 部署
1. 前置准备
```bash
# 1. mysql
mysql -u 用户名 -p 
source ./shorturl/shorturl-grpc/sql/create_db.sql;
USE shorturl;
source ./shorturl/shorturl-grpc/sql/create_table.sql;
# 2. 确保redis启用
```
2. 构建镜像
```bash
cd ./mediahub
docker build -t mediahub:0.1.0 .
cd ./mediahub-web
docker build -t mediahub-web:0.1.0 .
cd ./shorturl/shorturl-crontab
docker build -t shorturl-crontab:0.1.0 .
cd ./shorturl/shorturl-grpc
docker build -t shorturl-grpc:0.1.0 .
cd ./shorturl/shorturl-proxy
docker build -t shorturl-proxy:0.1.0 .
```
3. 修改配置
```bash
cd ./stack/configs # 这个路径下可以找到全部的配置，修改成自己的对应COS token、redis、mysql连接方式
``` 
4. 统一部署服务
```bash
cd ./stack
docker stack deploy -c compose.yaml imageChain-stack
# 访问
http://<IP>:9997/mediahub-web/
```
5. 统一删除服务
```bash
# 配置不会被删除
docker stack rm imageChain-stack
```