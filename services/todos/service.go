package todos

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	DB     *sql.DB
}

func New(logger *zap.Logger, db *sql.DB) *Service {
	return &Service{
		Logger: logger,
		DB:     db,
	}
}

func (service *Service) GetTodo() ([]*Entry, error) {
	const allEntries = "SELECT * FROM todos"
	var todos []*Entry

	rows, err := service.DB.Query(allEntries)
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

func (service *Service) GetTodoByID(id string) (*Entry, error) {
	const entryByID = "SELECT id, name, description FROM todos WHERE id = %s"
	todo := &Entry{}
	row := service.DB.QueryRow(fmt.Sprintf(entryByID, id))
	if err := row.Scan(&todo.ID, &todo.Name, &todo.Description); err != nil {
		return nil, err
	}

	return todo, nil
}

func (service *Service) PutTodo(entry *Entry) error {
	putEntry := fmt.Sprintf("UPDATE todos SET (name, description) = ('%s', '%s') WHERE id = %s;", entry.Name, entry.Description, entry.ID)
	_, err := service.DB.Exec(putEntry)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) PostTodo(entry *Entry) error {
	createEntry := fmt.Sprintf("INSERT INTO todos (name, description) VALUES ('%s', '%s') RETURNING id;", entry.Name, entry.Description)
	_, err := service.DB.Exec(createEntry)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) DeleteTodo(id string) error {
	deleteEntry := fmt.Sprintf("DELETE FROM todos WHERE id = %s ", id)
	_, err := service.DB.Exec(deleteEntry)
	if err != nil {
		return err
	}

	return nil
}
