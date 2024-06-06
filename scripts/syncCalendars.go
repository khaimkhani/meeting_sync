package scripts

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"msync/models/user"
	"os"
  "json"
)

// run this intermittently throughout a day
// look into having a hook listen for changes
func SyncCalendars(u *user.User) {
	ClientID := os.Getenv("GOOGLE_CLIENT_ID")
	ClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
  
  json.

}
