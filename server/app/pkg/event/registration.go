package event

import "github.com/yoda/common/pkg/eventbus"

func RegistrationQueue(e *eventbus.RegistrationRequest, msgId string) {
	AddQueue(e.QueueName)
}
