package repos

import (
	"contracts"
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
func (repo *repository) CreateIncomingActivity(activity *IncomingActivity) error {
	if err := repo.Storage.Create(activity).Error; err != nil {
		return err
	}

	return nil
}

// FindIncomingActivity searches the incoming activities table with the condition given
// and returns a single record.
func (repo *repository) FindIncomingActivity(conds ...interface{}) (*IncomingActivity, error) {
	dest := &IncomingActivity{}
	if err := repo.Storage.Model(&IncomingActivity{}).Take(dest, conds...).Error; err != nil {
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
func (repo *repository) FindIncomingActivitiesForUser(userIden interface{}) ([]IncomingActivity, error) {
	result := &[]IncomingActivity{}
	if err := repo.Storage.Model(&IncomingActivity{}).Find(result, "`to` = ?", userIden).Error; err != nil {
		return *result, err
	}

	return *result, nil
}

// FindIncomingActivityById searches the incoming activities table with the id given
func (repo *repository) FindIncomingActivityById(id uint) (*IncomingActivity, error) {
	return repo.FindIncomingActivity("id = ?", id)
}

// FindIncomingActivityByGuid searches the incoming activities table with the guid given
func (repo *repository) FindIncomingActivityByGuid(guid string) (*IncomingActivity, error) {
	return repo.FindIncomingActivity("guid = ?", guid)
}
