# wolf-bff
## 概要
「今日こと(What happened today)」システム(hodie)のBFF

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
            ↓ ↑ gRPC
           [auth]
```

### 機能
- 今日こと
- インセンティブ

### wolf-front
GraphQLフロントエンド（SSR）

### wolf-auth
認証認可（Auth0）

## ユースケース
### 今日こと(What happened today)
- ユーザは「今日こと」(wht)を登録する。
  - テキスト入力ができる。
  - 書いたものを写真に撮って画像を添付できる。
  - 録音した音声データを添付できる。
  - 撮影した動画データを添付できる。
- ユーザは登録した「今日こと」を確認できる。
  - 一覧で確認できる。
    - 最新順で確認できる。
    - 登録月や特定ワードを含むもの(テキスト入力のみ)といったフィルタリングができる。

### インセンティブ(Incentive)
- 保護者は「インセンティブ」(incentive)を登録する。
