package todos

import "fmt"

func (service *Service) Delete(id string) error {
	deleteEntry := fmt.Sprintf("DELETE FROM todos WHERE id = %s ", id)
	_, err := service.DB.Exec(deleteEntry)
	if err != nil {
		return err
	}

	return nil
}
