# 使用官方的 Go 基礎鏡像
FROM golang:1.17 as builder

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下載所有依賴
RUN go mod download

# 複製項目文件
COPY . .

# 編譯 Go 程序為二進制文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 使用 alpine 鏡像開始另一個階段
FROM alpine:latest
WORKDIR /root/

# 從構建器鏡像中複製編譯好的應用
COPY --from=builder /app/myapp .

# 運行你的應用
CMD ["./myapp"]
