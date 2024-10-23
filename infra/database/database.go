package database

import (
	"fmt"
	"goTemplate/modules/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type database struct {
	database *gorm.DB
}

func (d *database) createDatabase() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	d.database = db
}

func (d *database) migrateDatabase() {
	err := d.database.AutoMigrate(&user.User{})

	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
}

func (d *database) InitDatabase() {
	d.createDatabase()
	d.migrateDatabase()
}

func (d *database) GetDatabase() *gorm.DB {

	if d.database == nil {
		d.InitDatabase()
	}

	return d.database
}
