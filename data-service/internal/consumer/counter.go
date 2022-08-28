package consumer

import (
	"expvar"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var successfullyProcessedMessages utils.Counter
var messagesErrors = utils.NewErrorCounter()
var responsesGiven utils.Counter

func init() {
	expvar.Publish("Successfully processed messages", &successfullyProcessedMessages)
	expvar.Publish("Messages error count", &messagesErrors)
	expvar.Publish("Given response count", &responsesGiven)
}
