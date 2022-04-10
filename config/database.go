package config

import (
	"content-service/content"
	"content-service/content_type"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal Koneksi Ke Database MySQL")
	}

	fmt.Println("Berhasil Koneksi Ke Database")

	return db
}

func ConnectPostgres() (*gorm.DB, *sql.DB, error) {

	dsn := "host=localhost user=postgres password=root dbname=content-service port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dbConn.AutoMigrate(&content.Content{})
	dbConn.AutoMigrate(&content_type.Content_Type{})
	
	if err != nil {
		log.Println("Error connect to PostgreSQL: ", err.Error())
		return nil, nil, err
	}

	sqlDb, errDb := dbConn.DB()
	if errDb != nil {
		log.Println(errDb)
	} else {
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(1000)
	}

	log.Println("Postgres connection success")
	return dbConn, sqlDb, nil
}