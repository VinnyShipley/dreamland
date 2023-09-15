package api

import (
	"context"
	"time"

	goHttp "net/http"

	"github.com/pterm/pterm"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

type multiverseService struct {
	rest httpIface.Service
	common.Multiverse
}

func BigBang() error {
	// err := error would be the same and might keep style more symmetrical?
	var err error

	srv := &multiverseService{
		Multiverse: services.NewMultiVerse(),
	}

	srv.rest, err = http.New(srv.Context(), options.Listen(common.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))

	// Logic check for errors passed into the function 
	if err != nil {
		return err
	}

	// API actually being created here
	srv.setUpHttpRoutes().Start()

	// Is this to give the server time to have the API be created before timing out or throwing error?
	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), 10*time.Second)
	defer waitCtxC()

	for {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		case <-time.After(100 * time.Millisecond):
			// If there is an error that crops up in the .1 second the switch statement, error message is presented
			if srv.rest.Error() != nil {
				pterm.Error.Println("Dreamland failed to start")
				return srv.rest.Error()
			}
			_, err := goHttp.Get("http://" + common.DreamlandApiListen)
			// If there is no errors Dreamland instance, or Universe, is created
			if err == nil {
				pterm.Info.Println("Dreamland ready")
				return nil
			}
		}
	}
}
