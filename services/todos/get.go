package todos

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
