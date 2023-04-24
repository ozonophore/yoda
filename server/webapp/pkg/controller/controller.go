package controller

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/yoda/webapp/pkg/api"
	"github.com/yoda/webapp/pkg/service"
	"net/http"
	"time"
)

type ServerApi struct {
	logger *logrus.Logger
}

func (sa ServerApi) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	info, err := service.GetTasksInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}

func (sa ServerApi) GetStocksInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	info, err := service.GetStocksInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}

func (sa ServerApi) GetTransactionsInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	info, err := service.GetTransactionsInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}

func (sa ServerApi) GetSalesForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	sales, err := service.GetSalesForWeek()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sales)
}

func (sa ServerApi) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	date := `{"date": "` + time.Now().Format(time.DateTime) + `"}`
	w.Write([]byte(date))
}

func (sa ServerApi) GetOrders(w http.ResponseWriter, r *http.Request, params api.GetOrdersParams) {
	w.Header().Add("Content-Type", "application/json")
	result, err := service.GetOrders(params.Limit, params.Offset, params.Date, params.Search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (sa ServerApi) GetStocks(w http.ResponseWriter, r *http.Request, params api.GetStocksParams) {
	w.Header().Add("Content-Type", "application/json")
	result, err := service.GetStocks(params.Limit, params.Offset, params.Date, params.Search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (sa ServerApi) GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result, err := service.GetRooms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)
}

func (sa ServerApi) CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, err := api.NewCreateRoom(r)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := service.CreateRoom(*requestBody)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (sa ServerApi) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (sa ServerApi) GetRoomById(w http.ResponseWriter, r *http.Request, code string) {
	//TODO implement me
	panic("implement me")
}

// Get all jobs
// (GET /jobs)
func (sa ServerApi) GetJobs(w http.ResponseWriter, r *http.Request) {
	sa.logger.Debug("GetJobs")
	w.Header().Set("Content-Type", "application/json")
	jobs, err := service.GetJobs()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*jobs)
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

func NewServerApi(logger *logrus.Logger) api.ServerInterface {
	return ServerApi{
		logger: logger,
	}
}
