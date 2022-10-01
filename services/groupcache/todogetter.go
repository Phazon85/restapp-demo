package groupcache

import (
	"encoding/json"
	"fmt"
	"time"

	todoService "github.com/Phazon85/restapp-demo/services/todos"
	"github.com/mailgun/groupcache"
)

func (s *Service) NewTodoGetter(todoServ *todoService.Service) groupcache.GetterFunc {
	return groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			s.Logger.Info(fmt.Sprintf("Cache Miss, Hitting DB for key: %s", key))
			oneMinuteFromNow := time.Now().Add(time.Minute)
			entry, err := todoServ.GetByID(key)
			if err != nil {
				return err
			}
			bytes, err := json.Marshal(entry)
			dest.SetBytes(bytes, oneMinuteFromNow)
			return nil
		},
	)
}
