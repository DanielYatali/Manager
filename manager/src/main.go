package src

import (
	"log"

	weos "github.com/wepala/weos/controllers/rest"
	"golang.org/x/net/context"
)

func Init(ctxt context.Context, specPath string) (*weos.RESTAPI, error) {
	//instantiate weos with a reference to the OpenAPI specification
	api, err := weos.New(specPath)
	if err != nil {
		log.Fatalln("error loading api config", err)
	}
	api.RegisterGlobalInitializer(weos.SQLDatabase)
	api.RegisterGlobalInitializer(weos.DefaultProjection)
	api.RegisterGlobalInitializer(weos.DefaultEventStore)
	// api.RegisterGlobalInitializer(DefaultJobsInitializer)
	// api.RegisterGlobalInitializer(DefaultEmployeesInitializer)

	// api.RegisterGlobalInitializer(MainThemeInitializer)
	// api.RegisterOperationInitializer(ViewDataInitializer)

	//register middleware

	api.RegisterController("CreateJob", CreateJob)

	//initialize API so that standard middleware,controller,projections etc are registered
	err = api.Initialize(ctxt)
	if err != nil {
		log.Fatalln("error initializing api", err)
	}
	dispatcher, err := api.GetCommandDispatcher("Default")
	if err != nil {
		log.Fatalln("error getting dispatcher", err)
	}
	// dispatcher.AddSubscriber(CreateCommand(ctxt, CREATE_SCRIPT_COMMAND, ""), scriptReceiver.Create)

	api.RegisterCommandDispatcher("Default", dispatcher)

	return api, err
}
