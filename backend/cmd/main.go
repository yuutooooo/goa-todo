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

// CORSミドルウェアを定義
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS対応のヘッダーを設定
		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000, http://localhost:8081")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONSリクエストの場合は、ここで処理を終了
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 次のハンドラを呼び出す
		next.ServeHTTP(w, r)
	})
}

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

	// OpenAPI仕様ファイルの提供
	mux.Handle("GET", "/openapi.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w, r, "gen/http/openapi.json")
	}))

	// Swagger UI の提供
	mux.Handle("GET", "/swagger/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.StripPrefix("/swagger/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
	}))

	// Swagger UIのルートパスへのリダイレクト
	mux.Handle("GET", "/swagger", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	}))

	// CORSミドルウェアを適用
	handler := corsMiddleware(mux)

	// サーバーの起動
	port := "8000"
	logger.Printf("HTTP server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		logger.Fatalf("サーバー起動エラー: %v", err)
	}
}
