package interfaces

type RoomController interface {
	Index(Context) error
	NewRoom(Context) error
	EnterRoom(Context) error
}

type UserController interface {
	LoginMenu(Context) error
	Logout(Context) error
}
