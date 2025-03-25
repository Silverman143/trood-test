package eventdispatcher

type UnresolvedIntentEvent struct {
	ChatID int64 
}

func (e *UnresolvedIntentEvent) GetNamespace() string {
	return "unresolved_intent"
}

func (e *UnresolvedIntentEvent) GetAggregateID() int64 {
	return e.ChatID
}