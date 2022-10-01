package repos

import (
	"gorm.io/gorm"
)

// Following struct defines an account that the user follows
type Following struct {
	gorm.Model
	Target string `gorm:"not null"`
	Handle string `gorm:"not null"`
}

// CreateFollowing creates a new entry in the following's table
func (repo *repository) CreateFollowing(following *Following) error {
	if err := repo.Storage.Create(following).Error; err != nil {
		return err
	}

	return nil
}

// FindFollowing finds what accounts the user is following
func (repo *repository) FindFollowing(userIden interface{}) ([]Following, error) {
	followings := &[]Following{}
	if err := repo.Storage.Model(&Following{}).Find(followings, "`target` = ?", userIden).Error; err != nil {
		return *followings, err
	}

	return *followings, nil
}
