package repository

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"strings"
	"time"

	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/system"
)

//go:embed scripts
var scripts embed.FS

type baseRepository struct {
	name       string
	pluralName string
	entityType reflect.Type
	logger     ILogger
	database   ISqlDatabase
	volatile   bool
}

func newBaseRepository(name, pluralName string, entityType reflect.Type, logger ILogger, volatile bool) baseRepository {
	return baseRepository{
		name:       name,
		pluralName: pluralName,
		entityType: entityType,
		logger:     logger,
		volatile:   volatile,
	}
}

func (repository *baseRepository) Name() string {
	return repository.name
}

func (repository *baseRepository) Migrate() error {
	schema := repository.GetSqlDatabase().GetSchema()
	if schema == nil {
		return fmt.Errorf("invalid_db_schema: %s", repository.GetSqlDatabase().GetName())
	}

	if !repository.volatile {
		createTablesScript := repository.LoadScript("scripts/%s.sql")
		createTriggersScript := repository.LoadScript("scripts/%s_triggers.sql")

		if !schema.HasTable(repository.pluralName) {
			script := createTablesScript + createTriggersScript
			if err := repository.GetSqlDatabase().RunScript(script, "##########"); err != nil {
				return err
			}

			schema = repository.GetSqlDatabase().GetSchema()

			if !schema.HasTable(repository.pluralName) {
				return fmt.Errorf("DB_MIGRATION: %s", repository.pluralName)
			}

			if !schema.HasHistoryTable(repository.pluralName) {
				return fmt.Errorf("DB_MIGRATION: history.%s", repository.pluralName)
			}

			if !schema.HasTrigger(fmt.Sprintf("%s_after_update_trigger", repository.pluralName)) {
				return fmt.Errorf("DB_MIGRATION: %s_after_update_trigger", repository.pluralName)
			}

			if !schema.HasTrigger(fmt.Sprintf("%s_after_delete_trigger", repository.pluralName)) {
				return fmt.Errorf("DB_MIGRATION: %s_after_delete_trigger", repository.pluralName)
			}

			_, err := repository.database.Execute(`INSERT INTO "__system__" ("script") VALUES ($1);`, script)
			if err != nil {
				repository.logger.Alert(fmt.Sprintf("DB_MIGRATION: %s", err))
			}

			repository.logger.Debug(fmt.Sprintf("DB_MIGRATION: ✓ %s.%s", repository.GetSqlDatabase().GetName(), repository.pluralName))
		}

		changes := make([]string, 0)
		commands := make([]string, 0)
		historyCommands := make([]string, 0)

		for i := 0; i < repository.entityType.NumField(); i++ {
			field := repository.entityType.Field(i)
			column := field.Tag.Get("json")
			dbType := field.Tag.Get("storage")
			defaultValue := field.Tag.Get("default")

			if field.Name != "entity" && !schema.HasColumn(repository.pluralName, column) {
				changes = append(changes, fmt.Sprintf("%s.%s", repository.pluralName, column))

				commands = append(commands, fmt.Sprintf(`ALTER TABLE "%s" ADD COLUMN "%s" %s NOT NULL;`, repository.pluralName, column, dbType))
				commands = append(commands, fmt.Sprintf(`ALTER TABLE "%s" ALTER COLUMN "%s" SET DEFAULT %s;`, repository.pluralName, column, defaultValue))

				historyCommands = append(historyCommands, fmt.Sprintf(`ALTER TABLE "%s_history" ADD COLUMN "%s" %s NOT NULL;`, repository.pluralName, column, dbType))
				historyCommands = append(historyCommands, fmt.Sprintf(`ALTER TABLE "%s_history" ALTER COLUMN "%s" SET DEFAULT %s;`, repository.pluralName, column, defaultValue))
			}
		}

		var scriptLines []string
		if len(commands) > 0 {
			scriptLines = append([]string{}, historyCommands...)
			scriptLines = append(scriptLines, fmt.Sprintf(`DROP TRIGGER "%s_after_update_trigger" ON "%s";`, repository.pluralName, repository.pluralName))
			scriptLines = append(scriptLines, fmt.Sprintf(`DROP TRIGGER "%s_after_delete_trigger" ON "%s";`, repository.pluralName, repository.pluralName))
			scriptLines = append(scriptLines, commands...)
			scriptLines = append(scriptLines, createTriggersScript)
			script := strings.Join(scriptLines, "\n##########\n")

			if err := repository.GetSqlDatabase().RunScript(script, "##########"); err != nil {
				return err
			}

			schema = repository.GetSqlDatabase().GetSchema()

			for i := 0; i < repository.entityType.NumField(); i++ {
				field := repository.entityType.Field(i)
				column := field.Tag.Get("json")

				if field.Name != "entity" && !schema.HasColumn(repository.pluralName, column) {
					return fmt.Errorf("DB_MIGRATION: %s.%s", repository.pluralName, column)
				}
			}

			if !schema.HasTrigger(fmt.Sprintf("%s_after_update_trigger", repository.pluralName)) {
				return fmt.Errorf("DB_MIGRATION: %s_after_update_trigger", repository.pluralName)
			}

			if !schema.HasTrigger(fmt.Sprintf("%s_after_delete_trigger", repository.pluralName)) {
				return fmt.Errorf("DB_MIGRATION: %s_after_delete_trigger", repository.pluralName)
			}

			_, err := repository.database.Execute(`INSERT INTO "__system__" ("script") VALUES ($1);`, script)
			if err != nil {
				repository.logger.Alert(fmt.Sprintf("DB_MIGRATION: %s", err))
			}

			for _, change := range changes {
				repository.logger.Debug(fmt.Sprintf("DB_MIGRATION: ✓ %s.%s", repository.GetSqlDatabase().GetName(), change))
			}
		}
	}

	return nil
}

func (repository *baseRepository) LoadScript(resource string) string {
	scriptData, err := scripts.ReadFile(fmt.Sprintf(resource, repository.name))
	if err != nil {
		panic(err)
	}

	return strings.ReplaceAll(string(scriptData), "###DATABASE###", repository.GetSqlDatabase().GetName())
}

func (repository *baseRepository) GetSqlDatabase() ISqlDatabase {
	return repository.database
}

func (repository *baseRepository) SetSqlDatabase(database ISqlDatabase) {
	repository.database = database
}

func (repository *baseRepository) Serialize(pointer Pointer, cause error) {
	if data, err := json.Marshal(pointer); err != nil {
		//TODO: Handle
		repository.logger.Critical(fmt.Sprintf("EMERGENCY SERIALIZATION MARSHALLING FAILURE"))
	} else {
		timestamp := time.Now().UnixNano()
		randomFactor := rand.Intn(1000)
		dataFileName := fmt.Sprintf("%sE_%d_%d.json", repository.logger.SerializationPath(), timestamp, randomFactor)

		if err := ioutil.WriteFile(dataFileName, data, 0644); err != nil {
			//TODO: Handle
			repository.logger.Critical(fmt.Sprintf("EMERGENCY SERIALIZATION PERSISTENCE FAILURE: %s", err))
		}

		errorFileName := fmt.Sprintf("%sE_%d_%d_ERROR.log", repository.logger.SerializationPath(), timestamp, randomFactor)
		if err := ioutil.WriteFile(errorFileName, []byte(cause.Error()), 0644); err != nil {
			//TODO: Handle
			repository.logger.Critical(fmt.Sprintf("EMERGENCY SERIALIZATION ERROR PERSISTENCE FAILURE: %s", err))
		}
	}
}
