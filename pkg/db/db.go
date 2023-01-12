package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"news/pkg/models"
	"os"
	"sync"
)

var (
	db         *sql.DB
	dbConnOnce sync.Once
)

func GetDB() (*sql.DB, error) {

	var err error

	dbConnOnce.Do(func() {

		USER := os.Getenv("DB_USER")
		PASS := os.Getenv("DB_PASSWORD")
		HOST := os.Getenv("DB_HOST")
		DBNAME := os.Getenv("DB_NAME")

		URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
			HOST, DBNAME)

		d, _err := sql.Open("mysql", URL)
		if _err != nil {
			fmt.Printf("sql.Open failed: %v\n", _err)
			err = _err
			return
		}
		db = d

	})

	return db, err
}

func InsertData(ctx context.Context, data []*models.MainData) error {
	conn, err := GetDB()
	if err != nil {
		return fmt.Errorf("initDB failed: %w", err)
	}

	for _, value := range data {
		var _err error

		stmt := fmt.Sprintf("INSERT INTO main_data(uuid, headline, description, keywords, snippet, url) VALUES(?, ?, ?, ?, ?, ?)")

		_, err := conn.QueryContext(ctx, stmt, value.Uuid, value.Headline, value.Description, value.Keywords, value.Snippet, value.Url)
		if err != nil {
			_err = fmt.Errorf("insert failed: %w", err)
			return _err
		}
	}

	return nil
}
