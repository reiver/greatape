package jobs

import . "rail.town/infrastructure/components/contracts"

func Hourly(x IDispatcher, config string) {
	x.Logger().Debug("✓ SYSTEM_SCHEDULE_HOURLY")
}
