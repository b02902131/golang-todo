# golang-todo

> learning golang by build-up a todo-application

## RESTful API

- GET /todos
- POST /todos
- GET /todos/{id}
- PUT /todos/{id}
- DELETE /todos/{id}

## SCHEME

- Todo
  - ID
  - Title: string
  - Description: string
  - Completed: bool

## Packages

- Router: gorilla/mux

## Docker

### Run local

```
export DATABASE_URL="postgres://golangtodo:golangtodo@localhost:5432/golangtodo?sslmode=disable" 
go run .
```

```
# build docker
docker build -t myapp .

# run docker
docker run -p 8080:8080 -e DATABASE_URL="postgres://golangtodo:golangtodo@host.docker.internal:5432/golangtodo?sslmode=disable" myapp

# 用 host.docker.internal 代替 localhost Docker 才能指向宿主機
```

### Run on GCP(GKE)

```
# 設置 GCP 項目：
gcloud config set project [YOUR_PROJECT_ID]

# 配置 Docker 以使用 gcloud 作為憑證助手：
gcloud auth configure-docker

# 構建並推送鏡像到 GCR：
docker build -t gcr.io/[YOUR_PROJECT_ID]/myapp:v1 .
docker push gcr.io/[YOUR_PROJECT_ID]/myapp:v1
```

### Run on AWS

```
# build docker
docker build -t myapp . 

# tag & push docker
docker tag myapp:latest your-aws-account-id.dkr.ecr.your-region.amazonaws.com/myapp:latest
# 替換掉 `your-aws-account-id` 與 `your-region`
```

```
# Test RDS DB with local server
export DATABASE_URL="postgres://postgres:Abcd1234@database-3.cluster-cdqgyymqumh0.ap-southeast-2.rds.amazonaws.com:5432/golangtodo"
go run .
```