package repos

import (
	"db"

	"gorm.io/gorm"
)

// OutgoingActivity struct defines the activity
type OutgoingActivity struct {
	gorm.Model
	Timestamp int64
	From      string `gorm:"not null"`
	To        string `gorm:"not null"`
	Guid      string `gorm:"uniqueIndex;not null"`
	Content   string `gorm:"not null"`
}

// CreateOutgoingActivity creates an activity entry in the outgoing activities table
func CreateOutgoingActivity(activity *OutgoingActivity) *gorm.DB {
	return db.Executor.Create(activity)
}

// FindOutgoingActivitiesByUser finds the activities posted by user
func FindOutgoingActivitiesByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return db.Executor.Model(&OutgoingActivity{}).Find(dest, "`from` = ?", userIden)
}

// FindOutgoingActivity searches the outgoing activities table with the condition given
// and returns a single record.
func FindOutgoingActivity(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Executor.Model(&OutgoingActivity{}).Take(dest, conds...)
}

// FindOutgoingActivityById searches the outgoing activities table with the id given
func FindOutgoingActivityById(dest interface{}, id uint) *gorm.DB {
	return FindOutgoingActivity(dest, "id = ?", id)
}

// FindOutgoingActivityByGuid searches the outgoing activities table with the guid given
func FindOutgoingActivityByGuid(dest interface{}, guid string) *gorm.DB {
	return FindOutgoingActivity(dest, "guid = ?", guid)
}
