package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type user struct {
	object
	github string
}

// noinspection GoUnusedExportedFunction
func InitializeUser() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewUser(id int64, github string) (IUser, error) {
	instance := &user{
		object: object{
			id: id,
		},
		github: github,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewUserFromEntity(entity IUserEntity) (IUser, error) {
	instance := &user{
		object: object{
			id: entity.Id(),
		},
		github: entity.Github(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (user *user) Github() string {
	return user.github
}

func (user *user) UpdateGithub(github string, editor Identity) {
	if err := repository.Users.UpdateGithub(user.id, github, editor.Id()); err != nil {
		panic(err.Error())
	}

	user.github = github
}

func (user *user) UpdateGithubAtomic(transaction ITransaction, github string, editor Identity) {
	transaction.OnCommit(func() {
		user.github = github
	})

	if err := repository.Users.UpdateGithubAtomic(transaction, user.id, github, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (user *user) Validate() error {
	return nil
}

func (user *user) String() string {
	return fmt.Sprintf("User (Id: %d, Github: %v)", user.Id(), user.Github())
}

//------------------------------------------------------------------------------

type users struct {
	collection Users
}

// NewUsers creates an empty collection of 'User' which is not thread-safe.
func NewUsers() IUserCollection {
	return &users{
		collection: make(Users, 0),
	}
}

func (users *users) Count() int {
	return len(users.collection)
}

func (users *users) IsEmpty() bool {
	return len(users.collection) == 0
}

func (users *users) IsNotEmpty() bool {
	return len(users.collection) > 0
}

func (users *users) HasExactlyOneItem() bool {
	return len(users.collection) == 1
}

func (users *users) HasAtLeastOneItem() bool {
	return len(users.collection) >= 1
}

func (users *users) First() IUser {
	return users.collection[0]
}

func (users *users) Append(user IUser) {
	users.collection = append(users.collection, user)
}

func (users *users) Reverse() IUserCollection {
	slice := users.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	users.collection = slice

	return users
}

func (users *users) ForEach(iterator UserIterator) {
	if iterator == nil {
		return
	}

	for _, value := range users.collection {
		iterator(value)
	}
}

func (users *users) Array() Users {
	return users.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) UserExists(id int64) bool {
	return dispatcher.conductor.UserManager().Exists(id)
}

func (dispatcher *dispatcher) UserExistsWhich(condition UserCondition) bool {
	return dispatcher.conductor.UserManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListUsers() IUserCollection {
	return dispatcher.conductor.UserManager().ListUsers(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachUser(iterator UserIterator) {
	dispatcher.conductor.UserManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterUsers(predicate UserFilterPredicate) IUserCollection {
	return dispatcher.conductor.UserManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapUsers(predicate UserMapPredicate) IUserCollection {
	return dispatcher.conductor.UserManager().Map(predicate)
}

func (dispatcher *dispatcher) GetUser(id int64) IUser {
	if user, err := dispatcher.conductor.UserManager().GetUser(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return user
	}
}

func (dispatcher *dispatcher) AddUser(identityId int64, github string) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().AddUserAtomic(transaction, identityId, github, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().AddUser(identityId, github, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}

func (dispatcher *dispatcher) AddUserObject(identity IIdentity, user IUser) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().AddUserObjectAtomic(transaction, identity.Id(), user, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().AddUserObject(identity.Id(), user, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}

func (dispatcher *dispatcher) LogUser(identityId int64, github string, source string, payload string) {
	dispatcher.conductor.UserManager().Log(identityId, github, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateUser(id int64, github string) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().UpdateUserAtomic(transaction, id, github, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().UpdateUser(id, github, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateUserObject(object IObject, user IUser) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().UpdateUserAtomic(transaction, object.Id(), user.Github(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().UpdateUser(object.Id(), user.Github(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateUserObject(object IObject, user IUser) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().AddOrUpdateUserObjectAtomic(transaction, object.Id(), user, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().AddOrUpdateUserObject(object.Id(), user, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}

func (dispatcher *dispatcher) RemoveUser(id int64) IUser {
	transaction := dispatcher.transaction
	if transaction != nil {
		if user, err := dispatcher.conductor.UserManager().RemoveUserAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	} else {
		if user, err := dispatcher.conductor.UserManager().RemoveUser(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return user
		}
	}
}
