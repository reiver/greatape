package jobs

import . "rail.town/infrastructure/components/contracts"

func Daily(x IDispatcher, config string) {
	x.Logger().Debug("✓ SYSTEM_SCHEDULE_DAILY")
}
