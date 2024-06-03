package scheduling

import (
	"msync/models/user"
)

type Day struct {
	id   int
	date string
	user *user.User

	// would have descripts possibly
	taken []*Block
}

func (d *Day) GetOpenBlocks(tz string) []string {
	// adjust for tz?
	// or give tz in reply

	return []string{"0200-0300", "0400-1200"}
}
