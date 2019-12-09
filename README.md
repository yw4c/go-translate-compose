# Running with docker-compose
### Set up local env variable
````
export GO111MODULE=on
export GOPATH=$HOME/go
export PATH="$PATH:$GOPATH/bin"
````

### recover sub-projects
````
git submodule update --init --recursive
````

### run docker
````
cp .env.example .env
docker-compose up --build
````

## Running with kubernetes


## Services
name  | description
---- | --- 
p005api |  restful api 
p010user | User 用戶
p012dict | Dictionary 字典
p013keep | Keep 單字卡


## Occupied ports
port  | description
---- | --- 
6001-6999 | GRPC testing 
81 | p005api
8080 | Adminer - mysql cli
8081 | Kibana - ES UI
8500 | Consul UI
9200 | Elasticsearch
9300 | Elasticsearch




## Protobuf
###  生成 pb 
````
sh ./pb/gen.sh
````

