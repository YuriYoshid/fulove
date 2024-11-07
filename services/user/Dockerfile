FROM golang:1.21.0-alpine

WORKDIR /app

# 必要なパッケージのインストール
RUN apk add --no-cache gcc musl-dev postgresql-client

# Go.modをコピー
COPY go.mod .

# 依存関係のダウンロード
RUN go mod download && go mod verify

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# アプリケーションの実行
CMD ["./main"]
