package CurrentTime

import "time"

func GetCurrentTime() string {
	now := time.Now()
	return now.Format("03:04 PM")
}