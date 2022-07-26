package src_test

import (
	"context"
	"encoding/json"
	"manager/src"
	"testing"

	"github.com/wepala/weos/model"
	"github.com/wepala/weos/projections"
)

func TestDefaultJobsInitializer(t *testing.T) {
	swagger, err := LoadConfig(t)
	if err != nil {
		t.Fatalf("error loading api specification '%s'", err)
	}

	t.Run("create default Jobs", func(t *testing.T) {

		ctxt := context.Background()
		commandDispatcher := &CommandDispatcherMock{DispatchFunc: func(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error {
			if command.Type != "create" {
				t.Errorf("expected the command to be '%s', got '%s'", "create", command.Type)
			}
			if command.Metadata.EntityType != "Job" {
				t.Errorf("expected the entity type to be '%s', got '%s'", "Job", command.Metadata.EntityType)
			}
			var payload map[string]interface{}
			err := json.Unmarshal(command.Payload, &payload)
			if err != nil {
				t.Fatalf("unexpected error reading the payload '%s'", err)
			}
			if _, ok := payload["title"]; !ok {
				t.Fatalf("expected the title to be set")
			}
			if _, ok := payload["description"]; !ok {
				t.Fatalf("expected the description to be set")
			}
			return nil
		}}
		container := &ContainerMock{
			GetCommandDispatcherFunc: func(name string) (model.CommandDispatcher, error) {
				if name == "Default" {
					return commandDispatcher, nil
				}

				return nil, nil
			},
			GetEventStoreFunc: func(name string) (model.EventRepository, error) {
				return nil, nil
			},
			GetProjectionFunc: func(name string) (projections.Projection, error) {
				return nil, nil
			},
			GetLogFunc: func(name string) (model.Log, error) {
				return &LogMock{
					InfofFunc: func(format string, args ...interface{}) {

					},
					DebugfFunc: func(format string, args ...interface{}) {

					},
					DebugFunc: func(args ...interface{}) {

					},
				}, nil
			},
		}
		//trigger initializer
		ctxt, err := src.DefaultJobsInitializer(ctxt, container, swagger)
		if err != nil {
			t.Fatalf("error encountered running initializer '%s'", err)
		}
		//check that a command is dispatched
		if len(container.GetCommandDispatcherCalls()) != 1 {
			t.Error("expected the command dispatcher for system to be retrieved")
		}

		if len(commandDispatcher.DispatchCalls()) != 4 {
			t.Errorf("expected the command dispatcher to be called %d time, called %d times", 1, len(commandDispatcher.DispatchCalls()))
		}
	})
}

func TestDefaultEmployeesInitializer(t *testing.T) {
	swagger, err := LoadConfig(t)
	if err != nil {
		t.Fatalf("error loading api specification '%s'", err)
	}

	t.Run("create default Employees", func(t *testing.T) {

		ctxt := context.Background()
		commandDispatcher := &CommandDispatcherMock{DispatchFunc: func(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error {
			if command.Type != "create" {
				t.Errorf("expected the command to be '%s', got '%s'", "create", command.Type)
			}
			if command.Metadata.EntityType != "Employee" {
				t.Errorf("expected the entity type to be '%s', got '%s'", "Employee", command.Metadata.EntityType)
			}
			var payload map[string]interface{}
			err := json.Unmarshal(command.Payload, &payload)
			if err != nil {
				t.Fatalf("unexpected error reading the payload '%s'", err)
			}
			if _, ok := payload["firstName"]; !ok {
				t.Fatalf("expected the firstName to be set")
			}
			if _, ok := payload["lastName"]; !ok {
				t.Fatalf("expected the lastName to be set")
			}
			return nil
		}}
		container := &ContainerMock{
			GetCommandDispatcherFunc: func(name string) (model.CommandDispatcher, error) {
				if name == "Default" {
					return commandDispatcher, nil
				}

				return nil, nil
			},
			GetEventStoreFunc: func(name string) (model.EventRepository, error) {
				return nil, nil
			},
			GetProjectionFunc: func(name string) (projections.Projection, error) {
				return nil, nil
			},
			GetLogFunc: func(name string) (model.Log, error) {
				return &LogMock{
					InfofFunc: func(format string, args ...interface{}) {

					},
					DebugfFunc: func(format string, args ...interface{}) {

					},
					DebugFunc: func(args ...interface{}) {

					},
				}, nil
			},
		}
		//trigger initializer
		ctxt, err := src.DefaultEmployeesInitializer(ctxt, container, swagger)
		if err != nil {
			t.Fatalf("error encountered running initializer '%s'", err)
		}
		//check that a command is dispatched
		if len(container.GetCommandDispatcherCalls()) != 1 {
			t.Error("expected the command dispatcher for system to be retrieved")
		}

		if len(commandDispatcher.DispatchCalls()) != 4 {
			t.Errorf("expected the command dispatcher to be called %d time, called %d times", 1, len(commandDispatcher.DispatchCalls()))
		}
	})
}
