package routes

import "contracts"

var All = []contracts.IRoute{
	Health,
	Root,
	Profile,
	Signup,
	Login,
	GetProfile,
	UpdateProfile,
	WebFinger,
	User,
	Message,
	InboxPost,
	InboxGet,
	OutboxPost,
	OutboxGet,
	Followers,
	Follow,
	Authorize,
}
