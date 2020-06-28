# wolf-bff
## 概要
「今日こと」システムのBFF

## 使用ライブラリバージョン
### go
```
$ go version
go version go1.14.2 linux/amd64
```

### gqlgen
```
$ gqlgen version
v0.11.3-dev
```

## 全体
### 図
```
      GraphQL   gRPC
[front] -> [bff] -> [xxxx]
            ↓ ↑ REST
           [auth]
```

### 機能
- 今日こと
- インセンティブ

### wolf-front
GraphQLフロントエンド（SSR）

### wolf-auth
認証認可（Auth0）
