package cached

import (
	"expvar"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var hit utils.Counter
var miss utils.Counter

func init() {
	expvar.Publish("Cache hits", &hit)
	expvar.Publish("Cache miss", &miss)
}
