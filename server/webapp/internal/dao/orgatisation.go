package dao

import "github.com/yoda/common/pkg/model"

func GetOrganisations() (*[]model.Organisation, error) {
	var orgs []model.Organisation
	if err := dao.database.Find(&orgs).Error; err != nil {
		return nil, err
	}
	return &orgs, nil
}
