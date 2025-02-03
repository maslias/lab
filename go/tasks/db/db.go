package db

import (
	"database/sql"
	"log"

	"github.com/maslias/tasks/config"
	"github.com/maslias/tasks/db/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type AppDb struct {
	Tasks *sqlite.TaskModel
}

func (a *AppDb) Execute() {

    config.Load()

	db, err := sql.Open(config.GET().DB_DRIVER, config.GET().DB_PATH)
	// db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}

    a.Tasks = &sqlite.TaskModel{
        DB: db,
    }
}

func NewAppDb() *AppDb {


    config.Load()

	db, err := sql.Open(config.GET().DB_DRIVER, config.GET().DB_PATH)
	// db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	return &AppDb{
		Tasks: &sqlite.TaskModel{
            DB: db,
        },
	}
}
