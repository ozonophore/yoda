package api

import (
	"encoding/json"
	"net/http"
)

func NewCreateRoom(r *http.Request) (*CreateRoomJSONRequestBody, error) {
	rq := CreateRoomJSONRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		return nil, err
	}
	return &rq, nil
}
