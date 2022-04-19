package storage

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name string
}
