package boot

import (
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbStart() *gorm.DB {
	env := envService.GetEnvServiceInstance()
	dsn := env.Env.DBString
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
