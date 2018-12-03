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
	if err != nil {
		return nil, err
	}
	return rooms, err
}
