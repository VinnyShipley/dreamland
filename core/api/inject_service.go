package api

import (
	"fmt"

	commonIface "github.com/taubyte/go-interfaces/common"
	httpIface "github.com/taubyte/http"
)

func injectServiceHttp() {
	// Path to create services in a universe
	api.POST(&httpIface.RouteDefinition{
		Path: "/service/{universe}/{name}",
		Vars: httpIface.Variables{
			Required: []string{"universe", "name", "config"},
		},
		Handler: apiHandlerService,
	})
}

func apiHandlerService(ctx httpIface.Context) (interface{}, error) {
	// Grab the universe
	universe, err := getUniverse(ctx)
	if err != nil {
		return nil, fmt.Errorf("killing service failed with: %s", err.Error())
	}

	name, err := ctx.GetStringVariable("name")
	if err != nil {
		return nil, fmt.Errorf("failed getting name error %w", err)
	}

	config := struct {
		Config *commonIface.ServiceConfig
	}{}

	err = ctx.ParseBody(&config)
	if err != nil {
		return nil, err
	}

	err = universe.Service(name, config.Config)
	if err != nil {
		return nil, fmt.Errorf("failed creating service `%s` failed with: %v", name, err)
	}

	return nil, nil
}
