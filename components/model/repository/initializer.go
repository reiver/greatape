package repository

import (
	"errors"
	"strings"

	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/database/drivers/postgres"
)

var (
	Pipe                          IPipeRepository
	Documents                     IDocumentsRepository
	SystemSchedules               ISystemSchedulesRepository
	Identities                    IIdentitiesRepository
	AccessControls                IAccessControlsRepository
	RemoteActivities              IRemoteActivitiesRepository
	CategoryTypes                 ICategoryTypesRepository
	Categories                    ICategoriesRepository
	Users                         IUsersRepository
	ActivityPubIncomingActivities IActivityPubIncomingActivitiesRepository
	ActivityPubOutgoingActivities IActivityPubOutgoingActivitiesRepository
	ActivityPubFollowers          IActivityPubFollowersRepository
)

var database ISqlDatabase

func Initialize(configuration IConfiguration, logger ILogger) error {
	databaseName := configuration.GetPostgreSQLConfiguration().GetDatabase()
	if strings.TrimSpace(databaseName) == "" {
		return errors.New("database_required")
	}

	database = NewDatabase(configuration, logger, databaseName)
	if err := database.Initialize(); err != nil {
		return err
	}

	Pipe = newPipeRepository(logger)
	Documents = newDocumentsRepository(logger)
	SystemSchedules = newSystemSchedulesRepository(logger)
	Identities = newIdentitiesRepository(logger)
	AccessControls = newAccessControlsRepository(logger)
	RemoteActivities = newRemoteActivitiesRepository(logger)
	CategoryTypes = newCategoryTypesRepository(logger)
	Categories = newCategoriesRepository(logger)
	Users = newUsersRepository(logger)
	ActivityPubIncomingActivities = newActivityPubIncomingActivitiesRepository(logger)
	ActivityPubOutgoingActivities = newActivityPubOutgoingActivitiesRepository(logger)
	ActivityPubFollowers = newActivityPubFollowersRepository(logger)

	repositories := []IRepository{
		Pipe,
		Documents,
		SystemSchedules,
		Identities,
		AccessControls,
		RemoteActivities,
		CategoryTypes,
		Categories,
		Users,
		ActivityPubIncomingActivities,
		ActivityPubOutgoingActivities,
		ActivityPubFollowers,
	}

	for _, repository := range repositories {
		repository.SetSqlDatabase(database)
	}

	for _, repository := range repositories {
		if err := repository.Migrate(); err != nil {
			return err
		}
	}

	return nil
}

func WithTransaction(handler RepositoryTransactionHandler) error {
	if database == nil {
		panic("repository_not_initialized")
	}

	return database.WithTransaction(func(transaction ISqlTransaction) error {
		return handler(transaction)
	})
}
