package mqclient

import (
	"encoding/json"
	"github.com/yoda/common/pkg/eventbus"
	"github.com/yoda/webapp/internal/event"
	"time"
)

func SendRunTask(id string) {
	request := eventbus.MessageRunTaskRequest{
		ID:        id,
		Date:      time.Now(),
		JobId:     1,
		QueueName: event.CONSUMER_QUEUE_NAME,
	}
	msg, _ := json.Marshal(request)
	eventbus.Publish(event.PUBLISHER_QUEUE_NAME, eventbus.EVENT_RUN_JOB, id, msg)
}
