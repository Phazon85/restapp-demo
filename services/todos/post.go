package todos

import "fmt"

func (service *Service) Post(entry *Entry) error {
	createEntry := fmt.Sprintf("INSERT INTO todos (name, description) VALUES ('%s', '%s') RETURNING id;", entry.Name, entry.Description)
	_, err := service.DB.Exec(createEntry)
	if err != nil {
		return err
	}

	return nil
}
