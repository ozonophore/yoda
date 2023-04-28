package service

import (
	"errors"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/webapp/internal/api"
	"github.com/yoda/webapp/internal/dao"
	"github.com/yoda/webapp/internal/mapper"
	"strings"
)

func CreateRoom(newRoom api.NewRoom) (*api.Room, error) {
	owner := mapper.MapRoomToOwner(newRoom)
	ozon := mapper.MapRoomToOzon(newRoom)
	wb := mapper.MapRoomToWb(newRoom)
	err := dao.SaveOwner(owner, ozon, wb)
	if err != nil {
		return nil, err
	}

	room := api.Room{
		Code: newRoom.Code,
		Name: newRoom.Name,
		Ozon: newRoom.Ozon,
		Wb:   newRoom.Wb,
	}

	job, err := dao.GetJobById(1)
	if err != nil {
		return nil, err
	}
	days := strings.Split(*job.WeekDays, ",")
	time := strings.Split(*job.AtTime, ",")

	newDays := make([]api.WeekDay, len(days))
	for i, day := range days {
		newDays[i] = api.WeekDay(day)
	}
	room.Days = newDays
	room.Times = time

	return &room, nil
}

func GetRooms() (*[]api.Room, error) {
	owners, err := dao.GetOwners()
	if err != nil {
		return nil, err
	}
	if len(owners) == 0 {
		return &[]api.Room{}, nil
	}
	rooms := mapper.MapOwnersToRooms(owners)
	mps, err := dao.GetMarketplacesByOwners(owners)
	if err != nil {
		return nil, err
	}
	job, err := dao.GetJobById(1)
	if err != nil {
		return nil, err
	}
	days := strings.Split(*job.WeekDays, ",")
	newTimes := strings.Split(*job.AtTime, ",")
	newDays := make([]api.WeekDay, len(days))
	for i := 0; len(days) > i; i++ {
		newDays[i] = api.WeekDay(days[i])
	}
	roomCodes := make(map[string]*api.Room)
	for i := 0; len(*rooms) > i; i++ {
		room := &(*rooms)[i]
		room.Days = newDays
		room.Times = newTimes
		roomCodes[room.Code] = room
	}
	for _, mp := range *mps {
		room := roomCodes[mp.OwnerCode]
		switch mp.Source {
		case types.SourceTypeOzon:
			room.Ozon = api.Ozon{
				ApiKey:   *mp.Password,
				ClientId: *mp.ClientID,
			}
		case types.SourceTypeWB:
			room.Wb = api.Wb{
				AuthToken: *mp.Password,
			}
		}
	}
	return rooms, nil
}

func GetJobs() (*api.Job, error) {
	jobs, err := dao.GetJobs()
	if err != nil {
		return nil, errors.Join(err)
	}
	return mapper.MapJobsToApi(jobs), nil
}
