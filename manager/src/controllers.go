package src

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	context2 "github.com/wepala/weos/context"
	"github.com/wepala/weos/controllers/rest"
	"github.com/wepala/weos/model"
	"github.com/wepala/weos/projections"
	"golang.org/x/net/context"
)

//Create Job
func CreateJob(api rest.Container, projection projections.Projection, commandDispatcher model.CommandDispatcher, eventSource model.EventRepository, entityFactory model.EntityFactory) echo.HandlerFunc {
	return func(c echo.Context) error {
		var command *model.Command
		ctxt := c.Request().Context()
		body := ctxt.Value(context2.PAYLOAD).([]byte)
		payload := make(map[string]interface{})
		err := json.Unmarshal(body, &payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)

		}
		ctxt = context.WithValue(ctxt, context2.ENTITY_FACTORY, entityFactory)
		var jobTitle string
		var jobDescription string
		resp := make(map[string]interface{})
		jobTitle = payload["title"].(string)
		jobDescription = payload["description"].(string)

		if jobTitle == "" || jobDescription == "" {
			resp["error"] = "title and description are required"
			return c.JSON(http.StatusBadRequest, resp)
		}
		filter := make(map[string]interface{})
		filterProperty := make(map[string]interface{})
		filterProperty["field"] = "title"
		filterProperty["operator"] = "eq"
		filterProperty["value"] = jobTitle
		filter["job"] = filterProperty

		_, len, err := projection.GetList(ctxt, entityFactory, 0, 0, "", nil, filter)
		if err != nil {
			resp["error"] = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}
		if len > 0 {
			resp["error"] = "job already exists"
			return c.JSON(http.StatusBadRequest, resp)
		}
		command = &model.Command{
			Type:    "create",
			Payload: []byte(`{"title":"` + jobTitle + `","description":"` + jobDescription + `"}`),
			Metadata: model.CommandMetadata{
				EntityType: "Job",
			},
		}
		err = commandDispatcher.Dispatch(ctxt, command, eventSource, projection, c.Logger())
		if err != nil {
			c.Logger().Debugf("there was an error creating job '%s'", err)
			resp["error"] = err.Error()
			return c.JSON(http.StatusBadRequest, resp)
		}
		resp["success"] = true
		return c.JSON(http.StatusOK, resp)
	}
}
