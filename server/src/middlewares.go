package src

import (
	"encoding/json"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/wepala/weos/controllers/rest"
	"github.com/wepala/weos/model"
	"github.com/wepala/weos/projections"
	"golang.org/x/net/context"
)

//ViewDataMiddleware fetch data to inject into the view
func ViewDataMiddleware(api rest.Container, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory, path *openapi3.PathItem, operation *openapi3.Operation) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		log, _ := api.GetLog("Default")
		httpClient, err := api.GetHTTPClient("Default")
		if err != nil {
			log.Errorf("error setting up http client '%s'", err)
		}
		return func(ctxt echo.Context) error {
			newContext := ctxt.Request().Context()
			//loop through the x-data extension and fetch data
			if raw, ok := operation.Extensions["x-data"].(json.RawMessage); ok {
				var data map[string]interface{}
				json.Unmarshal(raw, &data)
				//get the data using the value as an url
				for k, v := range data {
					if url, ok := v.(string); ok {
						resp, err := httpClient.Get(url)
						if err != nil {
							log.Debugf("error getting data from '%s' '%s'", url, err)
						}
						var payload map[string]interface{}
						if resp != nil {
							data, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								log.Debugf("error reading response from url '%s'", err)
							}
							json.Unmarshal(data, &payload)
							newContext = context.WithValue(newContext, k, payload)
						}
					}
				}
			}

			//add id to context for use by controller

			request := ctxt.Request().WithContext(newContext)
			ctxt.SetRequest(request)
			return next(ctxt)
		}
	}
}

//ScriptsMiddleware get the list of javascript files
func ScriptsMiddleware(api rest.Container, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory, path *openapi3.PathItem, operation *openapi3.Operation) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO read the folder to see what files are in there
			//TODO create list of files using the filenames without the extension
			//TODO set list to the context
			return next(c)
		}
	}
}
