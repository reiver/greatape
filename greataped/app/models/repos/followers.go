package repos

import (
	"db"

	"gorm.io/gorm"
)

// Follower struct defines a follower
type Follower struct {
	gorm.Model
	Target   string `gorm:"not null"`
	Handle   string `gorm:"not null"`
	Accepted bool   `gorm:"not null"`
}

// CreateFollower creates a new entry in the followers's table
func CreateFollower(follower *Follower) *gorm.DB {
	return db.DB.Create(follower)
}

// FindFollowers finds the user's followers
func FindFollowers(dest interface{}, userIden interface{}) *gorm.DB {
	return db.DB.Model(&Follower{}).Find(dest, "`target` = ?", userIden)
}

// AcceptFollower accepts a follow request
func AcceptFollower(id interface{}) *gorm.DB {
	return db.DB.Model(&Follower{}).Where("id = ?", id).Update("accepted", true)
}
