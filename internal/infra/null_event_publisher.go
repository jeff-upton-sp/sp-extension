package infra

import (
	"context"
	"fmt"

	"github.com/sailpoint/atlas-go/atlas/application"
	"github.com/sailpoint/atlas-go/atlas/event"
)

type nullEventPublisher struct {
}

func (p *nullEventPublisher) BulkPublish(ctx context.Context, events []event.EventAndTopic) ([]*event.FailedEventAndTopic, error) {
	return nil, fmt.Errorf("not implemented")
}

func (p *nullEventPublisher) Publish(ctx context.Context, td event.TopicDescriptor, event *event.Event) error {
	return fmt.Errorf("not implemented")
}

func (p *nullEventPublisher) PublishToTopic(ctx context.Context, topic event.Topic, event *event.Event) error {
	return fmt.Errorf("not implemented")
}

func (p *nullEventPublisher) PublishWithDelay(ctx context.Context, topic event.Topic, event *event.Event, delaySeconds int) error {
	return fmt.Errorf("not implemented")
}

func WithNullEventPublisher() application.ConfigurationOption {
	return func(app *application.Application) error {
		app.EventPublisher = &nullEventPublisher{}
		return nil
	}
}
