## Running Project
### 1. Load sub-modules 
````
git submodule ...
````


## Services
name  | description
---- | --- 
p005api |  restful api 
p010user | User 用戶
p011trans | Translation 翻譯
p012dict | Dictionary 字典
p013keep | Keep 單字卡
p014google-svc | Google Cloud Service
p015oxford-svc | Oxford Service
p016detect-lang-com | detect-lang.com Service

## ports
port  | description
---- | --- 
6001-6999 | GRPC testing 
81 | p005api
8080 | Adminer - mysql cli
8081 | Kibana - ES UI
8500 | Consul UI
9200 | Elasticsearch
9300 | Elasticsearch




### Protobuf
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
