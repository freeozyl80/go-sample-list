package test

import (
	"testing"
  "pattern/factory/event"
)

func TestEventFactory(t *testing.T) {
    factory := event.Factory{}
    e := factory.Create(event.Start)
    if e.EventType() != event.Start {
        t.Errorf("expect event.Start, but actual %v.", e.EventType())
    }
    e = factory.Create(event.End)
    if e.EventType() != event.End {
        t.Errorf("expect event.End, but actual %v.", e.EventType())
    }
}