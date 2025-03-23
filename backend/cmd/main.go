package main

import (
	memosvr "backend/gen/http/memo/server"
	todosvr "backend/gen/http/todo/server"
	usersvr "backend/gen/http/user/server"
	"backend/gen/memo"
	"backend/gen/todo"
	"backend/gen/user"
	"backend/internal/db"
	"backend/internal/repository"
	"backend/pkg/service"
	"log"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
)

func main() {
	// ロガーの設定
	logger := log.New(os.Stdout, "[todo-app] ", log.Ltime)

	// データベース接続の初期化
	database, err := db.InitDB()
	if err != nil {
		logger.Fatalf("データベース接続エラー: %v", err)
	}

	// マイグレーションの実行
	if err := db.Migrate(database); err != nil {
		logger.Fatalf("マイグレーションエラー: %v", err)
	}

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(database)
	todoRepo := repository.NewTodoRepository(database)
	memoRepo := repository.NewMemoRepository(database)

	// サービスの初期化
	userSvc := service.NewUserService(userRepo)
	todoSvc := service.NewTodoService(todoRepo)
	memoSvc := service.NewMemoService(memoRepo)

	// エンドポイントの初期化
	userEndpoints := user.NewEndpoints(userSvc)
	todoEndpoints := todo.NewEndpoints(todoSvc)
	memoEndpoints := memo.NewEndpoints(memoSvc)

	// MUXの設定
	mux := goahttp.NewMuxer()

	// エンコーダ/デコーダの設定
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	// HTTPハンドラーの作成
	userHandler := usersvr.New(userEndpoints, mux, dec, enc, nil, nil)
	todoHandler := todosvr.New(todoEndpoints, mux, dec, enc, nil, nil)
	memoHandler := memosvr.New(memoEndpoints, mux, dec, enc, nil, nil)

	// ルートのマウント
	usersvr.Mount(mux, userHandler)
	todosvr.Mount(mux, todoHandler)
	memosvr.Mount(mux, memoHandler)

	// ルートのログ出力
	for _, m := range userHandler.Mounts {
		logger.Printf("User API: %s %s mounted on %s", m.Method, m.Verb, m.Pattern)
	}

	for _, m := range todoHandler.Mounts {
		logger.Printf("Todo API: %s %s mounted on %s", m.Method, m.Verb, m.Pattern)
	}

	for _, m := range memoHandler.Mounts {
		logger.Printf("Memo API: %s %s mounted on %s", m.Method, m.Verb, m.Pattern)
	}

	// サーバーの起動
	port := "8000"
	logger.Printf("HTTP server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		logger.Fatalf("サーバー起動エラー: %v", err)
	}
}
