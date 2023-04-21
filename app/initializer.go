package functions

import (
	"github.com/reiver/greatape/app/jobs"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func Initialize(x IDispatcher) error {
	// system schedules
	x.Ensure(
		x.Schedule(jobs.SYSTEM_SCHEDULE_HOURLY, CRON_EVERY_HOUR, jobs.Hourly),
		x.Schedule(jobs.SYSTEM_SCHEDULE_DAILY, CRON_EVERY_DAY_6PM, jobs.Daily),
	)

	return nil
}
