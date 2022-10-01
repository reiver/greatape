package repos

import (
	"contracts"
	"errors"

	"github.com/gofiber/fiber/v2"
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
func (repo *repository) CreateFollower(follower *Follower) error {
	if err := repo.Storage.Create(follower).Error; err != nil {
		return err
	}

	return nil
}

// FindFollower searches the followers table with the condition given
func (repo *repository) FindFollower(conds ...any) (*Follower, error) {
	dest := &Follower{}
	if err := repo.Storage.Model(&Follower{}).Take(dest, conds...).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &fiber.Error{
				Code:    contracts.StatusNotFound,
				Message: "follower not found",
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

// FindFollowers finds the user's followers
func (repo *repository) FindFollowers(userIden interface{}) ([]Follower, error) {
	result := &[]Follower{}
	if err := repo.Storage.Model(&Follower{}).Find(result, "`target` = ?", userIden).Error; err != nil {
		return *result, err
	}

	return *result, nil
}

// FindFollowerById searches the followers's table with the id given
func (repo *repository) FindFollowerById(id uint64) (*Follower, error) {
	return repo.FindFollower("id = ?", id)
}

// AcceptFollower accepts a follow request
func (repo *repository) AcceptFollower(id interface{}) error {
	if err := repo.Storage.Model(&Follower{}).Where("id = ?", id).Update("accepted", true).Error; err != nil {
		return err
	}

	return nil
}
