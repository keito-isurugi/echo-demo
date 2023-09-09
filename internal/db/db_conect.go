package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	PostgresHost     string `required:"true" split_words:"true"`
	PostgresPort     string `required:"true" split_words:"true"`
	PostgresDatabase string `required:"true" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}

func dbenv() DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("envを読み込めませんでした： %v", err)
	}

	dbenv := DB{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
	}

	return dbenv
}

func NewConect() (*gorm.DB, error) {
	env := dbenv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		env.PostgresHost,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDatabase,
		env.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("PostgreSQLデータベースに接続しました")
	return db, nil
}
