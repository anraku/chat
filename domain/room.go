package domain

// Roomは一つのチャットルームを表します
type Room struct {
	ID          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}
