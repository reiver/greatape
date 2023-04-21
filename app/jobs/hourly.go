package jobs

import . "github.com/reiver/greatape/components/contracts"

func Hourly(x IDispatcher, config string) {
	x.Logger().Debug("✓ SYSTEM_SCHEDULE_HOURLY")
}
