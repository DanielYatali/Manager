package src_test

//import (
//	"manager/src"
//	"encoding/json"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/labstack/echo/v4"
//	"github.com/wepala/weos/model"
//	"golang.org/x/net/context"
//)
//
//func TestSystemComponentController_Delete(t *testing.T) {
//
//	container := &ContainerMock{
//		GetLogFunc: func(name string) (model.Log, error) {
//			return &LogMock{
//				DebugfFunc: func(format string, args ...interface{}) {
//
//				},
//				ErrorfFunc: func(format string, args ...interface{}) {
//
//				},
//			}, nil
//		},
//	}
//
//	dispatcher := &CommandDispatcherMock{
//		DispatchFunc: func(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error {
//			if command.Type != src.DELETE_SCRIPT_COMMAND {
//				t.Errorf("expected the script command to be '%s', got '%s'", src.DELETE_SCRIPT_COMMAND, command.Type)
//			}
//
//			payload := make(map[string]string)
//			json.Unmarshal(command.Payload, &payload)
//			if _, ok := payload["title"]; !ok {
//				t.Fatalf("expected title in the payload")
//			}
//
//			if payload["title"] != "foo" {
//				t.Errorf("expected the title to be '%s', got '%s'", "test", payload["title"])
//			}
//			return nil
//		},
//	}
//	handler := fileMiddleware.Delete(container, &ProjectionMock{}, dispatcher, &EventRepositoryMock{}, &EntityFactoryMock{})
//	e := echo.New()
//	resp := httptest.NewRecorder()
//	req := httptest.NewRequest(http.MethodDelete, "/scripts/foo", nil)
//	e.DELETE("/scripts/foo", handler, func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(ctxt echo.Context) error {
//			newContext := context.WithValue(ctxt.Request().Context(), "componentType", "script")
//			newContext = context.WithValue(newContext, "id", "foo")
//			request := ctxt.Request().WithContext(newContext)
//			ctxt.SetRequest(request)
//			return next(ctxt)
//		}
//	})
//	e.ServeHTTP(resp, req)
//
//	if resp.Code != http.StatusOK {
//		t.Errorf("expected the response code to be %d, got %d", http.StatusOK, resp.Code)
//	}
//
//	if len(dispatcher.DispatchCalls()) != 1 {
//		t.Errorf("expected the dispatch function to be called %d time, called %d times", 1, len(dispatcher.DispatchCalls()))
//	}
//
//	//check that success true is returned
//	var respPayload map[string]bool
//	err := json.NewDecoder(resp.Body).Decode(&respPayload)
//	if err != nil {
//		t.Error("expected the response to be a map[string]bool")
//	}
//
//	if success, ok := respPayload["success"]; !ok || !success {
//		t.Errorf("expected the response to have property 'success' and for it to be true")
//	}
//}
