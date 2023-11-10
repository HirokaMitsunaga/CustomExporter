FROM golang:1.20-alpine AS build

# Goコンテナのマルチステージビルドを実行
WORKDIR /src/
COPY main.go go.* /src/
# templateバイナリを作成
RUN CGO_ENABLED=0 go build -o /bin/template

# スクラッチイメージ(空のコンテナイメージ)を用意
FROM scratch
# スクラッチイメージにtemplateバイナリをコピー
COPY --from=build /bin/template /bin/template
# templateバイナリの実行(exec)
ENTRYPOINT ["/bin/template"]
