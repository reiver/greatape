package repository

import (
	"sync"
	"time"

	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/xeronith/diamante/contracts/logging"
)

func (repository *pipeRepository) GetPipeDescriptors() []*pipeDescriptor {
	descriptors := []*pipeDescriptor{
		{
			PIPE_DOCUMENT,
			&sync.Mutex{},
			"INSERT INTO `documents` (`id`, `content`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IDocumentPipeEntity)
				return Parameters{e.Id(), e.Content(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_SYSTEM_SCHEDULE,
			&sync.Mutex{},
			"INSERT INTO `system_schedules` (`id`, `enabled`, `config`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(ISystemSchedulePipeEntity)
				return Parameters{e.Id(), e.Enabled(), e.Config(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_IDENTITY,
			&sync.Mutex{},
			"INSERT INTO `identities` (`id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IIdentityPipeEntity)
				return Parameters{e.Id(), e.Username(), e.PhoneNumber(), e.PhoneNumberConfirmed(), e.FirstName(), e.LastName(), e.DisplayName(), e.Email(), e.EmailConfirmed(), e.Avatar(), e.Banner(), e.Summary(), e.Token(), e.MultiFactor(), e.Hash(), e.Salt(), e.PublicKey(), e.PrivateKey(), e.Permission(), e.Restriction(), e.LastLogin(), e.LoginCount(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_ACCESS_CONTROL,
			&sync.Mutex{},
			"INSERT INTO `access_controls` (`id`, `key`, `value`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IAccessControlPipeEntity)
				return Parameters{e.Id(), e.Key(), e.Value(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_REMOTE_ACTIVITY,
			&sync.Mutex{},
			"INSERT INTO `remote_activities` (`id`, `entry_point`, `duration`, `successful`, `error_message`, `remote_address`, `user_agent`, `event_type`, `timestamp`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IRemoteActivityPipeEntity)
				return Parameters{e.Id(), e.EntryPoint(), e.Duration(), e.Successful(), e.ErrorMessage(), e.RemoteAddress(), e.UserAgent(), e.EventType(), e.Timestamp(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_CATEGORY_TYPE,
			&sync.Mutex{},
			"INSERT INTO `category_types` (`id`, `description`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(ICategoryTypePipeEntity)
				return Parameters{e.Id(), e.Description(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_CATEGORY,
			&sync.Mutex{},
			"INSERT INTO `categories` (`id`, `category_type_id`, `category_id`, `title`, `description`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(ICategoryPipeEntity)
				return Parameters{e.Id(), e.CategoryTypeId(), e.CategoryId(), e.Title(), e.Description(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_USER,
			&sync.Mutex{},
			"INSERT INTO `users` (`id`, `github`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IUserPipeEntity)
				return Parameters{e.Id(), e.Github(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_ACTIVITY_PUB_INCOMING_ACTIVITY,
			&sync.Mutex{},
			"INSERT INTO `activity_pub_incoming_activities` (`id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IActivityPubIncomingActivityPipeEntity)
				return Parameters{e.Id(), e.IdentityId(), e.UniqueIdentifier(), e.Timestamp(), e.From(), e.To(), e.Content(), e.Raw(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_ACTIVITY_PUB_OUTGOING_ACTIVITY,
			&sync.Mutex{},
			"INSERT INTO `activity_pub_outgoing_activities` (`id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IActivityPubOutgoingActivityPipeEntity)
				return Parameters{e.Id(), e.IdentityId(), e.UniqueIdentifier(), e.Timestamp(), e.From(), e.To(), e.Content(), e.Raw(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
		{
			PIPE_ACTIVITY_PUB_FOLLOWER,
			&sync.Mutex{},
			"INSERT INTO `activity_pub_followers` (`id`, `handle`, `inbox`, `subject`, `activity`, `accepted`, `editor`, `queued_at`, `payload`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);",
			func(entity IPipeEntity) Parameters {
				e := entity.(IActivityPubFollowerPipeEntity)
				return Parameters{e.Id(), e.Handle(), e.Inbox(), e.Subject(), e.Activity(), e.Accepted(), e.GetEditor(), e.GetQueueTimestamp().UnixNano(), e.Payload()}
			},
		},
	}

	return descriptors
}

func (repository *pipeRepository) Insert(entities ...IPipeEntity) {
	for _, entity := range entities {
		repository.pipes[entity.GetPipe()].Input() <- entity
	}
}

type pipeRepository struct {
	baseRepository
	pipes      map[int]IPipe
	dispatcher <-chan time.Time
}

func newPipeRepository(logger ILogger) IPipeRepository {
	repository := &pipeRepository{
		baseRepository: newBaseRepository("pipe", "pipes", nil, logger, true),
		pipes:          make(map[int]IPipe),
		dispatcher:     time.NewTicker(AUTO_FLUSH_DURATION).C,
	}

	for _, descriptor := range repository.GetPipeDescriptors() {
		repository.pipes[descriptor.Id()] = NewPipe(descriptor, repository)
	}

	go func() {
		for {
			select {
			case <-repository.dispatcher:
				for _, pipe := range repository.pipes {
					pipe.Signal() <- SIGNAL_FLUSH
				}
			}
		}
	}()

	for _, pipe := range repository.pipes {
		go pipe.OpenValve()
	}

	return repository
}
