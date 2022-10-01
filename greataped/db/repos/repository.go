package repos

import "gorm.io/gorm"

type repository struct {
	Storage *gorm.DB
}
