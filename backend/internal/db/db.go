package db

import (
	"backend/internal/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 環境変数から設定を取得、または既定値を使用
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// データベース接続を初期化
func InitDB() (*gorm.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "mysql")
	password := getEnv("DB_PASSWORD", "mysql")
	dbname := getEnv("DB_NAME", "todo_app")

	// MySQL接続文字列: user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// GORMロガーの設定
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Info, // ログレベル
			IgnoreRecordNotFoundError: true,        // レコード未検出エラーを無視
			Colorful:                  true,        // カラー出力
		},
	)

	// MySQL接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// グローバル変数に設定
	DB = db

	return db, nil
}

// マイグレーションを実行
func Migrate(db *gorm.DB) error {
	// 全モデルのマイグレーション
	return db.AutoMigrate(
		&model.User{},
		&model.Todo{},
		&model.Memo{},
	)
}
