package eventmanager

import (
	grafeas "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
)

//go:generate mockgen -destination=../../mocks/pkg/eventmanager_mock/event_manager.go -package=eventmanager_mock . EventManager

type EventManager interface {
	Initialize(attesterName string) error
	Publish(attesterName string, occurrence *grafeas.Occurrence) error
	Subscribe(attesterName string) error
	Unsubscribe(attesterName string) error
}
