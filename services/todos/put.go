package todos

import "fmt"

func (service *Service) Put(entry *Entry) error {
	putEntry := fmt.Sprintf("UPDATE todos SET (name, description) = ('%s', '%s') WHERE id = %s;", entry.Name, entry.Description, entry.ID)
	_, err := service.DB.Exec(putEntry)
	if err != nil {
		return err
	}

	return nil
}
