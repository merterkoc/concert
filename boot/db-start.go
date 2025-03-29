package boot

import (
	"database/sql"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbStart() *gorm.DB {
	env := envService.GetEnvServiceInstance()
	var dsnWithoutDB = "root:1234@tcp(localhost:3306)/?parseTime=True"

	dsn := env.Env.DBString

	dbTemp, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		log.Fatal("Failed to connect to MySQL server:", err)
	}
	defer func(dbTemp *sql.DB) {
		err := dbTemp.Close()
		if err != nil {

		}
	}(dbTemp)

	_, err = dbTemp.Exec("CREATE DATABASE IF NOT EXISTS gigbuddy;")
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	_, err = dbTemp.Exec("USE gigbuddy;")
	if err != nil {
		log.Fatal("Failed to select database:", err)
	}

	_, err = dbTemp.Exec("CREATE TABLE IF NOT EXISTS goose_db_version (\n    id BIGINT AUTO_INCREMENT PRIMARY KEY,\n    version_id BIGINT NOT NULL,\n    is_applied BOOLEAN NOT NULL,\n    tstamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP\n);")
	if err != nil {
		log.Fatal("Failed to create goose_db_version table:", err)
	}

	_, err = dbTemp.Exec("INSERT INTO goose_db_version (version_id, is_applied)\nSELECT 0, 1\nFROM DUAL\nWHERE NOT EXISTS (\n    SELECT 1 FROM goose_db_version WHERE version_id = 0 AND is_applied = 1\n);;")
	if err != nil {
		log.Fatal("Failed to insert goose_db_version table:", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB:", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to set dialect:", err)
		return nil
	}
	err = goose.SetDialect("mysql")
	if err != nil {
		return nil
	}

	if err := goose.Up(sqlDB, "./migrations"); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Database connected and migrations applied successfully!")
	return db
}
