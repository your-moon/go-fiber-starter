package models

type User struct {
	ID    uint   `gorm:"primaryKey"   json:"id"`
	Email string `gorm:"column:email" json:"email" fake:"email"`
}
