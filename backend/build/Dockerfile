FROM golang:1.17-alpine

RUN apk update && apk add git

# 作業ディレクトリの指定
WORKDIR  /go/src/github.com/Seiji-Ikeda32/go-todo-app/backend

# コンテナの作業ディレクトリにファイルを移行
ADD . /go/src/github.com/Seiji-Ikeda32/go-todo-app/backend

# モジュールをインストール
RUN go mod download

# ホットリロードライブラリをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air"]

