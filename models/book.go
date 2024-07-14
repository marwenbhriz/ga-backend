package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	authorName  string `gorm:"type:varchar(300)" json:"author_name"`
	title       string `gorm:"type:varchar(300)" json:"title"`
	description string `gorm:"type:varchar(300)" json:"description"`
}
