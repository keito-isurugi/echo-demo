package main

import (
	"fmt"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/labstack/echo"

	"echo-demo/internal/db"
)

type Todo struct {
	ID int
  Title         string
  Content       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	e := echo.New()

	db, _ := db.NewConnect()

	ag := e.Group("/architecture")
	ag.GET("/beta/todos", func(c echo.Context) error {
		var todos []Todo
		db.Find(&todos)
		fmt.Println(todos)
		return c.String(http.StatusOK, "todo一覧取得")
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