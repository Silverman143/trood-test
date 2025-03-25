package eventdispatcher

import (
	"context"
)

type Event interface{
	GetNamespace() string
	GetAggregateID() int64
}

type EventDispatcher interface{
	Dispatch(ctx context.Context, event Event) error
	Close() error
}