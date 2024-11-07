FROM golang:1.21.0-alpine

WORKDIR /app

# 必要なパッケージのインストール
RUN apk add --no-cache gcc musl-dev postgresql-client

# アップロードディレクトリの作成
RUN mkdir -p /app/uploads/images

# Go.modとGo.sumをコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# UUIDパッケージの追加
RUN go get github.com/google/uuid

# ソースコードのコピー
COPY . .

# ビルド
RUN go build -o main .

# ポート公開
EXPOSE 8080

# 実行
CMD ["./main"]