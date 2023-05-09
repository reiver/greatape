package commands

import . "github.com/reiver/greatape/components/contracts"

func UpdateProfileByUser(x IDispatcher,
	displayName string,
	avatar string,
	banner string,
	summary string,
	github string,
) (IUpdateProfileByUserResult, error) {
	editor := x.Identity()
	identity := editor.(IIdentity)
	user := x.GetUser(identity.Id())

	x.Atomic(func() error {
		identity.UpdateDisplayNameAtomic(x.Transaction(), displayName, editor)
		identity.UpdateAvatarAtomic(x.Transaction(), avatar, editor)
		identity.UpdateBannerAtomic(x.Transaction(), banner, editor)
		identity.UpdateSummaryAtomic(x.Transaction(), summary, editor)
		user.UpdateGithubAtomic(x.Transaction(), github, editor)
		return nil
	})

	identity = x.GetIdentity(x.Identity().Id())
	user = x.GetUser(identity.Id())

	return x.NewUpdateProfileByUserResult(
		identity.DisplayName(),
		identity.Avatar(),
		identity.Banner(),
		identity.Summary(),
		user.Github(),
	), nil
}
