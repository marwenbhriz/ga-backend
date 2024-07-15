package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	AuthorName  string `gorm:"type:varchar(300)" json:"author_name"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Review      string `gorm:"type:text" json:"review"`
}
