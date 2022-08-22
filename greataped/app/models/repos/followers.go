package repos

import (
	"db"

	"gorm.io/gorm"
)

// Follower struct defines a follower
type Follower struct {
	gorm.Model
	Target      string `gorm:"not null"`
	Handle      string `gorm:"not null"`
	HandleInbox string
	Activity    string
	Accepted    bool
}

// CreateFollower creates a new entry in the followers's table
func CreateFollower(follower *Follower) *gorm.DB {
	return db.Executor.Create(follower)
}

// FindFollowers finds the user's followers
func FindFollowers(dest interface{}, userIden interface{}) *gorm.DB {
	return db.Executor.Model(&Follower{}).Find(dest, "`target` = ?", userIden)
}

// FindFollower searches the follower's table with the condition given
func FindFollower(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Executor.Model(&Follower{}).Take(dest, conds...)
}

// FindFollowerById searches the followers's table with the id given
func FindFollowerById(dest interface{}, id uint64) *gorm.DB {
	return FindFollower(dest, "id = ?", id)
}

// AcceptFollower accepts a follow request
func AcceptFollower(id interface{}) *gorm.DB {
	return db.Executor.Model(&Follower{}).Where("id = ?", id).Update("accepted", true)
}
