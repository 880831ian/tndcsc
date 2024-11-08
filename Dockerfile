# 基礎 image
FROM golang:1.22-alpine as builder

# 設定目錄
WORKDIR /app

# 複製程式碼
COPY . .

# 編譯程式碼
RUN go build -o tndcsc

# 使用輕量 image
FROM alpine:latest

# 設定工作目錄
WORKDIR /app

# 複製編譯後的程式碼到最終 image
COPY --from=builder /app/tndcsc .

COPY . .

# 設定權限
RUN chmod +x ./tndcsc

# 暴露端口
EXPOSE 80

# 運行應用程式
CMD ["./tndcsc"]