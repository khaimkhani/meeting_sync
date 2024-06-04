package scripts

import (
	"msync/models/user"
)

// run this intermittently throughout a day
// look into having a hook listen for changes
func SyncCalendars(u *user.User) {
	if u.googleCalendarKey == "" {
		return
	}
	// call google api
	// loop through events and creats blocks
	// with appropriate tz info

}
