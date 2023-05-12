package CurrentTime

import "time"

func GetCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}
