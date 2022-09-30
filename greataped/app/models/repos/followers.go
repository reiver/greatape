package repos

import (
	"contracts"
	"db"
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
func CreateFollower(follower *Follower) error {
	if err := db.Executor.Create(follower).Error; err != nil {
		return err
	}

	return nil
}

// FindFollower searches the followers table with the condition given
func FindFollower(conds ...any) (*Follower, error) {
	dest := &Follower{}
	if err := db.Executor.Model(&Follower{}).Take(dest, conds...).Error; err != nil {
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
func FindFollowers(userIden interface{}) ([]Follower, error) {
	followers := &[]Follower{}
	if err := db.Executor.Model(&Follower{}).Find(followers, "`target` = ?", userIden).Error; err != nil {
		return *followers, err
	}

	return *followers, nil
}

// FindFollowerById searches the followers's table with the id given
func FindFollowerById(id uint64) (*Follower, error) {
	return FindFollower("id = ?", id)
}

// AcceptFollower accepts a follow request
func AcceptFollower(id interface{}) error {
	if err := db.Executor.Model(&Follower{}).Where("id = ?", id).Update("accepted", true).Error; err != nil {
		return err
	}

	return nil
}
