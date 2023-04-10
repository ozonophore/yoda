package mqclient

import (
	"encoding/json"
	"github.com/yoda/common/pkg/mq"
)

func SendMessageEtlInfo(message mq.MessageETLInfoRequest) error {
	data, _ := json.Marshal(message)
	return mq.SendMessage(mq.HEADER_ETL_INFO, data)
}
