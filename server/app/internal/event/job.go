package event

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/yoda/common/pkg/eventbus"
	"time"
)

func RefreshJobs(jobs []eventbus.MQJob) {
	body, _ := json.Marshal(&jobs)
	PublishToAll(&body, eventbus.EVENT_JOB_REFRESH, "0")
}

func RunJob(id int) {
	response := eventbus.MessageRunTaskResponse{
		JobId: id,
		Date:  time.Now(),
	}
	body, _ := json.Marshal(&response)
	logrus.Debug("Send message to run job: ", id)
	PublishToAll(&body, eventbus.EVENT_RUN_JOB, uuid.New().String())
}
