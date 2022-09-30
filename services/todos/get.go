package todos

import (
	"fmt"
)

func (service *Service) Get() ([]*Entry, error) {
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

func (service *Service) GetByID(id string) (*Entry, error) {
	const entryByID = "SELECT id, name, description FROM todos WHERE id = %s"
	todo := &Entry{}
	row := service.DB.QueryRow(fmt.Sprintf(entryByID, id))
	if err := row.Scan(&todo.ID, &todo.Name, &todo.Description); err != nil {
		return nil, err
	}

	return todo, nil
}
