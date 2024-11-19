package db

import (
	"task_manager_app/internal/env"
	"task_manager_app/internal/gen"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Q *gen.Query

func InitializeGen(db *gorm.DB) {
	// genのクエリオブジェクトを初期化
	gen.SetDefault(db)
	Q = gen.Use(db)
}

func ConnectDB() error {

	// envの情報からDSNを設定
	dsn := env.DB_USER + ":" + env.DB_PASS + "@tcp(" + env.DB_HOST + ":" + env.DB_PORT + ")/" + env.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local" 

	// データベース接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
			return err
	}

	// グローバル変数にデータベース接続を設定
	DB = db

	InitializeGen(DB)

	return nil
}