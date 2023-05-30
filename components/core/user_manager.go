package core

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
)

// noinspection GoSnakeCaseUsage
const USER_MANAGER = "UserManager"

type userManager struct {
	systemComponent
	cache ICache
}

func newUserManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IUserManager {
	manager := &userManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *userManager) Name() string {
	return USER_MANAGER
}

func (manager *userManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *userManager) Load() error {
	userEntities, err := repository.Users.FetchAll()
	if err != nil {
		return err
	}

	users := make(SystemObjectCache)
	for _, userEntity := range userEntities {
		if user, err := NewUserFromEntity(userEntity); err == nil {
			users[user.Id()] = user
		} else {
			return err
		}
	}

	manager.cache.Load(users)
	return nil
}

func (manager *userManager) Reload() error {
	return manager.Load()
}

func (manager *userManager) OnCacheChanged(callback UserCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *userManager) Count() int {
	return manager.cache.Size()
}

func (manager *userManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *userManager) ExistsWhich(condition UserCondition) bool {
	var users Users
	manager.ForEach(func(user IUser) {
		if condition(user) {
			users = append(users, user)
		}
	})

	return len(users) > 0
}

func (manager *userManager) ListUsers(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IUserCollection {
	return manager.Filter(UserPassThroughFilter)
}

func (manager *userManager) GetUser(id int64, _ Identity) (IUser, error) {
	if user := manager.Find(id); user == nil {
		return nil, ERROR_USER_NOT_FOUND
	} else {
		return user, nil
	}
}

func (manager *userManager) AddUser(identityId int64, github string, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(identityId, github)
	return manager.Apply(userEntity, repository.Users.Add, manager.cache.Put, editor)
}

func (manager *userManager) AddUserObject(identityId int64, user IUser, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(identityId, user.Github())
	return manager.Apply(userEntity, repository.Users.Add, manager.cache.Put, editor)
}

func (manager *userManager) AddUserAtomic(transaction ITransaction, identityId int64, github string, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(identityId, github)
	return manager.ApplyAtomic(transaction, userEntity, repository.Users.AddAtomic, manager.cache.Put, editor)
}

func (manager *userManager) AddUserObjectAtomic(transaction ITransaction, identityId int64, user IUser, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(identityId, user.Github())
	return manager.ApplyAtomic(transaction, userEntity, repository.Users.AddAtomic, manager.cache.Put, editor)
}

func (manager *userManager) Log(identityId int64, github string, source string, editor Identity, payload string) {
	userPipeEntity := NewUserPipeEntity(identityId, github, source, editor.Id(), payload)
	repository.Pipe.Insert(userPipeEntity)

	user, err := NewUserFromEntity(userPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(user.Id(), user)
	}
}

func (manager *userManager) UpdateUser(id int64, github string, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, github)
	return manager.Apply(userEntity, repository.Users.Update, manager.cache.Put, editor)
}

func (manager *userManager) UpdateUserObject(id int64, user IUser, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, user.Github())
	return manager.Apply(userEntity, repository.Users.Update, manager.cache.Put, editor)
}

func (manager *userManager) UpdateUserAtomic(transaction ITransaction, id int64, github string, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, github)
	return manager.ApplyAtomic(transaction, userEntity, repository.Users.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *userManager) UpdateUserObjectAtomic(transaction ITransaction, id int64, user IUser, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, user.Github())
	return manager.ApplyAtomic(transaction, userEntity, repository.Users.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *userManager) AddOrUpdateUserObject(id int64, user IUser, editor Identity) (IUser, error) {
	if manager.Exists(id) {
		return manager.UpdateUserObject(id, user, editor)
	} else {
		return manager.AddUserObject(id, user, editor)
	}
}

func (manager *userManager) AddOrUpdateUserObjectAtomic(transaction ITransaction, id int64, user IUser, editor Identity) (IUser, error) {
	if manager.Exists(id) {
		return manager.UpdateUserObjectAtomic(transaction, id, user, editor)
	} else {
		return manager.AddUserObjectAtomic(transaction, id, user, editor)
	}
}

func (manager *userManager) RemoveUser(id int64, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, "")
	return manager.Apply(userEntity, repository.Users.Remove, manager.cache.Remove, editor)
}

func (manager *userManager) RemoveUserAtomic(transaction ITransaction, id int64, editor Identity) (IUser, error) {
	userEntity := NewUserEntity(id, "")
	return manager.ApplyAtomic(transaction, userEntity, repository.Users.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *userManager) Apply(userEntity IUserEntity, repositoryHandler func(IUserEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IUser, error) {
	result, err := NewUserFromEntity(userEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(userEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *userManager) ApplyAtomic(transaction ITransaction, userEntity IUserEntity, repositoryHandler func(IRepositoryTransaction, IUserEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IUser, error) {
	result, err := NewUserFromEntity(userEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, userEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *userManager) Find(id int64) IUser {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IUser)
	}
}

func (manager *userManager) ForEach(iterator UserIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IUser))
	})
}

func (manager *userManager) Filter(predicate UserFilterPredicate) IUserCollection {
	users := NewUsers()
	if predicate == nil {
		return users
	}

	manager.ForEach(func(user IUser) {
		if predicate(user) {
			users.Append(user)
		}
	})

	return users
}

func (manager *userManager) Map(predicate UserMapPredicate) IUserCollection {
	users := NewUsers()
	if predicate == nil {
		return users
	}

	manager.ForEach(func(user IUser) {
		users.Append(predicate(user))
	})

	return users
}
