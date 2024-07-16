package models

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(300)" json:"firstName"`
	LastName  string `gorm:"type:varchar(300)" json:"lastName"`
	Email     string `gorm:"type:varchar(300)" json:"email"`
}
