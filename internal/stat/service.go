package stat

import (
	"log"
	"url-shortener/pkg/event"
)

type Service struct {
	EventBus *event.EventBus
	StatRepo *Repository
}

type ServiceDeps struct {
	EventBus *event.EventBus
	StatRepo *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		EventBus: deps.EventBus,
		StatRepo: deps.StatRepo,
	}
}

func (s *Service) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data: ", msg.Data)
				continue
			}
			s.StatRepo.AddClick(id)
		}
	}
}
