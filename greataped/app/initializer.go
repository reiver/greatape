package functions

import (
	"rail.town/infrastructure/app/jobs"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

func Initialize(x IDispatcher) error {
	// system schedules
	x.Ensure(
		x.Schedule(jobs.SYSTEM_SCHEDULE_HOURLY, CRON_EVERY_HOUR, jobs.Hourly),
		x.Schedule(jobs.SYSTEM_SCHEDULE_DAILY, CRON_EVERY_DAY_6PM, jobs.Daily),
	)

	return nil
}
