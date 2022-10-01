package repos

var (
	All = []interface{}{
		&User{},
		&IncomingActivity{},
		&OutgoingActivity{},
		&Follower{},
		&Following{},
	}

	Default = &repository{}
)
