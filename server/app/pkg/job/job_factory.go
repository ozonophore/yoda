package job

import (
	"context"
	"errors"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/service"
	"github.com/yoda/common/pkg/types"
)

type DataLoader interface {
	Parsing(context context.Context, transactionID int64) error
}

func JobFactory(source, owner, password string, clientId *string, config *configuration.Config) (DataLoader, error) {
	switch source {
	case types.SourceTypeWB:
		return service.NewWBService(owner, password, config), nil
	case types.SourceTypeOzon:
		if clientId == nil {
			return nil, errors.New("client id is empty")
		}
		return service.NewOzonService(owner, *clientId, password, config), nil
	default:
		return nil, nil
	}
}
