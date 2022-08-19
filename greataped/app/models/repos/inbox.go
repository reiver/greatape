package repos

import (
	"db"

	"gorm.io/gorm"
)

// IncomingActivity struct defines the activity
type IncomingActivity struct {
	gorm.Model
	Timestamp int64
	From      string `gorm:"not null"`
	To        string `gorm:"not null"`
	Guid      string `gorm:"uniqueIndex;not null"`
	Content   string `gorm:"not null"`
}

// CreateIncomingActivity creates an activity entry in the incoming activities table
func CreateIncomingActivity(activity *IncomingActivity) *gorm.DB {
	return db.Executor.Create(activity)
}

// FindIncomingActivitiesForUser finds the activities posted to user
func FindIncomingActivitiesForUser(dest interface{}, userIden interface{}) *gorm.DB {
	return db.Executor.Model(&IncomingActivity{}).Find(dest, "`to` = ?", userIden)
}

// FindIncomingActivity searches the incoming activities table with the condition given
// and returns a single record.
func FindIncomingActivity(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Executor.Model(&IncomingActivity{}).Take(dest, conds...)
}

// FindIncomingActivityById searches the incoming activities table with the id given
func FindIncomingActivityById(dest interface{}, id uint) *gorm.DB {
	return FindIncomingActivity(dest, "id = ?", id)
}

// FindIncomingActivityByGuid searches the incoming activities table with the guid given
func FindIncomingActivityByGuid(dest interface{}, guid string) *gorm.DB {
	return FindIncomingActivity(dest, "guid = ?", guid)
}
