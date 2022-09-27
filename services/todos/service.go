package todos

import (
	"database/sql"
)

type Service struct {
	DB *sql.DB
}

func New(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (repo *Service) GetTodos() ([]*Entry, error) {
	var todos []*Entry

	rows, err := repo.DB.Query("SELECT * FROM todos")
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		newEntry := &Entry{}
		err = rows.Scan(&newEntry.ID, &newEntry.Name, &newEntry.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, newEntry)
	}

	return todos, nil
}
