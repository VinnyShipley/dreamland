package api

import (
	"time"

	"github.com/taubyte/dreamland/core/common"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
)

var api httpIface.Service
var mv common.Multiverse

func Start(m common.Multiverse) (err error) {
	mv = m
	api, err = http.New(m.Context(), options.Listen(common.DreamlandApiListen))
	if err != nil {
		return
	}

	setUpHttpRoutes()

	api.Start()

	time.Sleep(300 * time.Millisecond)

	err = api.Error()

	return
}
