package repos

import (
	"contracts"
	"errors"

	"github.com/gofiber/fiber/v2"
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
func (repo *repository) CreateOutgoingActivity(activity *OutgoingActivity) error {
	if err := repo.Storage.Create(activity).Error; err != nil {
		return err
	}

	return nil
}

// FindOutgoingActivity searches the outgoing activities table with the condition given
// and returns a single record.
func (repo *repository) FindOutgoingActivity(conds ...interface{}) (*OutgoingActivity, error) {
	dest := &OutgoingActivity{}
	if err := repo.Storage.Model(&OutgoingActivity{}).Take(dest, conds...).Error; err != nil {
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

// FindOutgoingActivitiesByUser finds the activities posted by user
func (repo *repository) FindOutgoingActivitiesByUser(userIden interface{}) ([]OutgoingActivity, error) {
	result := &[]OutgoingActivity{}
	if err := repo.Storage.Model(&OutgoingActivity{}).Find(result, "`from` = ?", userIden).Error; err != nil {
		return *result, err
	}

	return *result, nil
}

// FindOutgoingActivityById searches the outgoing activities table with the id given
func (repo *repository) FindOutgoingActivityById(id uint) (*OutgoingActivity, error) {
	return repo.FindOutgoingActivity("id = ?", id)
}

// FindOutgoingActivityByGuid searches the outgoing activities table with the guid given
func (repo *repository) FindOutgoingActivityByGuid(guid string) (*OutgoingActivity, error) {
	return repo.FindOutgoingActivity("guid = ?", guid)
}
