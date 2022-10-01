package repos

import (
	"contracts"
	"db"
	"errors"

	"github.com/gofiber/fiber/v2"
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
func CreateIncomingActivity(activity *IncomingActivity) error {
	if err := db.Executor.Create(activity).Error; err != nil {
		return err
	}

	return nil
}

// FindIncomingActivity searches the incoming activities table with the condition given
// and returns a single record.
func FindIncomingActivity(conds ...interface{}) (*IncomingActivity, error) {
	dest := &IncomingActivity{}
	if err := db.Executor.Model(&IncomingActivity{}).Take(dest, conds...).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &fiber.Error{
				Code:    contracts.StatusNotFound,
				Message: "activity not found",
			}
		} else {
			return nil, &fiber.Error{
				Code:    contracts.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return dest, nil
}

// FindIncomingActivitiesForUser finds the activities posted to user
func FindIncomingActivitiesForUser(userIden interface{}) ([]IncomingActivity, error) {
	result := &[]IncomingActivity{}
	if err := db.Executor.Model(&IncomingActivity{}).Find(result, "`to` = ?", userIden).Error; err != nil {
		return *result, err
	}

	return *result, nil
}

// FindIncomingActivityById searches the incoming activities table with the id given
func FindIncomingActivityById(id uint) (*IncomingActivity, error) {
	return FindIncomingActivity("id = ?", id)
}

// FindIncomingActivityByGuid searches the incoming activities table with the guid given
func FindIncomingActivityByGuid(guid string) (*IncomingActivity, error) {
	return FindIncomingActivity("guid = ?", guid)
}
