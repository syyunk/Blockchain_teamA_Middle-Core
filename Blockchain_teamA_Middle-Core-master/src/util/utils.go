package util

import "time"

func GetTimestamp() time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	t := now.In(loc)

	return t
}
