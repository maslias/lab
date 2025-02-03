package models

import (
	"database/sql"
	"time"
)

type Task struct {
	Id        int
	Title     string
	Details   sql.NullString
	CreatedAt time.Time
    TerminatedAt time.Time
	DoneAt    sql.NullTime
}
