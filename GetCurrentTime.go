package CurrentTime

import "time"

func GetCurrentTime() string {
	now := time.Now()
	return now.Format("2006-01-02")
}