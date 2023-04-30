package mapper

import (
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/webapp/internal/api"
	"time"
)

func MapRoomToOwner(s api.NewRoom) model.Owner {
	return model.Owner{
		Code:           s.Code,
		Name:           s.Name,
		OrganisationId: s.OrganisationId,
		CreateDate:     time.Now(),
	}
}

func MapRoomToOzon(s api.NewRoom) model.OwnerMarketplace {
	return model.OwnerMarketplace{
		OwnerCode: s.Code,
		Source:    types.SourceTypeOzon,
		Password:  &s.Ozon.ApiKey,
		ClientID:  &s.Ozon.ClientId,
	}
}

func MapRoomToWb(s api.NewRoom) model.OwnerMarketplace {
	return model.OwnerMarketplace{
		OwnerCode: s.Code,
		Source:    types.SourceTypeWB,
		Password:  &s.Wb.AuthToken,
	}
}

func MapOwnersToRooms(owners []model.Owner) *[]api.Room {
	rooms := make([]api.Room, len(owners))
	for i, owner := range owners {
		rooms[i] = MapOwnerToRoom(owner)
	}
	return &rooms
}

func MapOwnerToRoom(owner model.Owner) api.Room {
	return api.Room{
		Code:           owner.Code,
		Name:           owner.Name,
		OrganisationId: owner.OrganisationId,
		CreatedAt:      &owner.CreateDate,
		Organisation:   owner.OrganisationName,
	}
}
