# Running with docker-compose
### 1. Set up local env variable
````
export GO111MODULE=on
export GOPATH=$HOME/go
export PATH="$PATH:$GOPATH/bin"
````

### 2. Recover sub-projects
````
git submodule update --init --recursive
````
then set up env file of each service

### 3. Run Containers
````
cp .env.example .env
docker-compose up --build
````

### 4. Create database and user
reference: ./mysql/docker-entrypoint-initdb.d/createdb.sql.example

## Directories
name  | description 
---- | --- 
p005api |  web api service
p010user | User service
p012dict | Dictionary service (字典)
p013keep | Keep service (單字卡)
pb | protobuf of services
kubernetes | k8s deployment sample for production
mysql | used in docker-compose 
rabbitmq | used in docker-compose

