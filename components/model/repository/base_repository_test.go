package repository_test

import (
	"os"
	"testing"

	"github.com/reiver/greatape/components/model/repository"
	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/settings"
)

//region Initialization

func TestMain(main *testing.M) {
	logger := logging.NewLogger(false)
	configuration := settings.NewTestConfiguration()
	if err := repository.Initialize(configuration, logger); err != nil {
		os.Exit(1)
	}

	os.Exit(main.Run())
}

//endregion
