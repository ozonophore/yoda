package controller

import (
	"encoding/json"
	"github.com/yoda/webapp/internal/service"
	"net/http"
)

func (sa ServerApi) GetOrganisations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	orgs, err := service.GetOrganisations()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orgs)
}
