# README

## ビルド

```bash
make all
```

## ローカルでのDocker起動

```bash
# 開発環境（ホットリロード）
docker-compose up
# 本番環境（バイナリから起動）
docker run -p 8080:8080 --rm api-prod
```

## デプロイ

```bash
# ECRで表示されるプッシュコマンド（例）を
# prod環境用のターゲットでビルドするよう修正する
docker build -t echo-server-test .
# ↓
docker build -t echo-server-test --target prod .
```
