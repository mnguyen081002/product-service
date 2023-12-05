package mappermessage

import (
	"encoding/json"
	"productservice/internal/messaging/message"
	"github.com/segmentio/kafka-go"
)

func DecreaseProductQuantityMessage(kmsg kafka.Message) message.DecreaseProductQuantity {
	var msg message.DecreaseProductQuantity
	_ = json.Unmarshal(kmsg.Value, &msg)
	return msg
}
