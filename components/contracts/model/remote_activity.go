package model

type (
	RemoteActivityEntities []IRemoteActivityEntity

	IRemoteActivityEntity interface {
		IEntity
		EntryPoint() string
		Duration() int64
		Successful() bool
		ErrorMessage() string
		RemoteAddress() string
		UserAgent() string
		EventType() uint32
		Timestamp() int64
	}

	IRemoteActivityPipeEntity interface {
		IRemoteActivityEntity
		IPipeEntity
	}

	IRemoteActivitiesRepository interface {
		IRepository
		Add(entity IRemoteActivityEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error
		FetchById(editor int64) (IRemoteActivityEntity, error)
		Update(entity IRemoteActivityEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error
		Remove(entity IRemoteActivityEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IRemoteActivityEntity, editor int64) error
		FetchAll() (RemoteActivityEntities, error)
		UpdateEntryPoint(id int64, value string, editor int64) error
		UpdateEntryPointAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateDuration(id int64, value int64, editor int64) error
		UpdateDurationAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error
		UpdateSuccessful(id int64, value bool, editor int64) error
		UpdateSuccessfulAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
		UpdateErrorMessage(id int64, value string, editor int64) error
		UpdateErrorMessageAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateRemoteAddress(id int64, value string, editor int64) error
		UpdateRemoteAddressAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateUserAgent(id int64, value string, editor int64) error
		UpdateUserAgentAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateEventType(id int64, value uint32, editor int64) error
		UpdateEventTypeAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error
		UpdateTimestamp(id int64, value int64, editor int64) error
		UpdateTimestampAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error
	}
)
