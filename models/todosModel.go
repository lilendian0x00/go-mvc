package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title 			string	`json:"title"`
	Description 	string	`json:"description"`
	Done			bool	`json:"done"`
}