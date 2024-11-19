package main

import (
	"fmt"
	"task_manager_app/internal/db"
	"task_manager_app/internal/env"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	// generatorの設定
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/gen",
		Mode: gen.WithDefaultQuery | gen.WithoutContext,
		FieldCoverable: true,
	})
	
	// 環境変数を読み込んでDB接続を確立
	if err := env.Getenv("./internal/env/.env"); err != nil {
		panic(fmt.Errorf("cannot get environmental variables: %v", err))
	}
	if err := db.ConnectDB(); err != nil {
		panic(fmt.Errorf("cannot connect db: %v", err))
	}
	g.UseDB(db.DB)

	// userモデルとprojectモデルを作成
	user := g.GenerateModel("users")
	project := g.GenerateModel("projects")

	g.ApplyBasic(g.GenerateModel("users", gen.FieldRelate(field.HasMany, "Projects", project, &field.RelateConfig{
		RelateSlicePointer: true,
	})), g.GenerateModel("projects", gen.FieldRelate(field.BelongsTo, "User", user, &field.RelateConfig{
		RelatePointer: true,
	})))

	g.Execute()
}