package event

import (
	"encoding/json"
	"github.com/yoda/common/pkg/eventbus"
)

func RefreshJobs(jobs []eventbus.MQJob) {
	body, _ := json.Marshal(&jobs)
	PublishToAll(&body, eventbus.EVENT_JOB_REFRESH, "0")
}
