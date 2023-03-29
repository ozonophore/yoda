package mqserver

import (
	"encoding/json"
	"github.com/yoda/common/pkg/mq"
)

func SendMessageEtlInfoResponse(message mq.MessageETLInfoResponse) error {
	data, _ := json.Marshal(message)
	return mq.SendMessage(mq.HEADER_ETL_INFO, data)
}
