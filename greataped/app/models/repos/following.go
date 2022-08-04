package repos

import (
	"db"

	"gorm.io/gorm"
)

// Following struct defines a follower
type Following struct {
	gorm.Model
	Handle string `gorm:"not null"`
}

// CreateFollowing creates a new entry in the followers's table
func CreateFollowing(following *Following) *gorm.DB {
	return db.DB.Create(following)
}
