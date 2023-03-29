package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ServerApi struct {
	logger *logrus.Logger
}

// Get all jobs
// (GET /jobs)
func (sa ServerApi) GetJobs(w http.ResponseWriter, r *http.Request) {
	sa.logger.Debug("GetJobs")
	id := int64(1)
	name := "test"
	newJob := Job{
		Id:   &id,
		Name: &name,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newJob)
}

// Create job
// (POST /jobs)
func (sa ServerApi) CreateJob(w http.ResponseWriter, r *http.Request) {
	sa.logger.Debug("CreateJob")
}

// Get job by id
// (GET /jobs/{id})
func (sa ServerApi) GetJobById(w http.ResponseWriter, r *http.Request, id int64) {
	sa.logger.Debugf("GetJobById: %d", id)
}

// Run job by id
// (POST /jobs/{id}/run)
func (sa ServerApi) RunJobById(w http.ResponseWriter, r *http.Request, id int64) {
	sa.logger.Debugf("RunJobById: %d", id)
}

// Stop job by id
// (POST /jobs/{id}/stop)
func (sa ServerApi) StopJobById(w http.ResponseWriter, r *http.Request, id int64) {
	sa.logger.Debugf("StopJobById: %d", id)
}

func NewServerApi(logger *logrus.Logger) ServerInterface {
	return ServerApi{
		logger: logger,
	}
}
