package repos

import (
	"db"

	"gorm.io/gorm"
)

// Following struct defines an account that the user follows
type Following struct {
	gorm.Model
	Target string `gorm:"not null"`
	Handle string `gorm:"not null"`
}

// CreateFollowing creates a new entry in the following's table
func CreateFollowing(following *Following) *gorm.DB {
	return db.DB.Create(following)
}

// FindFollowing finds what accounts the user is following
func FindFollowing(dest interface{}, userIden interface{}) *gorm.DB {
	return db.DB.Model(&Follower{}).Find(dest, "`handle` = ?", userIden)
}
