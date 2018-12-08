// package usecase
package main

type RoomInteractor struct {
	roomRepository    *RoomRepository
	messageRepository *MessageRepository
}

func NewRoomInteractor(r *RoomRepository, m *MessageRepository) *RoomInteractor {
	return &RoomInteractor{
		roomRepository:    r,
		messageRepository: m,
	}
}

func (interactor *RoomInteractor) Fetch() (rooms []Room, err error) {
	rooms, err = interactor.roomRepository.Fetch()
	return
}

func (interactor *RoomInteractor) Create(room Room) (err error) {
	err = interactor.roomRepository.Create(room)
	return
}
