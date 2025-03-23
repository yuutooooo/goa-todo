# goa-todo

## アプリケーション概要

このアプリケーションは、ユーザー登録からタスク管理、メモ追加までの機能を持つTODOアプリです。シンプルなUIと直感的な操作で、日々のタスク管理を効率的に行うことができます。

## 使用技術スタック

### バックエンド
- **言語**: Go
- **フレームワーク**: Goa（APIデザインフレームワーク）
- **データベース**: MySQL（開発環境）
- **データベースクライアント**: PHPMyAdmin
- **API仕様**: OpenAPI（Swagger）

### フロントエンド
- **言語**: TypeScript
- **フレームワーク**: React
- **UIライブラリ**: Material UI
- **ルーティング**: React Router
- **HTTP通信**: Axios

### 開発・実行環境
- **コンテナ化**: Docker / Docker Compose
- **開発サーバー**: React開発サーバー (port: 3000)
- **APIサーバー**: Go (port: 8000)

## 主な機能

- **ユーザー管理**
  - ユーザー登録
  - ログイン/ログアウト
  - ユーザー情報の表示・編集

- **TODO管理**
  - TODO一覧表示
  - TODO作成・編集・削除
  - TODO完了状態の切り替え

- **メモ管理**
  - TODOに紐づくメモの作成
  - メモの一覧表示・編集・削除

## アーキテクチャ

- **フロントエンド**: SPA（Single Page Application）
- **バックエンド**: RESTful API（Goaで設計・実装）
- **データモデル**: ユーザー、TODO、メモの3層構造

## 開発者向け情報

### 起動方法
```bash
# バックエンド起動
cd backend && go run cmd/main.go

# フロントエンド起動
cd frontend && npm start
```

- アプリケーションURL
  http://localhost:3000
- phpmyadmin
  http://localhost:8080/
- api仕様書(swagger UI)
  http://localhost:8000/swagger/

