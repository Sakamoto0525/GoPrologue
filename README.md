# Go Prologue

### 初回

#### .env作成

.envを作成し、Github名とアプリ名を設定する

```bash
cp .env.default .env
```

#### Dockerビルド

```bash:
docker-compose build
```

### 起動

```bash:
docker-compose up
# バックグラウンドでの起動
docker-compose up -d
```

[localhost:8080](localhost:8080)で動作を確認

