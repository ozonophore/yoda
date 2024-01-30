package storage

import "context"

type Marketplace struct {
	Code      string `gorm:"column:code"`
	Name      string `gorm:"column:name"`
	ShortName string `gorm:"column:short_name"`
}

func (s *Storage) GetMarketplaces(ctx context.Context) (*[]Marketplace, error) {
	var marketplaces []Marketplace
	tx := s.db.WithContext(ctx).Table("ml.marketplace").Scan(&marketplaces)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &marketplaces, nil

}
