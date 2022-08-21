package models

type Config struct {
	Name  string `json:"name" gorm:"primary_key"`
	Value string `json:"value"`
}
