package eventbus

import "time"

const (
	EVENT_REGISTRATION = "registration"
	EVENT_JOB_ADD      = "job_add"
	EVENT_JOB_REFRESH  = "job_refresh"
	EVENT_RUN_JOB      = "run_job"
	EVENT_UPDATE_ORG   = "update_org"
)

type RegistrationRequest struct {
	QueueName string `json:"queue_name"`
}

type JobResponse = MQJob

type MessageRunTaskRequest struct {
	ID        string    `json:"id"`         // ID of the message
	Date      time.Time `json:"date"`       // Date of the message
	JobId     int       `json:"job_id"`     // ID of the job
	QueueName string    `json:"queue_name"` // Name of the queue
}

type MessageRunTaskResponse struct {
	ID    string    `json:"id"`     // ID of the message
	JobId int       `json:"job_id"` // ID of the job
	Date  time.Time `json:"date"`   // Date of the message
}

type EmptyResponse struct {
	ID string `json:"id"` // ID of the message
}
