package user

type User struct {
	id    int
	email string
	// googleCalendarKey string
	// alt calendars ...

	// this could be a TZ object with a string representation possibly
	tz string
}

func (u *User) create(atts map[string]string) int {

	// returns id after creating

	// DB.QueryRows("")
	// DB.Exec("")

	return 1
}
