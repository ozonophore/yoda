package event

import (
	"encoding/json"
	"github.com/yoda/common/pkg/eventbus"
)

func RegistrationQueue(msgId, name string) {
	event := &eventbus.RegistrationRequest{
		QueueName: name,
	}
	body, _ := json.Marshal(event)
	eventbus.Publish(PUBLISHER_QUEUE_NAME, eventbus.EVENT_REGISTRATION, msgId, body)
}
