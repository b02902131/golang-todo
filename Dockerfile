# 使用 Go 官方鏡像作為構建階段的基礎鏡像
FROM golang:1.22 as builder

# 設置工作目錄
WORKDIR /app

# 安裝 gcc 和其他必要的庫
RUN apt-get update && apt-get install -y gcc libc6-dev

# 複製 go.mod 和 go.sum 文件，並下載依賴
COPY todo-backend/go.* ./
RUN go mod download

# 複製應用的全部代碼到容器中
COPY todo-backend/ ./

# 構建應用程序，啟用 cgo
RUN CGO_ENABLED=1 GOOS=linux go build -v -o server
# 使用更新的 Ubuntu 版本作為生產階段的基礎鏡像
FROM ubuntu:latest
WORKDIR /app

# 安裝運行時所需的庫
RUN apt-get update && apt-get install -y libsqlite3-0

# 從構建階段拷貝編譯好的應用到生產鏡像中
COPY --from=builder /app/server /app/server

# 設置容器啟動時運行應用
CMD ["/app/server"]
