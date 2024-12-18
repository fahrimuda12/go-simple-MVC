package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}