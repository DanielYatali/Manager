package src

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	context2 "github.com/wepala/weos/context"
	"github.com/wepala/weos/controllers/rest"
	"github.com/wepala/weos/model"
	"golang.org/x/net/context"
)

func DefaultJobsInitializer(ctxt context.Context, tapi rest.Container, swagger *openapi3.Swagger) (context.Context, error) {
	//get default logger
	log, err := tapi.GetLog("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default jobs '%s'", err)
	}
	log.Infof("started setting up default jobs")
	//event store
	eventStore, err := tapi.GetEventStore("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default jobs '%s'", err)
	}
	//get default projection
	projection, err := tapi.GetProjection("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default jobs '%s'", err)
	}
	dispatcher, err := tapi.GetCommandDispatcher("Default")

	jobs := []map[string]interface{}{
		{
			"title":       "Intern",
			"description": "Assitant to developers",
		},
		{
			"title":       "Developer",
			"description": "Builds websites",
		},
		{
			"title":       "Scrum Master",
			"description": "Manages the team",
		},
		{
			"title":       "CEO",
			"description": "Manages the company",
		},
	}

	for _, job := range jobs {

		data, err := json.Marshal(job)
		if err != nil {
			return ctxt, fmt.Errorf("error setting up default jobs '%s'", err)
		}
		command := &model.Command{
			Type:    "create",
			Payload: data,
			Metadata: model.CommandMetadata{
				EntityType: "Job",
			},
		}
		//setup entity factory
		if swagger.Components.Schemas["Job"].Value == nil {
			return ctxt, fmt.Errorf("no schema '%s' is in the config", "Schema")
		}
		entityFactory := new(model.DefaultEntityFactory).FromSchema("Job", swagger.Components.Schemas["Job"].Value)
		ctxt = context.WithValue(ctxt, context2.ENTITY_FACTORY, entityFactory)

		err = dispatcher.Dispatch(ctxt, command, eventStore, projection, log)

		if err == nil {
			log.Debug("Job %s created", job["title"])
		}
	}

	return ctxt, nil
}

func DefaultEmployeesInitializer(ctxt context.Context, tapi rest.Container, swagger *openapi3.Swagger) (context.Context, error) {
	//get default logger
	log, err := tapi.GetLog("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default employees '%s'", err)
	}
	log.Infof("started setting up default employees")
	//event store
	eventStore, err := tapi.GetEventStore("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default employees '%s'", err)
	}
	//get default projection
	projection, err := tapi.GetProjection("Default")
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default employees '%s'", err)
	}
	dispatcher, err := tapi.GetCommandDispatcher("Default")

	//set up job entity factory
	if swagger.Components.Schemas["Job"].Value == nil {
		return ctxt, fmt.Errorf("no schema '%s' is in the config", "Schema")
	}
	jobEntityFactory := new(model.DefaultEntityFactory).FromSchema("Job", swagger.Components.Schemas["Job"].Value)
	jobContext := context.WithValue(ctxt, context2.ENTITY_FACTORY, jobEntityFactory)
	jobs, len, err := projection.GetList(jobContext, jobEntityFactory, 0, 0, "", nil, nil)
	if err != nil {
		return ctxt, fmt.Errorf("error setting up default employees '%s'", err)

	}
	if len != 4 {
		return ctxt, fmt.Errorf("error setting up default employees '%s'", err)
	}
	employees := []map[string]interface{}{
		{
			"firstName": "Bob",
			"lastName":  "Henry",
			"email":     "d@gmail.com",
			"dob":       "21st Jan 1980",
			"age":       20,
			"job":       jobs[0].ToMap(),
		},
		{
			"firstName": "John",
			"lastName":  "Doe",
			"email":     "johnDoe@gmail.com",
			"dob":       "31st Jan 1980",
			"age":       20,
			"job":       jobs[1].ToMap(),
		},
		{
			"firstName": "Jane",
			"lastName":  "Doe",
			"email":     "janeDoe@gmail.com",
			"dob":       "12th Jan 1980",
			"age":       20,
			"job":       jobs[2].ToMap(),
		},
		{
			"firstName": "John",
			"lastName":  "Paul",
			"email":     "johnPaul@gmail.com",
			"age":       20,
			"dob":       "4th Jan 1980",
			"job":       jobs[3].ToMap(),
		},
	}

	for _, employee := range employees {

		data, err := json.Marshal(employee)
		if err != nil {
			return ctxt, fmt.Errorf("error setting up default employees '%s'", err)
		}
		command := &model.Command{
			Type:    "create",
			Payload: data,
			Metadata: model.CommandMetadata{
				EntityType: "Employee",
			},
		}
		//setup entity factory
		if swagger.Components.Schemas["Employee"].Value == nil {
			return ctxt, fmt.Errorf("no schema '%s' is in the config", "Schema")
		}
		entityFactory := new(model.DefaultEntityFactory).FromSchema("Employee", swagger.Components.Schemas["Employee"].Value)
		ctxt = context.WithValue(ctxt, context2.ENTITY_FACTORY, entityFactory)

		err = dispatcher.Dispatch(ctxt, command, eventStore, projection, log)

		if err == nil {
			log.Debug("Employee %s created", employee["fistName"])
		}
	}

	return ctxt, nil
}
