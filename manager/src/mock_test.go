package src_test

import (
	"context"
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	ds "github.com/ompluscator/dynamic-struct"
	"github.com/wepala/weos/model"
	context2 "golang.org/x/net/context"
	"gorm.io/gorm"
)

// Ensure, that EventRepositoryMock does implement model.EventRepository.
// If this is not the case, regenerate this file with moq.
var _ model.EventRepository = &EventRepositoryMock{}

// EventRepositoryMock is a mock implementation of model.EventRepository.
//
// 	func TestSomethingThatUsesEventRepository(t *testing.T) {
//
// 		// make and configure a mocked model.EventRepository
// 		mockedEventRepository := &EventRepositoryMock{
// 			AddSubscriberFunc: func(handler model.EventHandler)  {
// 				panic("mock out the AddSubscriber method")
// 			},
// 			FlushFunc: func() error {
// 				panic("mock out the Flush method")
// 			},
// 			GetAggregateSequenceNumberFunc: func(ID string) (int64, error) {
// 				panic("mock out the GetAggregateSequenceNumber method")
// 			},
// 			GetByAggregateFunc: func(ID string) ([]*model.Event, error) {
// 				panic("mock out the GetByAggregate method")
// 			},
// 			GetByAggregateAndSequenceRangeFunc: func(ID string, start int64, end int64) ([]*model.Event, error) {
// 				panic("mock out the GetByAggregateAndSequenceRange method")
// 			},
// 			GetByAggregateAndTypeFunc: func(ID string, entityType string) ([]*model.Event, error) {
// 				panic("mock out the GetByAggregateAndType method")
// 			},
// 			GetByEntityAndAggregateFunc: func(entityID string, entityType string, rootID string) ([]*model.Event, error) {
// 				panic("mock out the GetByEntityAndAggregate method")
// 			},
// 			GetSubscribersFunc: func() ([]model.EventHandler, error) {
// 				panic("mock out the GetSubscribers method")
// 			},
// 			MigrateFunc: func(ctx context.Context) error {
// 				panic("mock out the Migrate method")
// 			},
// 			PersistFunc: func(ctxt context.Context, entity model.AggregateInterface) error {
// 				panic("mock out the Persist method")
// 			},
// 			ReplayEventsFunc: func(ctxt context.Context, date time.Time, entityFactories map[string]model.EntityFactory, projection model.Projection) (int, int, int, []error) {
// 				panic("mock out the ReplayEvents method")
// 			},
// 		}
//
// 		// use mockedEventRepository in code that requires model.EventRepository
// 		// and then make assertions.
//
// 	}
type EventRepositoryMock struct {
	// AddSubscriberFunc mocks the AddSubscriber method.
	AddSubscriberFunc func(handler model.EventHandler)

	// FlushFunc mocks the Flush method.
	FlushFunc func() error

	// GetAggregateSequenceNumberFunc mocks the GetAggregateSequenceNumber method.
	GetAggregateSequenceNumberFunc func(ID string) (int64, error)

	// GetByAggregateFunc mocks the GetByAggregate method.
	GetByAggregateFunc func(ID string) ([]*model.Event, error)

	// GetByAggregateAndSequenceRangeFunc mocks the GetByAggregateAndSequenceRange method.
	GetByAggregateAndSequenceRangeFunc func(ID string, start int64, end int64) ([]*model.Event, error)

	// GetByAggregateAndTypeFunc mocks the GetByAggregateAndType method.
	GetByAggregateAndTypeFunc func(ID string, entityType string) ([]*model.Event, error)

	// GetByEntityAndAggregateFunc mocks the GetByEntityAndAggregate method.
	GetByEntityAndAggregateFunc func(entityID string, entityType string, rootID string) ([]*model.Event, error)

	// GetSubscribersFunc mocks the GetSubscribers method.
	GetSubscribersFunc func() ([]model.EventHandler, error)

	// MigrateFunc mocks the Migrate method.
	MigrateFunc func(ctx context.Context) error

	// PersistFunc mocks the Persist method.
	PersistFunc func(ctxt context.Context, entity model.AggregateInterface) error

	// ReplayEventsFunc mocks the ReplayEvents method.
	ReplayEventsFunc func(ctxt context.Context, date time.Time, entityFactories map[string]model.EntityFactory, projection model.Projection) (int, int, int, []error)

	// calls tracks calls to the methods.
	calls struct {
		// AddSubscriber holds details about calls to the AddSubscriber method.
		AddSubscriber []struct {
			// Handler is the handler argument value.
			Handler model.EventHandler
		}
		// Flush holds details about calls to the Flush method.
		Flush []struct {
		}
		// GetAggregateSequenceNumber holds details about calls to the GetAggregateSequenceNumber method.
		GetAggregateSequenceNumber []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetByAggregate holds details about calls to the GetByAggregate method.
		GetByAggregate []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetByAggregateAndSequenceRange holds details about calls to the GetByAggregateAndSequenceRange method.
		GetByAggregateAndSequenceRange []struct {
			// ID is the ID argument value.
			ID string
			// Start is the start argument value.
			Start int64
			// End is the end argument value.
			End int64
		}
		// GetByAggregateAndType holds details about calls to the GetByAggregateAndType method.
		GetByAggregateAndType []struct {
			// ID is the ID argument value.
			ID string
			// EntityType is the entityType argument value.
			EntityType string
		}
		// GetByEntityAndAggregate holds details about calls to the GetByEntityAndAggregate method.
		GetByEntityAndAggregate []struct {
			// EntityID is the entityID argument value.
			EntityID string
			// EntityType is the entityType argument value.
			EntityType string
			// RootID is the rootID argument value.
			RootID string
		}
		// GetSubscribers holds details about calls to the GetSubscribers method.
		GetSubscribers []struct {
		}
		// Migrate holds details about calls to the Migrate method.
		Migrate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Persist holds details about calls to the Persist method.
		Persist []struct {
			// Ctxt is the ctxt argument value.
			Ctxt context.Context
			// Entity is the entity argument value.
			Entity model.AggregateInterface
		}
		// ReplayEvents holds details about calls to the ReplayEvents method.
		ReplayEvents []struct {
			// Ctxt is the ctxt argument value.
			Ctxt context.Context
			// Date is the date argument value.
			Date time.Time
			// EntityFactories is the entityFactories argument value.
			EntityFactories map[string]model.EntityFactory
			// Projection is the projection argument value.
			Projection model.Projection
		}
	}
	lockAddSubscriber                  sync.RWMutex
	lockFlush                          sync.RWMutex
	lockGetAggregateSequenceNumber     sync.RWMutex
	lockGetByAggregate                 sync.RWMutex
	lockGetByAggregateAndSequenceRange sync.RWMutex
	lockGetByAggregateAndType          sync.RWMutex
	lockGetByEntityAndAggregate        sync.RWMutex
	lockGetSubscribers                 sync.RWMutex
	lockMigrate                        sync.RWMutex
	lockPersist                        sync.RWMutex
	lockReplayEvents                   sync.RWMutex
}

// AddSubscriber calls AddSubscriberFunc.
func (mock *EventRepositoryMock) AddSubscriber(handler model.EventHandler) {
	if mock.AddSubscriberFunc == nil {
		panic("EventRepositoryMock.AddSubscriberFunc: method is nil but EventRepository.AddSubscriber was just called")
	}
	callInfo := struct {
		Handler model.EventHandler
	}{
		Handler: handler,
	}
	mock.lockAddSubscriber.Lock()
	mock.calls.AddSubscriber = append(mock.calls.AddSubscriber, callInfo)
	mock.lockAddSubscriber.Unlock()
	mock.AddSubscriberFunc(handler)
}

// AddSubscriberCalls gets all the calls that were made to AddSubscriber.
// Check the length with:
//     len(mockedEventRepository.AddSubscriberCalls())
func (mock *EventRepositoryMock) AddSubscriberCalls() []struct {
	Handler model.EventHandler
} {
	var calls []struct {
		Handler model.EventHandler
	}
	mock.lockAddSubscriber.RLock()
	calls = mock.calls.AddSubscriber
	mock.lockAddSubscriber.RUnlock()
	return calls
}

// Flush calls FlushFunc.
func (mock *EventRepositoryMock) Flush() error {
	if mock.FlushFunc == nil {
		panic("EventRepositoryMock.FlushFunc: method is nil but EventRepository.Flush was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFlush.Lock()
	mock.calls.Flush = append(mock.calls.Flush, callInfo)
	mock.lockFlush.Unlock()
	return mock.FlushFunc()
}

// FlushCalls gets all the calls that were made to Flush.
// Check the length with:
//     len(mockedEventRepository.FlushCalls())
func (mock *EventRepositoryMock) FlushCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFlush.RLock()
	calls = mock.calls.Flush
	mock.lockFlush.RUnlock()
	return calls
}

// GetAggregateSequenceNumber calls GetAggregateSequenceNumberFunc.
func (mock *EventRepositoryMock) GetAggregateSequenceNumber(ID string) (int64, error) {
	if mock.GetAggregateSequenceNumberFunc == nil {
		panic("EventRepositoryMock.GetAggregateSequenceNumberFunc: method is nil but EventRepository.GetAggregateSequenceNumber was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	mock.lockGetAggregateSequenceNumber.Lock()
	mock.calls.GetAggregateSequenceNumber = append(mock.calls.GetAggregateSequenceNumber, callInfo)
	mock.lockGetAggregateSequenceNumber.Unlock()
	return mock.GetAggregateSequenceNumberFunc(ID)
}

// GetAggregateSequenceNumberCalls gets all the calls that were made to GetAggregateSequenceNumber.
// Check the length with:
//     len(mockedEventRepository.GetAggregateSequenceNumberCalls())
func (mock *EventRepositoryMock) GetAggregateSequenceNumberCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetAggregateSequenceNumber.RLock()
	calls = mock.calls.GetAggregateSequenceNumber
	mock.lockGetAggregateSequenceNumber.RUnlock()
	return calls
}

// GetByAggregate calls GetByAggregateFunc.
func (mock *EventRepositoryMock) GetByAggregate(ID string) ([]*model.Event, error) {
	if mock.GetByAggregateFunc == nil {
		panic("EventRepositoryMock.GetByAggregateFunc: method is nil but EventRepository.GetByAggregate was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	mock.lockGetByAggregate.Lock()
	mock.calls.GetByAggregate = append(mock.calls.GetByAggregate, callInfo)
	mock.lockGetByAggregate.Unlock()
	return mock.GetByAggregateFunc(ID)
}

// GetByAggregateCalls gets all the calls that were made to GetByAggregate.
// Check the length with:
//     len(mockedEventRepository.GetByAggregateCalls())
func (mock *EventRepositoryMock) GetByAggregateCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetByAggregate.RLock()
	calls = mock.calls.GetByAggregate
	mock.lockGetByAggregate.RUnlock()
	return calls
}

// GetByAggregateAndSequenceRange calls GetByAggregateAndSequenceRangeFunc.
func (mock *EventRepositoryMock) GetByAggregateAndSequenceRange(ID string, start int64, end int64) ([]*model.Event, error) {
	if mock.GetByAggregateAndSequenceRangeFunc == nil {
		panic("EventRepositoryMock.GetByAggregateAndSequenceRangeFunc: method is nil but EventRepository.GetByAggregateAndSequenceRange was just called")
	}
	callInfo := struct {
		ID    string
		Start int64
		End   int64
	}{
		ID:    ID,
		Start: start,
		End:   end,
	}
	mock.lockGetByAggregateAndSequenceRange.Lock()
	mock.calls.GetByAggregateAndSequenceRange = append(mock.calls.GetByAggregateAndSequenceRange, callInfo)
	mock.lockGetByAggregateAndSequenceRange.Unlock()
	return mock.GetByAggregateAndSequenceRangeFunc(ID, start, end)
}

// GetByAggregateAndSequenceRangeCalls gets all the calls that were made to GetByAggregateAndSequenceRange.
// Check the length with:
//     len(mockedEventRepository.GetByAggregateAndSequenceRangeCalls())
func (mock *EventRepositoryMock) GetByAggregateAndSequenceRangeCalls() []struct {
	ID    string
	Start int64
	End   int64
} {
	var calls []struct {
		ID    string
		Start int64
		End   int64
	}
	mock.lockGetByAggregateAndSequenceRange.RLock()
	calls = mock.calls.GetByAggregateAndSequenceRange
	mock.lockGetByAggregateAndSequenceRange.RUnlock()
	return calls
}

// GetByAggregateAndType calls GetByAggregateAndTypeFunc.
func (mock *EventRepositoryMock) GetByAggregateAndType(ID string, entityType string) ([]*model.Event, error) {
	if mock.GetByAggregateAndTypeFunc == nil {
		panic("EventRepositoryMock.GetByAggregateAndTypeFunc: method is nil but EventRepository.GetByAggregateAndType was just called")
	}
	callInfo := struct {
		ID         string
		EntityType string
	}{
		ID:         ID,
		EntityType: entityType,
	}
	mock.lockGetByAggregateAndType.Lock()
	mock.calls.GetByAggregateAndType = append(mock.calls.GetByAggregateAndType, callInfo)
	mock.lockGetByAggregateAndType.Unlock()
	return mock.GetByAggregateAndTypeFunc(ID, entityType)
}

// GetByAggregateAndTypeCalls gets all the calls that were made to GetByAggregateAndType.
// Check the length with:
//     len(mockedEventRepository.GetByAggregateAndTypeCalls())
func (mock *EventRepositoryMock) GetByAggregateAndTypeCalls() []struct {
	ID         string
	EntityType string
} {
	var calls []struct {
		ID         string
		EntityType string
	}
	mock.lockGetByAggregateAndType.RLock()
	calls = mock.calls.GetByAggregateAndType
	mock.lockGetByAggregateAndType.RUnlock()
	return calls
}

// GetByEntityAndAggregate calls GetByEntityAndAggregateFunc.
func (mock *EventRepositoryMock) GetByEntityAndAggregate(entityID string, entityType string, rootID string) ([]*model.Event, error) {
	if mock.GetByEntityAndAggregateFunc == nil {
		panic("EventRepositoryMock.GetByEntityAndAggregateFunc: method is nil but EventRepository.GetByEntityAndAggregate was just called")
	}
	callInfo := struct {
		EntityID   string
		EntityType string
		RootID     string
	}{
		EntityID:   entityID,
		EntityType: entityType,
		RootID:     rootID,
	}
	mock.lockGetByEntityAndAggregate.Lock()
	mock.calls.GetByEntityAndAggregate = append(mock.calls.GetByEntityAndAggregate, callInfo)
	mock.lockGetByEntityAndAggregate.Unlock()
	return mock.GetByEntityAndAggregateFunc(entityID, entityType, rootID)
}

// GetByEntityAndAggregateCalls gets all the calls that were made to GetByEntityAndAggregate.
// Check the length with:
//     len(mockedEventRepository.GetByEntityAndAggregateCalls())
func (mock *EventRepositoryMock) GetByEntityAndAggregateCalls() []struct {
	EntityID   string
	EntityType string
	RootID     string
} {
	var calls []struct {
		EntityID   string
		EntityType string
		RootID     string
	}
	mock.lockGetByEntityAndAggregate.RLock()
	calls = mock.calls.GetByEntityAndAggregate
	mock.lockGetByEntityAndAggregate.RUnlock()
	return calls
}

// GetSubscribers calls GetSubscribersFunc.
func (mock *EventRepositoryMock) GetSubscribers() ([]model.EventHandler, error) {
	if mock.GetSubscribersFunc == nil {
		panic("EventRepositoryMock.GetSubscribersFunc: method is nil but EventRepository.GetSubscribers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetSubscribers.Lock()
	mock.calls.GetSubscribers = append(mock.calls.GetSubscribers, callInfo)
	mock.lockGetSubscribers.Unlock()
	return mock.GetSubscribersFunc()
}

// GetSubscribersCalls gets all the calls that were made to GetSubscribers.
// Check the length with:
//     len(mockedEventRepository.GetSubscribersCalls())
func (mock *EventRepositoryMock) GetSubscribersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetSubscribers.RLock()
	calls = mock.calls.GetSubscribers
	mock.lockGetSubscribers.RUnlock()
	return calls
}

// Migrate calls MigrateFunc.
func (mock *EventRepositoryMock) Migrate(ctx context.Context) error {
	if mock.MigrateFunc == nil {
		panic("EventRepositoryMock.MigrateFunc: method is nil but EventRepository.Migrate was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockMigrate.Lock()
	mock.calls.Migrate = append(mock.calls.Migrate, callInfo)
	mock.lockMigrate.Unlock()
	return mock.MigrateFunc(ctx)
}

// MigrateCalls gets all the calls that were made to Migrate.
// Check the length with:
//     len(mockedEventRepository.MigrateCalls())
func (mock *EventRepositoryMock) MigrateCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockMigrate.RLock()
	calls = mock.calls.Migrate
	mock.lockMigrate.RUnlock()
	return calls
}

// Persist calls PersistFunc.
func (mock *EventRepositoryMock) Persist(ctxt context.Context, entity model.AggregateInterface) error {
	if mock.PersistFunc == nil {
		panic("EventRepositoryMock.PersistFunc: method is nil but EventRepository.Persist was just called")
	}
	callInfo := struct {
		Ctxt   context.Context
		Entity model.AggregateInterface
	}{
		Ctxt:   ctxt,
		Entity: entity,
	}
	mock.lockPersist.Lock()
	mock.calls.Persist = append(mock.calls.Persist, callInfo)
	mock.lockPersist.Unlock()
	return mock.PersistFunc(ctxt, entity)
}

// PersistCalls gets all the calls that were made to Persist.
// Check the length with:
//     len(mockedEventRepository.PersistCalls())
func (mock *EventRepositoryMock) PersistCalls() []struct {
	Ctxt   context.Context
	Entity model.AggregateInterface
} {
	var calls []struct {
		Ctxt   context.Context
		Entity model.AggregateInterface
	}
	mock.lockPersist.RLock()
	calls = mock.calls.Persist
	mock.lockPersist.RUnlock()
	return calls
}

// ReplayEvents calls ReplayEventsFunc.
func (mock *EventRepositoryMock) ReplayEvents(ctxt context2.Context, date time.Time, entityFactories map[string]model.EntityFactory, projection model.Projection, schema *openapi3.Swagger) (int, int, int, []error) {
	if mock.ReplayEventsFunc == nil {
		panic("EventRepositoryMock.ReplayEventsFunc: method is nil but EventRepository.ReplayEvents was just called")
	}
	callInfo := struct {
		Ctxt            context.Context
		Date            time.Time
		EntityFactories map[string]model.EntityFactory
		Projection      model.Projection
	}{
		Ctxt:            ctxt,
		Date:            date,
		EntityFactories: entityFactories,
		Projection:      projection,
	}
	mock.lockReplayEvents.Lock()
	mock.calls.ReplayEvents = append(mock.calls.ReplayEvents, callInfo)
	mock.lockReplayEvents.Unlock()
	return mock.ReplayEventsFunc(ctxt, date, entityFactories, projection)
}

// ReplayEventsCalls gets all the calls that were made to ReplayEvents.
// Check the length with:
//     len(mockedEventRepository.ReplayEventsCalls())
func (mock *EventRepositoryMock) ReplayEventsCalls() []struct {
	Ctxt            context.Context
	Date            time.Time
	EntityFactories map[string]model.EntityFactory
	Projection      model.Projection
} {
	var calls []struct {
		Ctxt            context.Context
		Date            time.Time
		EntityFactories map[string]model.EntityFactory
		Projection      model.Projection
	}
	mock.lockReplayEvents.RLock()
	calls = mock.calls.ReplayEvents
	mock.lockReplayEvents.RUnlock()
	return calls
}

// Ensure, that ProjectionMock does implement model.Projection.
// If this is not the case, regenerate this file with moq.
var _ model.Projection = &ProjectionMock{}

// ProjectionMock is a mock implementation of model.Projection.
//
// 	func TestSomethingThatUsesProjection(t *testing.T) {
//
// 		// make and configure a mocked model.Projection
// 		mockedProjection := &ProjectionMock{
// 			GetByEntityIDFunc: func(ctxt context.Context, entityFactory model.EntityFactory, id string) (map[string]interface{}, error) {
// 				panic("mock out the GetByEntityID method")
// 			},
// 			GetByKeyFunc: func(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) (*model.ContentEntity, error) {
// 				panic("mock out the GetByKey method")
// 			},
// 			GetByPropertiesFunc: func(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) ([]*model.ContentEntity, error) {
// 				panic("mock out the GetByProperties method")
// 			},
// 			GetContentEntitiesFunc: func(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]map[string]interface{}, int64, error) {
// 				panic("mock out the GetContentEntities method")
// 			},
// 			GetContentEntityFunc: func(ctx context.Context, entityFactory model.EntityFactory, weosID string) (*model.ContentEntity, error) {
// 				panic("mock out the GetContentEntity method")
// 			},
// 			GetEventHandlerFunc: func() model.EventHandler {
// 				panic("mock out the GetEventHandler method")
// 			},
// 			GetListFunc: func(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]*model.ContentEntity, int64, error) {
// 				panic("mock out the GetList method")
// 			},
// 			MigrateFunc: func(ctx context.Context, builders map[string]ds.Builder, deletedFields map[string][]string) error {
// 				panic("mock out the Migrate method")
// 			},
// 		}
//
// 		// use mockedProjection in code that requires model.Projection
// 		// and then make assertions.
//
// 	}
type ProjectionMock struct {
	// GetByEntityIDFunc mocks the GetByEntityID method.
	GetByEntityIDFunc func(ctxt context.Context, entityFactory model.EntityFactory, id string) (map[string]interface{}, error)

	// GetByKeyFunc mocks the GetByKey method.
	GetByKeyFunc func(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) (*model.ContentEntity, error)

	// GetByPropertiesFunc mocks the GetByProperties method.
	GetByPropertiesFunc func(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) ([]*model.ContentEntity, error)

	// GetContentEntitiesFunc mocks the GetContentEntities method.
	GetContentEntitiesFunc func(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]map[string]interface{}, int64, error)

	// GetContentEntityFunc mocks the GetContentEntity method.
	GetContentEntityFunc func(ctx context.Context, entityFactory model.EntityFactory, weosID string) (*model.ContentEntity, error)

	// GetEventHandlerFunc mocks the GetEventHandler method.
	GetEventHandlerFunc func() model.EventHandler

	// GetListFunc mocks the GetList method.
	GetListFunc func(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]*model.ContentEntity, int64, error)

	// MigrateFunc mocks the Migrate method.
	MigrateFunc func(ctx context.Context, schema *openapi3.Swagger) error

	// calls tracks calls to the methods.
	calls struct {
		// GetByEntityID holds details about calls to the GetByEntityID method.
		GetByEntityID []struct {
			// Ctxt is the ctxt argument value.
			Ctxt context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// ID is the id argument value.
			ID string
		}
		// GetByKey holds details about calls to the GetByKey method.
		GetByKey []struct {
			// Ctxt is the ctxt argument value.
			Ctxt context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// Identifiers is the identifiers argument value.
			Identifiers map[string]interface{}
		}
		// GetByProperties holds details about calls to the GetByProperties method.
		GetByProperties []struct {
			// Ctxt is the ctxt argument value.
			Ctxt context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// Identifiers is the identifiers argument value.
			Identifiers map[string]interface{}
		}
		// GetContentEntities holds details about calls to the GetContentEntities method.
		GetContentEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// Page is the page argument value.
			Page int
			// Limit is the limit argument value.
			Limit int
			// Query is the query argument value.
			Query string
			// SortOptions is the sortOptions argument value.
			SortOptions map[string]string
			// FilterOptions is the filterOptions argument value.
			FilterOptions map[string]interface{}
		}
		// GetContentEntity holds details about calls to the GetContentEntity method.
		GetContentEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// WeosID is the weosID argument value.
			WeosID string
		}
		// GetEventHandler holds details about calls to the GetEventHandler method.
		GetEventHandler []struct {
		}
		// GetList holds details about calls to the GetList method.
		GetList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityFactory is the entityFactory argument value.
			EntityFactory model.EntityFactory
			// Page is the page argument value.
			Page int
			// Limit is the limit argument value.
			Limit int
			// Query is the query argument value.
			Query string
			// SortOptions is the sortOptions argument value.
			SortOptions map[string]string
			// FilterOptions is the filterOptions argument value.
			FilterOptions map[string]interface{}
		}
		// Migrate holds details about calls to the Migrate method.
		Migrate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Schema is the schema argument value.
			Schema *openapi3.Swagger
		}
	}
	lockGetByEntityID      sync.RWMutex
	lockGetByKey           sync.RWMutex
	lockGetByProperties    sync.RWMutex
	lockGetContentEntities sync.RWMutex
	lockGetContentEntity   sync.RWMutex
	lockGetEventHandler    sync.RWMutex
	lockGetList            sync.RWMutex
	lockMigrate            sync.RWMutex
}

// GetByEntityID calls GetByEntityIDFunc.
func (mock *ProjectionMock) GetByEntityID(ctxt context.Context, entityFactory model.EntityFactory, id string) (map[string]interface{}, error) {
	if mock.GetByEntityIDFunc == nil {
		panic("ProjectionMock.GetByEntityIDFunc: method is nil but Projection.GetByEntityID was just called")
	}
	callInfo := struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		ID            string
	}{
		Ctxt:          ctxt,
		EntityFactory: entityFactory,
		ID:            id,
	}
	mock.lockGetByEntityID.Lock()
	mock.calls.GetByEntityID = append(mock.calls.GetByEntityID, callInfo)
	mock.lockGetByEntityID.Unlock()
	return mock.GetByEntityIDFunc(ctxt, entityFactory, id)
}

// GetByEntityIDCalls gets all the calls that were made to GetByEntityID.
// Check the length with:
//     len(mockedProjection.GetByEntityIDCalls())
func (mock *ProjectionMock) GetByEntityIDCalls() []struct {
	Ctxt          context.Context
	EntityFactory model.EntityFactory
	ID            string
} {
	var calls []struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		ID            string
	}
	mock.lockGetByEntityID.RLock()
	calls = mock.calls.GetByEntityID
	mock.lockGetByEntityID.RUnlock()
	return calls
}

// GetByKey calls GetByKeyFunc.
func (mock *ProjectionMock) GetByKey(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) (*model.ContentEntity, error) {
	if mock.GetByKeyFunc == nil {
		panic("ProjectionMock.GetByKeyFunc: method is nil but Projection.GetByKey was just called")
	}
	callInfo := struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		Identifiers   map[string]interface{}
	}{
		Ctxt:          ctxt,
		EntityFactory: entityFactory,
		Identifiers:   identifiers,
	}
	mock.lockGetByKey.Lock()
	mock.calls.GetByKey = append(mock.calls.GetByKey, callInfo)
	mock.lockGetByKey.Unlock()
	return mock.GetByKeyFunc(ctxt, entityFactory, identifiers)
}

// GetByKeyCalls gets all the calls that were made to GetByKey.
// Check the length with:
//     len(mockedProjection.GetByKeyCalls())
func (mock *ProjectionMock) GetByKeyCalls() []struct {
	Ctxt          context.Context
	EntityFactory model.EntityFactory
	Identifiers   map[string]interface{}
} {
	var calls []struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		Identifiers   map[string]interface{}
	}
	mock.lockGetByKey.RLock()
	calls = mock.calls.GetByKey
	mock.lockGetByKey.RUnlock()
	return calls
}

// GetByProperties calls GetByPropertiesFunc.
func (mock *ProjectionMock) GetByProperties(ctxt context.Context, entityFactory model.EntityFactory, identifiers map[string]interface{}) ([]*model.ContentEntity, error) {
	if mock.GetByPropertiesFunc == nil {
		panic("ProjectionMock.GetByPropertiesFunc: method is nil but Projection.GetByProperties was just called")
	}
	callInfo := struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		Identifiers   map[string]interface{}
	}{
		Ctxt:          ctxt,
		EntityFactory: entityFactory,
		Identifiers:   identifiers,
	}
	mock.lockGetByProperties.Lock()
	mock.calls.GetByProperties = append(mock.calls.GetByProperties, callInfo)
	mock.lockGetByProperties.Unlock()
	return mock.GetByPropertiesFunc(ctxt, entityFactory, identifiers)
}

// GetByPropertiesCalls gets all the calls that were made to GetByProperties.
// Check the length with:
//     len(mockedProjection.GetByPropertiesCalls())
func (mock *ProjectionMock) GetByPropertiesCalls() []struct {
	Ctxt          context.Context
	EntityFactory model.EntityFactory
	Identifiers   map[string]interface{}
} {
	var calls []struct {
		Ctxt          context.Context
		EntityFactory model.EntityFactory
		Identifiers   map[string]interface{}
	}
	mock.lockGetByProperties.RLock()
	calls = mock.calls.GetByProperties
	mock.lockGetByProperties.RUnlock()
	return calls
}

// GetContentEntities calls GetContentEntitiesFunc.
func (mock *ProjectionMock) GetContentEntities(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]map[string]interface{}, int64, error) {
	if mock.GetContentEntitiesFunc == nil {
		panic("ProjectionMock.GetContentEntitiesFunc: method is nil but Projection.GetContentEntities was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		Page          int
		Limit         int
		Query         string
		SortOptions   map[string]string
		FilterOptions map[string]interface{}
	}{
		Ctx:           ctx,
		EntityFactory: entityFactory,
		Page:          page,
		Limit:         limit,
		Query:         query,
		SortOptions:   sortOptions,
		FilterOptions: filterOptions,
	}
	mock.lockGetContentEntities.Lock()
	mock.calls.GetContentEntities = append(mock.calls.GetContentEntities, callInfo)
	mock.lockGetContentEntities.Unlock()
	return mock.GetContentEntitiesFunc(ctx, entityFactory, page, limit, query, sortOptions, filterOptions)
}

// GetContentEntitiesCalls gets all the calls that were made to GetContentEntities.
// Check the length with:
//     len(mockedProjection.GetContentEntitiesCalls())
func (mock *ProjectionMock) GetContentEntitiesCalls() []struct {
	Ctx           context.Context
	EntityFactory model.EntityFactory
	Page          int
	Limit         int
	Query         string
	SortOptions   map[string]string
	FilterOptions map[string]interface{}
} {
	var calls []struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		Page          int
		Limit         int
		Query         string
		SortOptions   map[string]string
		FilterOptions map[string]interface{}
	}
	mock.lockGetContentEntities.RLock()
	calls = mock.calls.GetContentEntities
	mock.lockGetContentEntities.RUnlock()
	return calls
}

// GetContentEntity calls GetContentEntityFunc.
func (mock *ProjectionMock) GetContentEntity(ctx context.Context, entityFactory model.EntityFactory, weosID string) (*model.ContentEntity, error) {
	if mock.GetContentEntityFunc == nil {
		panic("ProjectionMock.GetContentEntityFunc: method is nil but Projection.GetContentEntity was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		WeosID        string
	}{
		Ctx:           ctx,
		EntityFactory: entityFactory,
		WeosID:        weosID,
	}
	mock.lockGetContentEntity.Lock()
	mock.calls.GetContentEntity = append(mock.calls.GetContentEntity, callInfo)
	mock.lockGetContentEntity.Unlock()
	return mock.GetContentEntityFunc(ctx, entityFactory, weosID)
}

// GetContentEntityCalls gets all the calls that were made to GetContentEntity.
// Check the length with:
//     len(mockedProjection.GetContentEntityCalls())
func (mock *ProjectionMock) GetContentEntityCalls() []struct {
	Ctx           context.Context
	EntityFactory model.EntityFactory
	WeosID        string
} {
	var calls []struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		WeosID        string
	}
	mock.lockGetContentEntity.RLock()
	calls = mock.calls.GetContentEntity
	mock.lockGetContentEntity.RUnlock()
	return calls
}

// GetEventHandler calls GetEventHandlerFunc.
func (mock *ProjectionMock) GetEventHandler() model.EventHandler {
	if mock.GetEventHandlerFunc == nil {
		panic("ProjectionMock.GetEventHandlerFunc: method is nil but Projection.GetEventHandler was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetEventHandler.Lock()
	mock.calls.GetEventHandler = append(mock.calls.GetEventHandler, callInfo)
	mock.lockGetEventHandler.Unlock()
	return mock.GetEventHandlerFunc()
}

// GetEventHandlerCalls gets all the calls that were made to GetEventHandler.
// Check the length with:
//     len(mockedProjection.GetEventHandlerCalls())
func (mock *ProjectionMock) GetEventHandlerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetEventHandler.RLock()
	calls = mock.calls.GetEventHandler
	mock.lockGetEventHandler.RUnlock()
	return calls
}

// GetList calls GetListFunc.
func (mock *ProjectionMock) GetList(ctx context.Context, entityFactory model.EntityFactory, page int, limit int, query string, sortOptions map[string]string, filterOptions map[string]interface{}) ([]*model.ContentEntity, int64, error) {
	if mock.GetListFunc == nil {
		panic("ProjectionMock.GetListFunc: method is nil but Projection.GetList was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		Page          int
		Limit         int
		Query         string
		SortOptions   map[string]string
		FilterOptions map[string]interface{}
	}{
		Ctx:           ctx,
		EntityFactory: entityFactory,
		Page:          page,
		Limit:         limit,
		Query:         query,
		SortOptions:   sortOptions,
		FilterOptions: filterOptions,
	}
	mock.lockGetList.Lock()
	mock.calls.GetList = append(mock.calls.GetList, callInfo)
	mock.lockGetList.Unlock()
	return mock.GetListFunc(ctx, entityFactory, page, limit, query, sortOptions, filterOptions)
}

// GetListCalls gets all the calls that were made to GetList.
// Check the length with:
//     len(mockedProjection.GetListCalls())
func (mock *ProjectionMock) GetListCalls() []struct {
	Ctx           context.Context
	EntityFactory model.EntityFactory
	Page          int
	Limit         int
	Query         string
	SortOptions   map[string]string
	FilterOptions map[string]interface{}
} {
	var calls []struct {
		Ctx           context.Context
		EntityFactory model.EntityFactory
		Page          int
		Limit         int
		Query         string
		SortOptions   map[string]string
		FilterOptions map[string]interface{}
	}
	mock.lockGetList.RLock()
	calls = mock.calls.GetList
	mock.lockGetList.RUnlock()
	return calls
}

// Migrate calls MigrateFunc.
func (mock *ProjectionMock) Migrate(ctx context2.Context, schema *openapi3.Swagger) error {
	if mock.MigrateFunc == nil {
		panic("ProjectionMock.MigrateFunc: method is nil but Projection.Migrate was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Schema *openapi3.Swagger
	}{
		Ctx:    ctx,
		Schema: schema,
	}
	mock.lockMigrate.Lock()
	mock.calls.Migrate = append(mock.calls.Migrate, callInfo)
	mock.lockMigrate.Unlock()
	return mock.MigrateFunc(ctx, schema)
}

// MigrateCalls gets all the calls that were made to Migrate.
// Check the length with:
//     len(mockedProjection.MigrateCalls())
func (mock *ProjectionMock) MigrateCalls() []struct {
	Ctx    context.Context
	Schema *openapi3.Swagger
} {
	var calls []struct {
		Ctx    context.Context
		Schema *openapi3.Swagger
	}
	mock.lockMigrate.RLock()
	calls = mock.calls.Migrate
	mock.lockMigrate.RUnlock()
	return calls
}

// Ensure, that LogMock does implement model.Log.
// If this is not the case, regenerate this file with moq.
var _ model.Log = &LogMock{}

// LogMock is a mock implementation of model.Log.
//
// 	func TestSomethingThatUsesLog(t *testing.T) {
//
// 		// make and configure a mocked model.Log
// 		mockedLog := &LogMock{
// 			DebugFunc: func(args ...interface{})  {
// 				panic("mock out the Debug method")
// 			},
// 			DebugfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Debugf method")
// 			},
// 			ErrorFunc: func(args ...interface{})  {
// 				panic("mock out the Error method")
// 			},
// 			ErrorfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Errorf method")
// 			},
// 			FatalFunc: func(args ...interface{})  {
// 				panic("mock out the Fatal method")
// 			},
// 			FatalfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Fatalf method")
// 			},
// 			InfoFunc: func(args ...interface{})  {
// 				panic("mock out the Info method")
// 			},
// 			InfofFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Infof method")
// 			},
// 			PanicFunc: func(args ...interface{})  {
// 				panic("mock out the Panic method")
// 			},
// 			PanicfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Panicf method")
// 			},
// 			PrintFunc: func(args ...interface{})  {
// 				panic("mock out the Print method")
// 			},
// 			PrintfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Printf method")
// 			},
// 		}
//
// 		// use mockedLog in code that requires model.Log
// 		// and then make assertions.
//
// 	}
type LogMock struct {
	// DebugFunc mocks the Debug method.
	DebugFunc func(args ...interface{})

	// DebugfFunc mocks the Debugf method.
	DebugfFunc func(format string, args ...interface{})

	// ErrorFunc mocks the Error method.
	ErrorFunc func(args ...interface{})

	// ErrorfFunc mocks the Errorf method.
	ErrorfFunc func(format string, args ...interface{})

	// FatalFunc mocks the Fatal method.
	FatalFunc func(args ...interface{})

	// FatalfFunc mocks the Fatalf method.
	FatalfFunc func(format string, args ...interface{})

	// InfoFunc mocks the Info method.
	InfoFunc func(args ...interface{})

	// InfofFunc mocks the Infof method.
	InfofFunc func(format string, args ...interface{})

	// PanicFunc mocks the Panic method.
	PanicFunc func(args ...interface{})

	// PanicfFunc mocks the Panicf method.
	PanicfFunc func(format string, args ...interface{})

	// PrintFunc mocks the Print method.
	PrintFunc func(args ...interface{})

	// PrintfFunc mocks the Printf method.
	PrintfFunc func(format string, args ...interface{})

	// calls tracks calls to the methods.
	calls struct {
		// Debug holds details about calls to the Debug method.
		Debug []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Debugf holds details about calls to the Debugf method.
		Debugf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Error holds details about calls to the Error method.
		Error []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Errorf holds details about calls to the Errorf method.
		Errorf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Fatal holds details about calls to the Fatal method.
		Fatal []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Fatalf holds details about calls to the Fatalf method.
		Fatalf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Info holds details about calls to the Info method.
		Info []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Infof holds details about calls to the Infof method.
		Infof []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Panic holds details about calls to the Panic method.
		Panic []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Panicf holds details about calls to the Panicf method.
		Panicf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
		// Print holds details about calls to the Print method.
		Print []struct {
			// Args is the args argument value.
			Args []interface{}
		}
		// Printf holds details about calls to the Printf method.
		Printf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
	}
	lockDebug  sync.RWMutex
	lockDebugf sync.RWMutex
	lockError  sync.RWMutex
	lockErrorf sync.RWMutex
	lockFatal  sync.RWMutex
	lockFatalf sync.RWMutex
	lockInfo   sync.RWMutex
	lockInfof  sync.RWMutex
	lockPanic  sync.RWMutex
	lockPanicf sync.RWMutex
	lockPrint  sync.RWMutex
	lockPrintf sync.RWMutex
}

// Debug calls DebugFunc.
func (mock *LogMock) Debug(args ...interface{}) {
	if mock.DebugFunc == nil {
		panic("LogMock.DebugFunc: method is nil but Log.Debug was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockDebug.Lock()
	mock.calls.Debug = append(mock.calls.Debug, callInfo)
	mock.lockDebug.Unlock()
	mock.DebugFunc(args...)
}

// DebugCalls gets all the calls that were made to Debug.
// Check the length with:
//     len(mockedLog.DebugCalls())
func (mock *LogMock) DebugCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockDebug.RLock()
	calls = mock.calls.Debug
	mock.lockDebug.RUnlock()
	return calls
}

// Debugf calls DebugfFunc.
func (mock *LogMock) Debugf(format string, args ...interface{}) {
	if mock.DebugfFunc == nil {
		panic("LogMock.DebugfFunc: method is nil but Log.Debugf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockDebugf.Lock()
	mock.calls.Debugf = append(mock.calls.Debugf, callInfo)
	mock.lockDebugf.Unlock()
	mock.DebugfFunc(format, args...)
}

// DebugfCalls gets all the calls that were made to Debugf.
// Check the length with:
//     len(mockedLog.DebugfCalls())
func (mock *LogMock) DebugfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockDebugf.RLock()
	calls = mock.calls.Debugf
	mock.lockDebugf.RUnlock()
	return calls
}

// Error calls ErrorFunc.
func (mock *LogMock) Error(args ...interface{}) {
	if mock.ErrorFunc == nil {
		panic("LogMock.ErrorFunc: method is nil but Log.Error was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	mock.ErrorFunc(args...)
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedLog.ErrorCalls())
func (mock *LogMock) ErrorCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// Errorf calls ErrorfFunc.
func (mock *LogMock) Errorf(format string, args ...interface{}) {
	if mock.ErrorfFunc == nil {
		panic("LogMock.ErrorfFunc: method is nil but Log.Errorf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockErrorf.Lock()
	mock.calls.Errorf = append(mock.calls.Errorf, callInfo)
	mock.lockErrorf.Unlock()
	mock.ErrorfFunc(format, args...)
}

// ErrorfCalls gets all the calls that were made to Errorf.
// Check the length with:
//     len(mockedLog.ErrorfCalls())
func (mock *LogMock) ErrorfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockErrorf.RLock()
	calls = mock.calls.Errorf
	mock.lockErrorf.RUnlock()
	return calls
}

// Fatal calls FatalFunc.
func (mock *LogMock) Fatal(args ...interface{}) {
	if mock.FatalFunc == nil {
		panic("LogMock.FatalFunc: method is nil but Log.Fatal was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockFatal.Lock()
	mock.calls.Fatal = append(mock.calls.Fatal, callInfo)
	mock.lockFatal.Unlock()
	mock.FatalFunc(args...)
}

// FatalCalls gets all the calls that were made to Fatal.
// Check the length with:
//     len(mockedLog.FatalCalls())
func (mock *LogMock) FatalCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockFatal.RLock()
	calls = mock.calls.Fatal
	mock.lockFatal.RUnlock()
	return calls
}

// Fatalf calls FatalfFunc.
func (mock *LogMock) Fatalf(format string, args ...interface{}) {
	if mock.FatalfFunc == nil {
		panic("LogMock.FatalfFunc: method is nil but Log.Fatalf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockFatalf.Lock()
	mock.calls.Fatalf = append(mock.calls.Fatalf, callInfo)
	mock.lockFatalf.Unlock()
	mock.FatalfFunc(format, args...)
}

// FatalfCalls gets all the calls that were made to Fatalf.
// Check the length with:
//     len(mockedLog.FatalfCalls())
func (mock *LogMock) FatalfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockFatalf.RLock()
	calls = mock.calls.Fatalf
	mock.lockFatalf.RUnlock()
	return calls
}

// Info calls InfoFunc.
func (mock *LogMock) Info(args ...interface{}) {
	if mock.InfoFunc == nil {
		panic("LogMock.InfoFunc: method is nil but Log.Info was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockInfo.Lock()
	mock.calls.Info = append(mock.calls.Info, callInfo)
	mock.lockInfo.Unlock()
	mock.InfoFunc(args...)
}

// InfoCalls gets all the calls that were made to Info.
// Check the length with:
//     len(mockedLog.InfoCalls())
func (mock *LogMock) InfoCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockInfo.RLock()
	calls = mock.calls.Info
	mock.lockInfo.RUnlock()
	return calls
}

// Infof calls InfofFunc.
func (mock *LogMock) Infof(format string, args ...interface{}) {
	if mock.InfofFunc == nil {
		panic("LogMock.InfofFunc: method is nil but Log.Infof was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockInfof.Lock()
	mock.calls.Infof = append(mock.calls.Infof, callInfo)
	mock.lockInfof.Unlock()
	mock.InfofFunc(format, args...)
}

// InfofCalls gets all the calls that were made to Infof.
// Check the length with:
//     len(mockedLog.InfofCalls())
func (mock *LogMock) InfofCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockInfof.RLock()
	calls = mock.calls.Infof
	mock.lockInfof.RUnlock()
	return calls
}

// Panic calls PanicFunc.
func (mock *LogMock) Panic(args ...interface{}) {
	if mock.PanicFunc == nil {
		panic("LogMock.PanicFunc: method is nil but Log.Panic was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockPanic.Lock()
	mock.calls.Panic = append(mock.calls.Panic, callInfo)
	mock.lockPanic.Unlock()
	mock.PanicFunc(args...)
}

// PanicCalls gets all the calls that were made to Panic.
// Check the length with:
//     len(mockedLog.PanicCalls())
func (mock *LogMock) PanicCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockPanic.RLock()
	calls = mock.calls.Panic
	mock.lockPanic.RUnlock()
	return calls
}

// Panicf calls PanicfFunc.
func (mock *LogMock) Panicf(format string, args ...interface{}) {
	if mock.PanicfFunc == nil {
		panic("LogMock.PanicfFunc: method is nil but Log.Panicf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockPanicf.Lock()
	mock.calls.Panicf = append(mock.calls.Panicf, callInfo)
	mock.lockPanicf.Unlock()
	mock.PanicfFunc(format, args...)
}

// PanicfCalls gets all the calls that were made to Panicf.
// Check the length with:
//     len(mockedLog.PanicfCalls())
func (mock *LogMock) PanicfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockPanicf.RLock()
	calls = mock.calls.Panicf
	mock.lockPanicf.RUnlock()
	return calls
}

// Print calls PrintFunc.
func (mock *LogMock) Print(args ...interface{}) {
	if mock.PrintFunc == nil {
		panic("LogMock.PrintFunc: method is nil but Log.Print was just called")
	}
	callInfo := struct {
		Args []interface{}
	}{
		Args: args,
	}
	mock.lockPrint.Lock()
	mock.calls.Print = append(mock.calls.Print, callInfo)
	mock.lockPrint.Unlock()
	mock.PrintFunc(args...)
}

// PrintCalls gets all the calls that were made to Print.
// Check the length with:
//     len(mockedLog.PrintCalls())
func (mock *LogMock) PrintCalls() []struct {
	Args []interface{}
} {
	var calls []struct {
		Args []interface{}
	}
	mock.lockPrint.RLock()
	calls = mock.calls.Print
	mock.lockPrint.RUnlock()
	return calls
}

// Printf calls PrintfFunc.
func (mock *LogMock) Printf(format string, args ...interface{}) {
	if mock.PrintfFunc == nil {
		panic("LogMock.PrintfFunc: method is nil but Log.Printf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockPrintf.Lock()
	mock.calls.Printf = append(mock.calls.Printf, callInfo)
	mock.lockPrintf.Unlock()
	mock.PrintfFunc(format, args...)
}

// PrintfCalls gets all the calls that were made to Printf.
// Check the length with:
//     len(mockedLog.PrintfCalls())
func (mock *LogMock) PrintfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockPrintf.RLock()
	calls = mock.calls.Printf
	mock.lockPrintf.RUnlock()
	return calls
}

// Ensure, that CommandDispatcherMock does implement model.CommandDispatcher.
// If this is not the case, regenerate this file with moq.
var _ model.CommandDispatcher = &CommandDispatcherMock{}

// CommandDispatcherMock is a mock implementation of model.CommandDispatcher.
//
// 	func TestSomethingThatUsesCommandDispatcher(t *testing.T) {
//
// 		// make and configure a mocked model.CommandDispatcher
// 		mockedCommandDispatcher := &CommandDispatcherMock{
// 			AddSubscriberFunc: func(command *model.Command, handler model.CommandHandler) map[string][]model.CommandHandler {
// 				panic("mock out the AddSubscriber method")
// 			},
// 			DispatchFunc: func(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error {
// 				panic("mock out the Dispatch method")
// 			},
// 			GetSubscribersFunc: func() map[string][]model.CommandHandler {
// 				panic("mock out the GetSubscribers method")
// 			},
// 		}
//
// 		// use mockedCommandDispatcher in code that requires model.CommandDispatcher
// 		// and then make assertions.
//
// 	}
type CommandDispatcherMock struct {
	// AddSubscriberFunc mocks the AddSubscriber method.
	AddSubscriberFunc func(command *model.Command, handler model.CommandHandler) map[string][]model.CommandHandler

	// DispatchFunc mocks the Dispatch method.
	DispatchFunc func(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error

	// GetSubscribersFunc mocks the GetSubscribers method.
	GetSubscribersFunc func() map[string][]model.CommandHandler

	// calls tracks calls to the methods.
	calls struct {
		// AddSubscriber holds details about calls to the AddSubscriber method.
		AddSubscriber []struct {
			// Command is the command argument value.
			Command *model.Command
			// Handler is the handler argument value.
			Handler model.CommandHandler
		}
		// Dispatch holds details about calls to the Dispatch method.
		Dispatch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Command is the command argument value.
			Command *model.Command
			// EventStore is the eventStore argument value.
			EventStore model.EventRepository
			// Projection is the projection argument value.
			Projection model.Projection
			// Logger is the logger argument value.
			Logger model.Log
		}
		// GetSubscribers holds details about calls to the GetSubscribers method.
		GetSubscribers []struct {
		}
	}
	lockAddSubscriber  sync.RWMutex
	lockDispatch       sync.RWMutex
	lockGetSubscribers sync.RWMutex
}

// AddSubscriber calls AddSubscriberFunc.
func (mock *CommandDispatcherMock) AddSubscriber(command *model.Command, handler model.CommandHandler) map[string][]model.CommandHandler {
	if mock.AddSubscriberFunc == nil {
		panic("CommandDispatcherMock.AddSubscriberFunc: method is nil but CommandDispatcher.AddSubscriber was just called")
	}
	callInfo := struct {
		Command *model.Command
		Handler model.CommandHandler
	}{
		Command: command,
		Handler: handler,
	}
	mock.lockAddSubscriber.Lock()
	mock.calls.AddSubscriber = append(mock.calls.AddSubscriber, callInfo)
	mock.lockAddSubscriber.Unlock()
	return mock.AddSubscriberFunc(command, handler)
}

// AddSubscriberCalls gets all the calls that were made to AddSubscriber.
// Check the length with:
//     len(mockedCommandDispatcher.AddSubscriberCalls())
func (mock *CommandDispatcherMock) AddSubscriberCalls() []struct {
	Command *model.Command
	Handler model.CommandHandler
} {
	var calls []struct {
		Command *model.Command
		Handler model.CommandHandler
	}
	mock.lockAddSubscriber.RLock()
	calls = mock.calls.AddSubscriber
	mock.lockAddSubscriber.RUnlock()
	return calls
}

// Dispatch calls DispatchFunc.
func (mock *CommandDispatcherMock) Dispatch(ctx context.Context, command *model.Command, eventStore model.EventRepository, projection model.Projection, logger model.Log) error {
	if mock.DispatchFunc == nil {
		panic("CommandDispatcherMock.DispatchFunc: method is nil but CommandDispatcher.Dispatch was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Command    *model.Command
		EventStore model.EventRepository
		Projection model.Projection
		Logger     model.Log
	}{
		Ctx:        ctx,
		Command:    command,
		EventStore: eventStore,
		Projection: projection,
		Logger:     logger,
	}
	mock.lockDispatch.Lock()
	mock.calls.Dispatch = append(mock.calls.Dispatch, callInfo)
	mock.lockDispatch.Unlock()
	return mock.DispatchFunc(ctx, command, eventStore, projection, logger)
}

// DispatchCalls gets all the calls that were made to Dispatch.
// Check the length with:
//     len(mockedCommandDispatcher.DispatchCalls())
func (mock *CommandDispatcherMock) DispatchCalls() []struct {
	Ctx        context.Context
	Command    *model.Command
	EventStore model.EventRepository
	Projection model.Projection
	Logger     model.Log
} {
	var calls []struct {
		Ctx        context.Context
		Command    *model.Command
		EventStore model.EventRepository
		Projection model.Projection
		Logger     model.Log
	}
	mock.lockDispatch.RLock()
	calls = mock.calls.Dispatch
	mock.lockDispatch.RUnlock()
	return calls
}

// GetSubscribers calls GetSubscribersFunc.
func (mock *CommandDispatcherMock) GetSubscribers() map[string][]model.CommandHandler {
	if mock.GetSubscribersFunc == nil {
		panic("CommandDispatcherMock.GetSubscribersFunc: method is nil but CommandDispatcher.GetSubscribers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetSubscribers.Lock()
	mock.calls.GetSubscribers = append(mock.calls.GetSubscribers, callInfo)
	mock.lockGetSubscribers.Unlock()
	return mock.GetSubscribersFunc()
}

// GetSubscribersCalls gets all the calls that were made to GetSubscribers.
// Check the length with:
//     len(mockedCommandDispatcher.GetSubscribersCalls())
func (mock *CommandDispatcherMock) GetSubscribersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetSubscribers.RLock()
	calls = mock.calls.GetSubscribers
	mock.lockGetSubscribers.RUnlock()
	return calls
}

// Ensure, that ServiceMock does implement model.Service.
// If this is not the case, regenerate this file with moq.
var _ model.Service = &ServiceMock{}

// ServiceMock is a mock implementation of model.Service.
//
// 	func TestSomethingThatUsesService(t *testing.T) {
//
// 		// make and configure a mocked model.Service
// 		mockedService := &ServiceMock{
// 			AddProjectionFunc: func(projection model.Projection) error {
// 				panic("mock out the AddProjection method")
// 			},
// 			ConfigFunc: func() *model.ServiceConfig {
// 				panic("mock out the Config method")
// 			},
// 			DBFunc: func() *gorm.DB {
// 				panic("mock out the DB method")
// 			},
// 			DBConnectionFunc: func() *sql.DB {
// 				panic("mock out the DBConnection method")
// 			},
// 			DispatcherFunc: func() model.CommandDispatcher {
// 				panic("mock out the Dispatcher method")
// 			},
// 			EventRepositoryFunc: func() model.EventRepository {
// 				panic("mock out the EventRepository method")
// 			},
// 			HTTPClientFunc: func() *http.Client {
// 				panic("mock out the HTTPClient method")
// 			},
// 			IDFunc: func() string {
// 				panic("mock out the ID method")
// 			},
// 			LoggerFunc: func() model.Log {
// 				panic("mock out the Logger method")
// 			},
// 			MigrateFunc: func(ctx context.Context, builders map[string]ds.Builder) error {
// 				panic("mock out the Migrate method")
// 			},
// 			ProjectionsFunc: func() []model.Projection {
// 				panic("mock out the Projections method")
// 			},
// 			TitleFunc: func() string {
// 				panic("mock out the Title method")
// 			},
// 		}
//
// 		// use mockedService in code that requires model.Service
// 		// and then make assertions.
//
// 	}
type ServiceMock struct {
	// AddProjectionFunc mocks the AddProjection method.
	AddProjectionFunc func(projection model.Projection) error

	// ConfigFunc mocks the Config method.
	ConfigFunc func() *model.ServiceConfig

	// DBFunc mocks the DB method.
	DBFunc func() *gorm.DB

	// DBConnectionFunc mocks the DBConnection method.
	DBConnectionFunc func() *sql.DB

	// DispatcherFunc mocks the Dispatcher method.
	DispatcherFunc func() model.CommandDispatcher

	// EventRepositoryFunc mocks the EventRepository method.
	EventRepositoryFunc func() model.EventRepository

	// HTTPClientFunc mocks the HTTPClient method.
	HTTPClientFunc func() *http.Client

	// IDFunc mocks the ID method.
	IDFunc func() string

	// LoggerFunc mocks the Logger method.
	LoggerFunc func() model.Log

	// MigrateFunc mocks the Migrate method.
	MigrateFunc func(ctx context.Context, builders map[string]ds.Builder) error

	// ProjectionsFunc mocks the Projections method.
	ProjectionsFunc func() []model.Projection

	// TitleFunc mocks the Title method.
	TitleFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// AddProjection holds details about calls to the AddProjection method.
		AddProjection []struct {
			// Projection is the projection argument value.
			Projection model.Projection
		}
		// Config holds details about calls to the Config method.
		Config []struct {
		}
		// DB holds details about calls to the DB method.
		DB []struct {
		}
		// DBConnection holds details about calls to the DBConnection method.
		DBConnection []struct {
		}
		// Dispatcher holds details about calls to the Dispatcher method.
		Dispatcher []struct {
		}
		// EventRepository holds details about calls to the EventRepository method.
		EventRepository []struct {
		}
		// HTTPClient holds details about calls to the HTTPClient method.
		HTTPClient []struct {
		}
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// Logger holds details about calls to the Logger method.
		Logger []struct {
		}
		// Migrate holds details about calls to the Migrate method.
		Migrate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Builders is the builders argument value.
			Builders map[string]ds.Builder
		}
		// Projections holds details about calls to the Projections method.
		Projections []struct {
		}
		// Title holds details about calls to the Title method.
		Title []struct {
		}
	}
	lockAddProjection   sync.RWMutex
	lockConfig          sync.RWMutex
	lockDB              sync.RWMutex
	lockDBConnection    sync.RWMutex
	lockDispatcher      sync.RWMutex
	lockEventRepository sync.RWMutex
	lockHTTPClient      sync.RWMutex
	lockID              sync.RWMutex
	lockLogger          sync.RWMutex
	lockMigrate         sync.RWMutex
	lockProjections     sync.RWMutex
	lockTitle           sync.RWMutex
}

// AddProjection calls AddProjectionFunc.
func (mock *ServiceMock) AddProjection(projection model.Projection) error {
	if mock.AddProjectionFunc == nil {
		panic("ServiceMock.AddProjectionFunc: method is nil but Service.AddProjection was just called")
	}
	callInfo := struct {
		Projection model.Projection
	}{
		Projection: projection,
	}
	mock.lockAddProjection.Lock()
	mock.calls.AddProjection = append(mock.calls.AddProjection, callInfo)
	mock.lockAddProjection.Unlock()
	return mock.AddProjectionFunc(projection)
}

// AddProjectionCalls gets all the calls that were made to AddProjection.
// Check the length with:
//     len(mockedService.AddProjectionCalls())
func (mock *ServiceMock) AddProjectionCalls() []struct {
	Projection model.Projection
} {
	var calls []struct {
		Projection model.Projection
	}
	mock.lockAddProjection.RLock()
	calls = mock.calls.AddProjection
	mock.lockAddProjection.RUnlock()
	return calls
}

// Config calls ConfigFunc.
func (mock *ServiceMock) Config() *model.ServiceConfig {
	if mock.ConfigFunc == nil {
		panic("ServiceMock.ConfigFunc: method is nil but Service.Config was just called")
	}
	callInfo := struct {
	}{}
	mock.lockConfig.Lock()
	mock.calls.Config = append(mock.calls.Config, callInfo)
	mock.lockConfig.Unlock()
	return mock.ConfigFunc()
}

// ConfigCalls gets all the calls that were made to Config.
// Check the length with:
//     len(mockedService.ConfigCalls())
func (mock *ServiceMock) ConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockConfig.RLock()
	calls = mock.calls.Config
	mock.lockConfig.RUnlock()
	return calls
}

// DB calls DBFunc.
func (mock *ServiceMock) DB() *gorm.DB {
	if mock.DBFunc == nil {
		panic("ServiceMock.DBFunc: method is nil but Service.DB was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDB.Lock()
	mock.calls.DB = append(mock.calls.DB, callInfo)
	mock.lockDB.Unlock()
	return mock.DBFunc()
}

// DBCalls gets all the calls that were made to DB.
// Check the length with:
//     len(mockedService.DBCalls())
func (mock *ServiceMock) DBCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDB.RLock()
	calls = mock.calls.DB
	mock.lockDB.RUnlock()
	return calls
}

// DBConnection calls DBConnectionFunc.
func (mock *ServiceMock) DBConnection() *sql.DB {
	if mock.DBConnectionFunc == nil {
		panic("ServiceMock.DBConnectionFunc: method is nil but Service.DBConnection was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDBConnection.Lock()
	mock.calls.DBConnection = append(mock.calls.DBConnection, callInfo)
	mock.lockDBConnection.Unlock()
	return mock.DBConnectionFunc()
}

// DBConnectionCalls gets all the calls that were made to DBConnection.
// Check the length with:
//     len(mockedService.DBConnectionCalls())
func (mock *ServiceMock) DBConnectionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDBConnection.RLock()
	calls = mock.calls.DBConnection
	mock.lockDBConnection.RUnlock()
	return calls
}

// Dispatcher calls DispatcherFunc.
func (mock *ServiceMock) Dispatcher() model.CommandDispatcher {
	if mock.DispatcherFunc == nil {
		panic("ServiceMock.DispatcherFunc: method is nil but Service.Dispatcher was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDispatcher.Lock()
	mock.calls.Dispatcher = append(mock.calls.Dispatcher, callInfo)
	mock.lockDispatcher.Unlock()
	return mock.DispatcherFunc()
}

// DispatcherCalls gets all the calls that were made to Dispatcher.
// Check the length with:
//     len(mockedService.DispatcherCalls())
func (mock *ServiceMock) DispatcherCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDispatcher.RLock()
	calls = mock.calls.Dispatcher
	mock.lockDispatcher.RUnlock()
	return calls
}

// EventRepository calls EventRepositoryFunc.
func (mock *ServiceMock) EventRepository() model.EventRepository {
	if mock.EventRepositoryFunc == nil {
		panic("ServiceMock.EventRepositoryFunc: method is nil but Service.EventRepository was just called")
	}
	callInfo := struct {
	}{}
	mock.lockEventRepository.Lock()
	mock.calls.EventRepository = append(mock.calls.EventRepository, callInfo)
	mock.lockEventRepository.Unlock()
	return mock.EventRepositoryFunc()
}

// EventRepositoryCalls gets all the calls that were made to EventRepository.
// Check the length with:
//     len(mockedService.EventRepositoryCalls())
func (mock *ServiceMock) EventRepositoryCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEventRepository.RLock()
	calls = mock.calls.EventRepository
	mock.lockEventRepository.RUnlock()
	return calls
}

// HTTPClient calls HTTPClientFunc.
func (mock *ServiceMock) HTTPClient() *http.Client {
	if mock.HTTPClientFunc == nil {
		panic("ServiceMock.HTTPClientFunc: method is nil but Service.HTTPClient was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHTTPClient.Lock()
	mock.calls.HTTPClient = append(mock.calls.HTTPClient, callInfo)
	mock.lockHTTPClient.Unlock()
	return mock.HTTPClientFunc()
}

// HTTPClientCalls gets all the calls that were made to HTTPClient.
// Check the length with:
//     len(mockedService.HTTPClientCalls())
func (mock *ServiceMock) HTTPClientCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHTTPClient.RLock()
	calls = mock.calls.HTTPClient
	mock.lockHTTPClient.RUnlock()
	return calls
}

// ID calls IDFunc.
func (mock *ServiceMock) ID() string {
	if mock.IDFunc == nil {
		panic("ServiceMock.IDFunc: method is nil but Service.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//     len(mockedService.IDCalls())
func (mock *ServiceMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// Logger calls LoggerFunc.
func (mock *ServiceMock) Logger() model.Log {
	if mock.LoggerFunc == nil {
		panic("ServiceMock.LoggerFunc: method is nil but Service.Logger was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLogger.Lock()
	mock.calls.Logger = append(mock.calls.Logger, callInfo)
	mock.lockLogger.Unlock()
	return mock.LoggerFunc()
}

// LoggerCalls gets all the calls that were made to Logger.
// Check the length with:
//     len(mockedService.LoggerCalls())
func (mock *ServiceMock) LoggerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLogger.RLock()
	calls = mock.calls.Logger
	mock.lockLogger.RUnlock()
	return calls
}

// Migrate calls MigrateFunc.
func (mock *ServiceMock) Migrate(ctx context.Context, builders map[string]ds.Builder) error {
	if mock.MigrateFunc == nil {
		panic("ServiceMock.MigrateFunc: method is nil but Service.Migrate was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Builders map[string]ds.Builder
	}{
		Ctx:      ctx,
		Builders: builders,
	}
	mock.lockMigrate.Lock()
	mock.calls.Migrate = append(mock.calls.Migrate, callInfo)
	mock.lockMigrate.Unlock()
	return mock.MigrateFunc(ctx, builders)
}

// MigrateCalls gets all the calls that were made to Migrate.
// Check the length with:
//     len(mockedService.MigrateCalls())
func (mock *ServiceMock) MigrateCalls() []struct {
	Ctx      context.Context
	Builders map[string]ds.Builder
} {
	var calls []struct {
		Ctx      context.Context
		Builders map[string]ds.Builder
	}
	mock.lockMigrate.RLock()
	calls = mock.calls.Migrate
	mock.lockMigrate.RUnlock()
	return calls
}

// Projections calls ProjectionsFunc.
func (mock *ServiceMock) Projections() []model.Projection {
	if mock.ProjectionsFunc == nil {
		panic("ServiceMock.ProjectionsFunc: method is nil but Service.Projections was just called")
	}
	callInfo := struct {
	}{}
	mock.lockProjections.Lock()
	mock.calls.Projections = append(mock.calls.Projections, callInfo)
	mock.lockProjections.Unlock()
	return mock.ProjectionsFunc()
}

// ProjectionsCalls gets all the calls that were made to Projections.
// Check the length with:
//     len(mockedService.ProjectionsCalls())
func (mock *ServiceMock) ProjectionsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockProjections.RLock()
	calls = mock.calls.Projections
	mock.lockProjections.RUnlock()
	return calls
}

// Title calls TitleFunc.
func (mock *ServiceMock) Title() string {
	if mock.TitleFunc == nil {
		panic("ServiceMock.TitleFunc: method is nil but Service.Title was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTitle.Lock()
	mock.calls.Title = append(mock.calls.Title, callInfo)
	mock.lockTitle.Unlock()
	return mock.TitleFunc()
}

// TitleCalls gets all the calls that were made to Title.
// Check the length with:
//     len(mockedService.TitleCalls())
func (mock *ServiceMock) TitleCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTitle.RLock()
	calls = mock.calls.Title
	mock.lockTitle.RUnlock()
	return calls
}

// Ensure, that EntityFactoryMock does implement model.EntityFactory.
// If this is not the case, regenerate this file with moq.
var _ model.EntityFactory = &EntityFactoryMock{}

// EntityFactoryMock is a mock implementation of model.EntityFactory.
//
// 	func TestSomethingThatUsesEntityFactory(t *testing.T) {
//
// 		// make and configure a mocked model.EntityFactory
// 		mockedEntityFactory := &EntityFactoryMock{
// 			BuilderFunc: func(ctx context.Context) ds.Builder {
// 				panic("mock out the Builder method")
// 			},
// 			CreateEntityWithValuesFunc: func(ctx context.Context, payload []byte) (*model.ContentEntity, error) {
// 				panic("mock out the CreateEntityWithValues method")
// 			},
// 			DynamicStructFunc: func(ctx context.Context) ds.DynamicStruct {
// 				panic("mock out the DynamicStruct method")
// 			},
// 			FromSchemaAndBuilderFunc: func(s string, schema *openapi3.Schema, builder ds.Builder) model.EntityFactory {
// 				panic("mock out the FromSchemaAndBuilder method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			NewEntityFunc: func(ctx context.Context) (*model.ContentEntity, error) {
// 				panic("mock out the NewEntity method")
// 			},
// 			SchemaFunc: func() *openapi3.Schema {
// 				panic("mock out the Schema method")
// 			},
// 			TableNameFunc: func() string {
// 				panic("mock out the TableName method")
// 			},
// 		}
//
// 		// use mockedEntityFactory in code that requires model.EntityFactory
// 		// and then make assertions.
//
// 	}
type EntityFactoryMock struct {
	// BuilderFunc mocks the Builder method.
	BuilderFunc func(ctx context.Context) ds.Builder

	// CreateEntityWithValuesFunc mocks the CreateEntityWithValues method.
	CreateEntityWithValuesFunc func(ctx context.Context, payload []byte) (*model.ContentEntity, error)

	// DynamicStructFunc mocks the DynamicStruct method.
	DynamicStructFunc func(ctx context.Context) ds.DynamicStruct

	// FromSchemaAndBuilderFunc mocks the FromSchemaAndBuilder method.
	FromSchemaAndBuilderFunc func(s string, schema *openapi3.Schema, builder ds.Builder) model.EntityFactory

	// NameFunc mocks the Name method.
	NameFunc func() string

	// NewEntityFunc mocks the NewEntity method.
	NewEntityFunc func(ctx context.Context) (*model.ContentEntity, error)

	// SchemaFunc mocks the Schema method.
	SchemaFunc func() *openapi3.Schema

	// TableNameFunc mocks the TableName method.
	TableNameFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Builder holds details about calls to the Builder method.
		Builder []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CreateEntityWithValues holds details about calls to the CreateEntityWithValues method.
		CreateEntityWithValues []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Payload is the payload argument value.
			Payload []byte
		}
		// DynamicStruct holds details about calls to the DynamicStruct method.
		DynamicStruct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FromSchemaAndBuilder holds details about calls to the FromSchemaAndBuilder method.
		FromSchemaAndBuilder []struct {
			// S is the s argument value.
			S string
			// Schema is the schema argument value.
			Schema *openapi3.Schema
			// Builder is the builder argument value.
			Builder ds.Builder
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// NewEntity holds details about calls to the NewEntity method.
		NewEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Schema holds details about calls to the Schema method.
		Schema []struct {
		}
		// TableName holds details about calls to the TableName method.
		TableName []struct {
		}
	}
	lockBuilder                sync.RWMutex
	lockCreateEntityWithValues sync.RWMutex
	lockDynamicStruct          sync.RWMutex
	lockFromSchemaAndBuilder   sync.RWMutex
	lockName                   sync.RWMutex
	lockNewEntity              sync.RWMutex
	lockSchema                 sync.RWMutex
	lockTableName              sync.RWMutex
}

// Builder calls BuilderFunc.
func (mock *EntityFactoryMock) Builder(ctx context.Context) ds.Builder {
	if mock.BuilderFunc == nil {
		panic("EntityFactoryMock.BuilderFunc: method is nil but EntityFactory.Builder was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockBuilder.Lock()
	mock.calls.Builder = append(mock.calls.Builder, callInfo)
	mock.lockBuilder.Unlock()
	return mock.BuilderFunc(ctx)
}

// BuilderCalls gets all the calls that were made to Builder.
// Check the length with:
//     len(mockedEntityFactory.BuilderCalls())
func (mock *EntityFactoryMock) BuilderCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockBuilder.RLock()
	calls = mock.calls.Builder
	mock.lockBuilder.RUnlock()
	return calls
}

// CreateEntityWithValues calls CreateEntityWithValuesFunc.
func (mock *EntityFactoryMock) CreateEntityWithValues(ctx context.Context, payload []byte) (*model.ContentEntity, error) {
	if mock.CreateEntityWithValuesFunc == nil {
		panic("EntityFactoryMock.CreateEntityWithValuesFunc: method is nil but EntityFactory.CreateEntityWithValues was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Payload []byte
	}{
		Ctx:     ctx,
		Payload: payload,
	}
	mock.lockCreateEntityWithValues.Lock()
	mock.calls.CreateEntityWithValues = append(mock.calls.CreateEntityWithValues, callInfo)
	mock.lockCreateEntityWithValues.Unlock()
	return mock.CreateEntityWithValuesFunc(ctx, payload)
}

// CreateEntityWithValuesCalls gets all the calls that were made to CreateEntityWithValues.
// Check the length with:
//     len(mockedEntityFactory.CreateEntityWithValuesCalls())
func (mock *EntityFactoryMock) CreateEntityWithValuesCalls() []struct {
	Ctx     context.Context
	Payload []byte
} {
	var calls []struct {
		Ctx     context.Context
		Payload []byte
	}
	mock.lockCreateEntityWithValues.RLock()
	calls = mock.calls.CreateEntityWithValues
	mock.lockCreateEntityWithValues.RUnlock()
	return calls
}

// DynamicStruct calls DynamicStructFunc.
func (mock *EntityFactoryMock) DynamicStruct(ctx context.Context) ds.DynamicStruct {
	if mock.DynamicStructFunc == nil {
		panic("EntityFactoryMock.DynamicStructFunc: method is nil but EntityFactory.DynamicStruct was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockDynamicStruct.Lock()
	mock.calls.DynamicStruct = append(mock.calls.DynamicStruct, callInfo)
	mock.lockDynamicStruct.Unlock()
	return mock.DynamicStructFunc(ctx)
}

// DynamicStructCalls gets all the calls that were made to DynamicStruct.
// Check the length with:
//     len(mockedEntityFactory.DynamicStructCalls())
func (mock *EntityFactoryMock) DynamicStructCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockDynamicStruct.RLock()
	calls = mock.calls.DynamicStruct
	mock.lockDynamicStruct.RUnlock()
	return calls
}

// FromSchemaAndBuilder calls FromSchemaAndBuilderFunc.
func (mock *EntityFactoryMock) FromSchemaAndBuilder(s string, schema *openapi3.Schema, builder ds.Builder) model.EntityFactory {
	if mock.FromSchemaAndBuilderFunc == nil {
		panic("EntityFactoryMock.FromSchemaAndBuilderFunc: method is nil but EntityFactory.FromSchemaAndBuilder was just called")
	}
	callInfo := struct {
		S       string
		Schema  *openapi3.Schema
		Builder ds.Builder
	}{
		S:       s,
		Schema:  schema,
		Builder: builder,
	}
	mock.lockFromSchemaAndBuilder.Lock()
	mock.calls.FromSchemaAndBuilder = append(mock.calls.FromSchemaAndBuilder, callInfo)
	mock.lockFromSchemaAndBuilder.Unlock()
	return mock.FromSchemaAndBuilderFunc(s, schema, builder)
}

// FromSchemaAndBuilderCalls gets all the calls that were made to FromSchemaAndBuilder.
// Check the length with:
//     len(mockedEntityFactory.FromSchemaAndBuilderCalls())
func (mock *EntityFactoryMock) FromSchemaAndBuilderCalls() []struct {
	S       string
	Schema  *openapi3.Schema
	Builder ds.Builder
} {
	var calls []struct {
		S       string
		Schema  *openapi3.Schema
		Builder ds.Builder
	}
	mock.lockFromSchemaAndBuilder.RLock()
	calls = mock.calls.FromSchemaAndBuilder
	mock.lockFromSchemaAndBuilder.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *EntityFactoryMock) Name() string {
	if mock.NameFunc == nil {
		panic("EntityFactoryMock.NameFunc: method is nil but EntityFactory.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedEntityFactory.NameCalls())
func (mock *EntityFactoryMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// NewEntity calls NewEntityFunc.
func (mock *EntityFactoryMock) NewEntity(ctx context.Context) (*model.ContentEntity, error) {
	if mock.NewEntityFunc == nil {
		panic("EntityFactoryMock.NewEntityFunc: method is nil but EntityFactory.NewEntity was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockNewEntity.Lock()
	mock.calls.NewEntity = append(mock.calls.NewEntity, callInfo)
	mock.lockNewEntity.Unlock()
	return mock.NewEntityFunc(ctx)
}

// NewEntityCalls gets all the calls that were made to NewEntity.
// Check the length with:
//     len(mockedEntityFactory.NewEntityCalls())
func (mock *EntityFactoryMock) NewEntityCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockNewEntity.RLock()
	calls = mock.calls.NewEntity
	mock.lockNewEntity.RUnlock()
	return calls
}

// Schema calls SchemaFunc.
func (mock *EntityFactoryMock) Schema() *openapi3.Schema {
	if mock.SchemaFunc == nil {
		panic("EntityFactoryMock.SchemaFunc: method is nil but EntityFactory.Schema was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSchema.Lock()
	mock.calls.Schema = append(mock.calls.Schema, callInfo)
	mock.lockSchema.Unlock()
	return mock.SchemaFunc()
}

// SchemaCalls gets all the calls that were made to Schema.
// Check the length with:
//     len(mockedEntityFactory.SchemaCalls())
func (mock *EntityFactoryMock) SchemaCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSchema.RLock()
	calls = mock.calls.Schema
	mock.lockSchema.RUnlock()
	return calls
}

// TableName calls TableNameFunc.
func (mock *EntityFactoryMock) TableName() string {
	if mock.TableNameFunc == nil {
		panic("EntityFactoryMock.TableNameFunc: method is nil but EntityFactory.TableName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTableName.Lock()
	mock.calls.TableName = append(mock.calls.TableName, callInfo)
	mock.lockTableName.Unlock()
	return mock.TableNameFunc()
}

// TableNameCalls gets all the calls that were made to TableName.
// Check the length with:
//     len(mockedEntityFactory.TableNameCalls())
func (mock *EntityFactoryMock) TableNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTableName.RLock()
	calls = mock.calls.TableName
	mock.lockTableName.RUnlock()
	return calls
}
