package service

import (
	"github.com/yoda/webapp/internal/api"
	"github.com/yoda/webapp/internal/dao"
)

func GetOrganisations() (*[]api.Organisation, error) {
	orgs, err := dao.GetOrganisations()
	if err != nil {
		return nil, err
	}
	if orgs == nil || len(*orgs) == 0 {
		return &[]api.Organisation{}, nil
	}
	var result []api.Organisation
	for _, org := range *orgs {
		result = append(result, api.Organisation{
			Id:   org.ID,
			Name: org.Name,
		})
	}
	return &result, nil
}
