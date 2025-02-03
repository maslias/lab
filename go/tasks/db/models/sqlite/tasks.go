package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/maslias/tasks/db/models"
)

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) All(isU bool, isF bool) ([]models.Task, error) {
	whereStmt := ``
	if isU && !isF {
		whereStmt = `where doneAt is null`
	} else if isF && !isU {
		whereStmt = `where doneAt is not null`
	}

	stmt := `select id, title, details, createdAt, terminatedAt, doneAt  from tasks ` +whereStmt+` order by (julianday(terminatedAt) - julianday(createdAt))`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	tasks := []models.Task{}
	for rows.Next() {
		t := models.Task{}
		err := rows.Scan(&t.Id, &t.Title, &t.Details, &t.CreatedAt,&t.TerminatedAt, &t.DoneAt)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (m *TaskModel) Add(title string, details string, terminatedAt int) ([]models.Task, error) {

	stmt := `insert into tasks (title, details, terminatedAt) values(?, ?, datetime(current_timestamp, ?)) returning id, title, details, createdAt, terminatedAt, doneAt`

    t := models.Task{}
	if err := m.DB.QueryRow(stmt, title, details, fmt.Sprintf(`+%d days`,terminatedAt)).Scan(&t.Id, &t.Title, &t.Details, &t.CreatedAt, &t.TerminatedAt, &t.DoneAt); err != nil {
        return nil, err
    }

    ts := []models.Task{}
    ts = append(ts, t)

    return ts, nil
}

func (m *TaskModel) Complete(id int) ([]models.Task, error){
    stmt := `update tasks set doneAt = current_timestamp where id = ? returning id, title, details, createdAt, terminatedAt, doneAt`
    t := models.Task{}
    if err:= m.DB.QueryRow(stmt, id).Scan(&t.Id, &t.Title, &t.Details, &t.CreatedAt, &t.TerminatedAt, &t.DoneAt); err != nil {
        return nil, err
    }

    ts := []models.Task{}
    ts = append(ts, t)

    return ts, nil
}

func (m *TaskModel) Delete(id int) ([]models.Task, error){
    stmt := `delete from tasks where id = ? returning id, title, details, createdAt, terminatedAt, doneAt`
    t := models.Task{}
    if err:= m.DB.QueryRow(stmt, id).Scan(&t.Id, &t.Title, &t.Details, &t.CreatedAt, &t.TerminatedAt, &t.DoneAt); err != nil {
        return nil, err
    }

    ts := []models.Task{}
    ts = append(ts, t)

    return ts, nil
}
