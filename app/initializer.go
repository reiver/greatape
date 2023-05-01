package functions

import (
	"github.com/reiver/greatape/app/jobs"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func Initialize(x IDispatcher) error {
	// data seed
	if !x.IdentityExists(0) {
		identity := x.AddIdentityWithCustomId(0, "invalid", "0", false, EMPTY, EMPTY, EMPTY, "invalid@localhost", false, EMPTY, EMPTY, EMPTY, "0", false, EMPTY, EMPTY, EMPTY, EMPTY, NOT_SET, ACL_RESTRICTION_ACCESS_DENIED, NOT_SET, NOT_SET)
		x.AddUser(identity.Id(), EMPTY)
	}

	if !x.IdentityExists(1) {
		identity := x.AddIdentityWithCustomId(1, "root", "1", false, EMPTY, EMPTY, EMPTY, "root@localhost", false, EMPTY, EMPTY, EMPTY, "1", true, EMPTY, EMPTY, EMPTY, EMPTY, ACL_PERMISSION_ADMIN, ACL_RESTRICTION_NONE, NOT_SET, NOT_SET)
		x.AddUser(identity.Id(), EMPTY)
	}

	if !x.CategoryTypeExists(0) {
		categoryType := x.AddCategoryTypeWithCustomId(0, EMPTY)
		_ = x.AddCategoryWithCustomId(0, categoryType.Id(), NOT_SET, EMPTY, EMPTY)
	}

	// system schedules
	x.Ensure(
		x.Schedule(jobs.SYSTEM_SCHEDULE_HOURLY, CRON_EVERY_HOUR, jobs.Hourly),
		x.Schedule(jobs.SYSTEM_SCHEDULE_DAILY, CRON_EVERY_DAY_6PM, jobs.Daily),
	)

	return nil
}
