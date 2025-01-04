package models

type Category struct {
    IDCategory uint   `json:"idCategory" gorm:"primaryKey;autoIncrement"`
    Name       string `json:"name"`
}
