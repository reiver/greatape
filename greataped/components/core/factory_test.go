package core_test

import (
	"os"
	"testing"

	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/settings"
	"rail.town/infrastructure/components/core"
	"rail.town/infrastructure/components/model/repository"
)

//region Initialization

func TestMain(main *testing.M) {
	logger := logging.NewLogger(false)
	configuration := settings.NewTestConfiguration()
	if err := repository.Initialize(configuration, logger); err != nil {
		os.Exit(1)
	}

	if err := core.Initialize(configuration, logger); err != nil {
		os.Exit(1)
	}

	os.Exit(main.Run())
}

//endregion
