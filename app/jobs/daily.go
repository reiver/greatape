package jobs

import . "github.com/reiver/greatape/components/contracts"

func Daily(x IDispatcher, config string) {
	x.Logger().Debug("✓ SYSTEM_SCHEDULE_DAILY")
}
