package api

import (
	"expvar"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var successGRPcRequests utils.Counter
var grpcErrors = utils.NewErrorCounter()

func init() {
	expvar.Publish("Success GRPc requests count", &successGRPcRequests)
	expvar.Publish("GRPc error count", &grpcErrors)
}
