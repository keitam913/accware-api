package account

import (
	"time"

	"github.com/keitam913/accware-api/domain/account"
)

type Service struct {
	Repository account.Respository
}

func (s *Service) GetMonth(year int, month time.Month) (account.Month, error) {
	return s.Repository.Month(year, month)
}

func (s *Service) AddItem(name string, amount int, personID string, date time.Time) error {
	item, err := account.NewItem(name, amount, personID, date)
	if err != nil {
		return err
	}
	return s.Repository.Add(item)
}
