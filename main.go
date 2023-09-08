package main

import (
	"fmt"
	"net/http"
	"time"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo"
	// "github.com/pkg/errors"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID int
  Title         string
  Content       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DB struct {
	PostgresHost     string `required:"true" split_words:"true"`
	PostgresPort     string `required:"true" split_words:"true"`
	PostgresDatabase string `required:"true" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}
func env() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("envを読み込めませんでした： %v", err)
	}

	message := os.Getenv("SAMPLE_MESSAGE")
	fmt.Println(message)
}
// func NewValue() (*Values, error) {
// 	var v Values
// 	err := envconfig.Process("", &v)

// 	// test環境の場合はtest用のDB情報を設定する
// 	if v.Env == "test" {
// 		v.DB.PostgresHost = v.TestDB.TestPostgresHost
// 		v.DB.PostgresPort = v.TestDB.TestPostgresPort
// 		v.DB.PostgresDatabase = v.TestDB.TestPostgresDatabase
// 		v.DB.PostgresUser = v.TestDB.TestPostgresUser
// 		v.DB.PostgresPassword = v.TestDB.TestPostgresPassword
// 	}

// 	if err != nil {
// 		s := fmt.Sprintf("need to set all env values %+v", v)
// 		return nil, errors.Wrap(err, s)
// 	}
// 	return &v, nil
// }


func main() {
	e := echo.New()

	ag := e.Group("/architecture")
	mg := ag.Group("/mono")
	mg.GET("/todo", func(c echo.Context) error {
		return c.String(http.StatusOK, "todo詳細取得")
	})


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.GET("/hoge", func(c echo.Context) error {
		var todos []Todo
		// TODO: envファイルから値を取ってくるようにする
		dsn := "host=db user=postgres password=postgres dbname=echo_demo port=5432 sslmode=disable TimeZone=Asia/Tokyo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		fmt.Println("PostgreSQLデータベースに接続しました")
		db.Find(&todos)
		fmt.Println(todos)
		return c.String(http.StatusOK, "hoge")
	})

	e.GET("/hoge/create", func(c echo.Context) error {
		todo := Todo{Title: "hogehogehoge", Content: "FOOFOOFO", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		// TODO: envファイルから値を取ってくるようにする
		dsn := "host=db user=postgres password=postgres dbname=echo_demo port=5432 sslmode=disable TimeZone=Asia/Tokyo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		fmt.Println("PostgreSQLデータベースに接続しました")
		fmt.Println(todo)
		fmt.Println(db)
		db.Create(&todo)
		return c.String(http.StatusOK, "hoge/create")
	})

	
	e.Logger.Fatal(e.Start(":8088"))
}