# README

## デプロイ

```bash
# ECRで表示されるプッシュコマンド（例）を
# prod環境用のターゲットでビルドするよう修正する
docker build -t hello-world-test .
# ↓
docker build -t hello-world-test --target prod .
```
