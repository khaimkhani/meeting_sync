package scheduling

type Block struct {
	blockedTime string
	tz          string
	desc        string
	// priority int
	// ^ can override some lower priority mfs
}
