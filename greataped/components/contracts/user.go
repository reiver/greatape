package contracts

import . "github.com/xeronith/diamante/contracts/security"

var UserPassThroughFilter = func(IUser) bool { return true }

type (
	Users               []IUser
	UserIterator        func(IUser)
	UserCondition       func(IUser) bool
	UserFilterPredicate func(IUser) bool
	UserMapPredicate    func(IUser) IUser
	UserCacheCallback   func()

	IUser interface {
		IObject
		// Github returns 'Github' of this 'User' instance.
		Github() string
		// UpdateGithub directly updates 'Github' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateGithub(github string, editor Identity)
		// UpdateGithubAtomic updates 'Github' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateGithubAtomic(transaction ITransaction, github string, editor Identity)
	}

	IUserCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IUser
		Append(user IUser)
		ForEach(UserIterator)
		Array() Users
	}

	IUserManager interface {
		ISystemComponent
		OnCacheChanged(UserCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition UserCondition) bool
		ListUsers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IUserCollection
		GetUser(id int64, editor Identity) (IUser, error)
		AddUser(identityId int64, github string, editor Identity) (IUser, error)
		AddUserObject(identityId int64, user IUser, editor Identity) (IUser, error)
		AddUserAtomic(transaction ITransaction, identityId int64, github string, editor Identity) (IUser, error)
		AddUserObjectAtomic(transaction ITransaction, identityId int64, user IUser, editor Identity) (IUser, error)
		Log(identityId int64, github string, source string, editor Identity, payload string)
		UpdateUser(id int64, github string, editor Identity) (IUser, error)
		UpdateUserObject(id int64, user IUser, editor Identity) (IUser, error)
		UpdateUserAtomic(transaction ITransaction, id int64, github string, editor Identity) (IUser, error)
		UpdateUserObjectAtomic(transaction ITransaction, id int64, user IUser, editor Identity) (IUser, error)
		AddOrUpdateUserObject(id int64, user IUser, editor Identity) (IUser, error)
		AddOrUpdateUserObjectAtomic(transaction ITransaction, id int64, user IUser, editor Identity) (IUser, error)
		RemoveUser(id int64, editor Identity) (IUser, error)
		RemoveUserAtomic(transaction ITransaction, id int64, editor Identity) (IUser, error)
		Find(id int64) IUser
		ForEach(iterator UserIterator)
		Filter(predicate UserFilterPredicate) IUserCollection
		Map(predicate UserMapPredicate) IUserCollection
	}
)
