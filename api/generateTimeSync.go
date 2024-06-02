package api

import (
	"msync/models/user"
)

type User = user.User

func generateTimeSyncLink(u *User) string {

	// create a link accessible by any to be able to
	// Probably something with a UUID which expires after x minutes
	// meetingsync.com/s3ZXd-Aqkaz-11SDz

	// store the UUID in DB or cache for newly created ones linking to a user

	// TimeSync.create(u, ...)

	return "xyz.com/pickATime"
}
