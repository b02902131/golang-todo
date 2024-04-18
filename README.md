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
# build docker
docker build -t myapp .

# run docker
docker run -p 4000:8080 -e DB_PATH=/app/todo-backend/todo.db -v /Users/bigyo/Documents/project/golang-todo/todo-backend/todo.db:/app/todo-backend/todo.db myapp
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
