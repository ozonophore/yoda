package eventbus

const (
	EVENT_REGISTRATION = "registration"
	EVENT_JOB_ADD      = "job_add"
	EVENT_JOB_REFRESH  = "job_refresh"
)

type RegistrationRequest struct {
	QueueName string `json:"queue_name"`
}

type JobResponse = MQJob
