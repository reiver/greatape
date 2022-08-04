package repos

import (
	"db"

	"gorm.io/gorm"
)

// Follower struct defines a follower
type Follower struct {
	gorm.Model
	Handle string `gorm:"not null"`
}

// CreateFollower creates a new entry in the followers's table
func CreateFollower(follower *Follower) *gorm.DB {
	return db.DB.Create(follower)
}
