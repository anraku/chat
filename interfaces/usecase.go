package interfaces

import "github.com/anraku/chat/domain"

type MessageInteractor interface {
	StoreData(m *domain.Message) error
}
