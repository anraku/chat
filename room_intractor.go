// package usecase
package main

type RoomInteractor struct {
	repository *RoomRepository
}

func NewRoomInteractor(r *RoomRepository) *RoomInteractor {
	return &RoomInteractor{
		repository: r,
	}
}

func (interactor *RoomInteractor) Fetch() (rooms []Room, err error) {
	rooms, err = interactor.repository.Fetch()
	return
}

func (interactor *RoomInteractor) Create(room Room) (err error) {
	err = interactor.repository.Create(room)
	return
}
