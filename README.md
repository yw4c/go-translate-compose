


## Services
name  | description
---- | --- 
p005api |  restful api 請求、回應
p010user | 用戶服務

## ports
port  | description
---- | --- 
6001-6999 | GRPC testing 
80 | p005api
8080 | Adminer - mysql cli
8081 | Kibana - ES UI
8500 | Consul UI
9200 | Elasticsearch
9300 | Elasticsearch


## Deploy
````
# migrate -path ./migration -database mysql://root:1234@tcp\(127.0.0.1:3306\)/translate up
````

## Protobuf
###  生成 pb 
````
# ./pb
protoc --go_out=plugins=grpc:. *.proto
````

### 同步 pb 庫
* docker 運行使用 volume 同步
* 本地 coding 使用連結
````
# e.g. in ./p010user/src
ln -s ../../pb ./
````
